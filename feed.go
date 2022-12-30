package feedly

import (
	"encoding/xml"
	"fmt"
)

// Item represents a single item in the RSS feed.
// It contains the title, description, link, and publication date.
type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	// media:content
	// media:thumbnail

	Media struct {
		Content struct {
			Url string `xml:"url,attr"`
		}
	} `xml:"media:content"`

	// Content
	Content string `xml:"content:encoded"`

	// Categories
	Categories []string `xml:"category"`

	// Enclosure
	Enclosure struct {
		Url string `xml:"url,attr"`
	} `xml:"enclosure"`
}

// Channel represents the RSS feed.
// It contains the title, description, link, and items.
type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Rss struct {
	// Rss represents the root of the RSS feed.
	// It contains the channel.
	Channel Channel `xml:"channel"`
}

type Feed struct {
	// Feed represents the RSS feed.
	// It contains the channel.
	Channel Channel `xml:"channel"`
}

type UnMarshallError struct {
	error
}

func NewRSS(src []byte) (*Rss, UnMarshallError) {
	var rss Rss
	err := xml.Unmarshal(src, &rss)
	if err != nil {
		return nil, UnMarshallError{err}
	}

	// check if media:content is empty
	for i, item := range rss.Channel.Items {
		if item.Media.Content.Url == "" {
			rss.Channel.Items[i].Media.Content.Url = item.Enclosure.Url
		}
	}

	// check if media content is empty and it to the first image in the content. img src
	for i, item := range rss.Channel.Items {
		if item.Media.Content.Url == "" && item.Content != "" {
			var imgSrc string
			_, err := fmt.Sscanf(item.Content, "<img src=\"%s\"", &imgSrc)
			if err == nil {
				rss.Channel.Items[i].Media.Content.Url = imgSrc
			}
		}
	}

	// check if media content is empty and it to the first image in the description. img src
	for i, item := range rss.Channel.Items {
		if item.Media.Content.Url == "" && item.Description != "" {
			var imgSrc string
			_, err := fmt.Sscanf(item.Description, "<img src=\"%s\"", &imgSrc)
			if err == nil {
				rss.Channel.Items[i].Media.Content.Url = imgSrc
			}
		}
	}

	return &rss, UnMarshallError{nil}
}
