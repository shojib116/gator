package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shojib116/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("no name and url provided \n usage: cli addfeed [name] [url]")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Println("Successfully creeated feed")
	printFeed(feed, user)
	fmt.Println()

	fmt.Println("following the feed for you")
	_, err = createFeedFollow(s, user, feed)
	if err != nil {
		return fmt.Errorf("failed to follow the feed for you: %w", err)
	}

	fmt.Printf("Successfully followed the feed: %s\n", feed.Name)
	fmt.Println("=====================================")

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch feeds: %w", err)
	}

	for _, feed := range feeds {
		printFeed(database.Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
		}, database.User{Name: feed.Username})

		fmt.Println("=====================================")

	}
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}
