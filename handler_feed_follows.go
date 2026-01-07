package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shojib116/gator/internal/database"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("no url provided \n usage: cli follow [url]")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("no feed found with the given url!: %w", err)
	}

	newFeedFollow, err := createFeedFollow(s, user, feed)
	if err != nil {
		return fmt.Errorf("failed to follow the feed: %w", err)
	}

	fmt.Printf("%s just followed the feed \"%s\"", newFeedFollow.Username, newFeedFollow.FeedName)

	return nil
}

func handlerFeedUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("no url provided \n usage: cli follow [url]")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("no feed found with the given url!: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow the feed: %w", err)
	}

	fmt.Printf("%s unfollowed the feed \"%s\"", user.Name, feed.Name)

	return nil
}

func handlerListFeedFollowing(s *state, _ command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to fetch latest feed follows: %w", err)
	}

	if len(feedFollows) > 0 {
		fmt.Println("you are currently following:")
	} else {
		fmt.Println("you are not following any feed yet")
	}
	for _, feedFollow := range feedFollows {
		fmt.Println("*", feedFollow.FeedName)
	}

	return nil
}

func createFeedFollow(s *state, user database.User, feed database.Feed) (database.CreateFeedFollowRow, error) {
	return s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
}
