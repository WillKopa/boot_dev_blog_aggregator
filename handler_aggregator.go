package main

import (
	"context"
	"database/sql"
	"fmt"
	// "strings"
	"time"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
	"github.com/google/uuid"
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
	
	return save_feed(s, feed, next_feed)
}

// func print_feed_response(r *RSSFeed) {
// 	line_break := strings.Repeat("-", 20)
// 	fmt.Printf("Feed: %s\n%s\n%s\n", r.Channel.Title, line_break, line_break)
// 	for _, item := range(r.Channel.Item) {
// 		fmt.Printf(" * %s\n", item.Title)
// 	}
// }

func save_feed(s *state, r *RSSFeed, feed database.Feed) error {
	for _, item := range(r.Channel.Item) {
		pub_time, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return fmt.Errorf("error parsing time: %v\nerror: %v", item.PubDate, err)
		}
		params := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			PublishedAt: pub_time,
			Url: item.Link,
			Title: item.Title,
			Description: item.Description,
			FeedID: feed.ID,
		}
		post, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			fmt.Printf("error saving post: %v\n", err)
		} else {
			fmt.Printf("Saved Post: %v\n", post.Title)
		}
	}
	return nil
}