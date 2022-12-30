package feedly

import (
	"reflect"
	"testing"
)

func TestNewRSS(t *testing.T) {
	c := NewClient()

	type args struct {
		src []byte
	}

	tests := make([]struct {
		name  string
		args  args
		want  *Rss
		want1 UnMarshallError
	}, 0)

	testUrls := []string{
		"https://www.caranddriver.com/rss/all.xml/",
		"https://jalopnik.com/rss",
		"https://www.autoblog.com/rss.xml",
		"http://feeds.feedburner.com/autonews/AutomakerNews",
		"http://feeds.feedburner.com/autonews/SupplierNews",
		"http://feeds.feedburner.com/autonews/EuropeNews",
		"http://feeds.feedburner.com/MobilityReport",
		"http://feeds.feedburner.com/autonews/RetailNews",
		"https://www.carthrottle.com/rss/",
		"https://feeds.highgearmedia.com/",
		"https://www.thetruthaboutcars.com/rss/feed/all",
		"https://insideevs.com/rss/articles/all/",
	}

	for _, url := range testUrls {
		t.Run(url, func(t *testing.T) {
			code, res, err := c.FetchFeed(url)
			if err.error != nil {
				t.Errorf("FetchFeed() err = %v", err)
			}
			if code != 200 {
				t.Errorf("FetchFeed() code = %v", code)
			}
			tests = append(tests, struct {
				name  string
				args  args
				want  *Rss
				want1 UnMarshallError
			}{
				name: url,
				args: args{
					src: res,
				},
				want:  &Rss{},
				want1: UnMarshallError{nil},
			})
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewRSS(tt.args.src)
			if got.Channel.Title == "" {
				t.Errorf("NewRSS() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewRSS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
