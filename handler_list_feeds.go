package main

import (
	"context"
	"fmt"
	"log"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
)

func handler_list_feeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		log.Fatal("error reading feeds from DB ", err)
	}
	print_feeds(s, feeds)
	return nil
}

func print_feeds(s *state, feeds []database.ListFeedsRow) {
	for _, feed := range(feeds) {
		fmt.Printf(" * %s | URL: %s | Added By %s\n", feed.Name, feed.Url, feed.UserName)
	}
}
