package main

import (
	"github.com/Alpha-Knight-Zero/rss-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		URL:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeed {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}
	return feeds
}
