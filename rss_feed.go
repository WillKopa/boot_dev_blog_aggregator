package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title			string		`xml:"title"`
		Link			string		`xml:"link"`
		Description		string		`xml:"description"`
		Item			[]RSSItem	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title			string		`xml:"title"`
	Link			string		`xml:"link"`
	Description		string		`xml:"description"`	
	PubDate			string		`xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("unable to create request with context for url %s: %v", feedUrl, err)
	}

	req.Header.Set("User-Agent", "gator")
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := httpClient.Do(req) 
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error getting rss feed: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading body of response: %v", err)
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshaling data in response: %v", err)
	}

	escapeHTML(&feed)
	return &feed, nil
}

func escapeHTML(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for idx, item := range(feed.Channel.Item) {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[idx] = item
	}
}