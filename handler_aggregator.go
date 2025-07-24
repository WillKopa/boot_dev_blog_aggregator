package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
)

func handler_aggregator(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("time between requests not set, please set it when calling this command")
	}
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing duration: %v", err)
	}
	
	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ;; <-ticker.C {
		err = scrapeFeeds(s, cmd)
		if err != nil {
			return fmt.Errorf("error scraping feeds: %v", err)
		}
	}
}

func scrapeFeeds(s *state, cmd command) error {
	next_feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed from DB: %v", err)
	}
	
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},	
		ID: next_feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v", err)
	}
	
	feed, err := fetchFeed(context.Background(), next_feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed from url: %v", err)
	}
	
	print_feed_response(feed)
	return nil
}

func print_feed_response(r *RSSFeed) {
	fmt.Printf("Feed: %s\n%s", r.Channel.Title, strings.Repeat("-", 20))
	for _, item := range(r.Channel.Item) {
		fmt.Printf(" * %s\n", item.Title)
	}
}
