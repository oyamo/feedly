package feedly

import (
	"fmt"
	"testing"
)

func TestParseMultipleFeeds(t *testing.T) {
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

	type args struct {
		urls     []string
		strategy OnErrorStrategy
	}

	tests := []struct {
		name    string
		args    args
		wantRss []FeedRes
		wantErr bool
	}{
		{
			name: "TestParseMultipleFeeds",
			args: args{
				urls:     testUrls,
				strategy: IGNORE,
			},
			wantRss: make([]FeedRes, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseMultipleFeeds(tt.args.urls, tt.args.strategy)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMultipleFeeds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestParseFromUrl(t *testing.T) {
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

	type args struct {
		url string
	}
	tests := make([]struct {
		name    string
		args    args
		wantRss *Rss
		wantErr bool
	}, 0)

	for _, url := range testUrls {
		tests = append(tests, struct {
			name    string
			args    args
			wantRss *Rss
			wantErr bool
		}{
			name: fmt.Sprintf("TestParseFromUrl %s", url),
			args: args{
				url: url,
			},
			wantRss: &Rss{},
			wantErr: false,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseFromUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFromUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
