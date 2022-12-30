![TEST COVERAGE](https://img.shields.io/badge/TEST%20COVERAGE-100%25-green?labelColor=GREEN&style=flat)
# Feedly: RSS Parser
This is a Go library for parsing RSS feeds. It includes functions for making HTTP requests to retrieve the feed, unmarshalling the response into structs, and handling errors. It also includes functions for parsing multiple feeds concurrently and specifying different error handling strategies. The library is designed to be easy to use and highly customizable.


## Installation
To install this Go library, you will need to have Go installed on your system. You can then use the go get command to install the library:

```go get github.com/oyamo/feedly```
This will download the library and install it in your $GOPATH. You can then import the library into your Go code using the following import path:

```go
import "github.com/oyamo/feedly"
```

Alternatively, you can clone the repository and install the library manually:
```bash
git clone https://github.com/oyamo/feedly.git
cd rss-parser
go install
```

## Usage
To use this Go library to parse an RSS feed, you can use the `ParseFromUrl` function. This function takes a 
single URL as input and returns the Rss struct for the feed at that URL, or an error if one occurred.

For example:
```go
package main

import (
	"fmt"
	"github.com/oyamo/feedly"
)

func main() {
	url := "https://example.com/rss"
	rss, err := feedly.ParseFromUrl(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rss)
}
```

You can also use the `ParseMultipleFeeds` function to parse multiple feeds concurrently. This function takes a slice of URLs
and an error handling strategy as input, and returns a slice of FeedRes structs and any errors that occurred. The error 
handling strategy can be either `IGNORE` or `ABORT`, which determines how the function handles errors when parsing a feed.
If the `ABORT` strategy is specified and an error occurs, the function cancels all other requests and returns the error.

For example:

```go
package main

import (
	"fmt"
	"github.com/oyamo/feedly"
)

func main() {
	urls := []string{"https://example.com/rss", "https://another-example.com/rss"}
	strategy := feedly.IGNORE
	feeds, err := feedly.ParseMultipleFeeds(urls, strategy)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, feed := range feeds {
		fmt.Println(feed.Url, feed.Rss, feed.Err, feed.Code)
	}
}
```


This will parse the two RSS feeds concurrently and print the results 
for each feed. If an error occurred while parsing a feed, it will be 
included in the `Err` field of the corresponding `FeedRes` struct. If the `IGNORE` strategy was specified, 
the function will continue parsing the remaining feeds even if an error occurred. If the `ABORT` strategy was specified,
the function will cancel all other requests and return the error if an error occurred while parsing a feed.

## Testing
This library has a 100% test coverage. To run the automated tests for this Go library, you will need to have Go 
installed on your system and have the library installed in your $GOPATH. 
Then, you can use the go test command to run the tests:

```
go test github.com/oyamo/feedly
```

## LICENSE
```
MIT License

Copyright (c) 2022 Oyamo Brian

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
