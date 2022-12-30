package feedly

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
)

type OnErrorStrategy int

// Error strategy
const (
	IGNORE = iota // ignore error
	ABORT         // Cancel all other requests
)

type FeedRes struct {
	Url  string
	Rss  *Rss
	Err  error
	Code uint32
}

type feedArgs struct {
	ctx      context.Context
	client   *Client
	outChan  chan FeedRes
	strategy OnErrorStrategy
}

// ParseFromUrl parses the RSS feed and returns a list of items.
func ParseFromUrl(url string) (rss *Rss, err error) {
	c := NewClient()
	code, res, reserr := c.FetchFeed(url)
	if reserr.error != nil {
		return nil, reserr
	}
	if code != 200 {
		return nil, err
	}

	rss, unmarshallerr := NewRSS(res)
	if unmarshallerr.error != nil {
		return nil, unmarshallerr
	}
	return rss, nil
}

func parseFeedChunk(args *feedArgs, urlChunk []string) {
	for _, url := range urlChunk {
		fmt.Println("Parsing url: ", url)
		select {
		case <-args.ctx.Done():
			return
		default:
			code, rssp, err := args.client.FetchFeed(url)
			if err.error != nil {
				if args.strategy == ABORT {
					return
				} else {
					args.outChan <- FeedRes{Url: url, Rss: nil, Err: err, Code: code}
					continue
				}
			}
			if code != 200 {
				fmt.Println("Error: ", err)
				if args.strategy == ABORT {
					return
				} else {
					args.outChan <- FeedRes{Url: url, Rss: nil, Err: err, Code: code}
					continue
				}
			}

			rss, rsserr := NewRSS(rssp)

			if rsserr.error != nil {
				if args.strategy == ABORT {
					return
				} else {
					args.outChan <- FeedRes{Url: url, Rss: nil, Err: rsserr, Code: code}
					continue
				}
			}
			args.outChan <- FeedRes{Url: url, Rss: rss, Err: nil, Code: code}
		}
	}
}

func ParseMultipleFeeds(urls []string, strategy OnErrorStrategy) (rss []FeedRes, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	args := feedArgs{
		ctx:      ctx,
		client:   NewClient(),
		outChan:  make(chan FeedRes),
		strategy: strategy,
	}

	var feeds = make([]FeedRes, 0)
	var nextChunk = 0
	for i := 0; i < len(urls); i += runtime.NumCPU() {
		if i+runtime.NumCPU() > len(urls) {
			nextChunk = len(urls) - 1
		} else {
			nextChunk = i + runtime.NumCPU()
		}
		go parseFeedChunk(&args, urls[i:nextChunk])
	}

	var count atomic.Uint32
	for res := range args.outChan {
		count.Add(1)
		if res.Err != nil && strategy == ABORT {
			cancel()
			close(args.outChan)
			return nil, res.Err
		}
		feeds = append(feeds, res)
		if count.Load() == uint32(len(urls)) {
			close(args.outChan)
		}
	}
	return feeds, nil
}
