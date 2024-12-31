package main

import (
	"github.com/Hamidspirit/http_server.git/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID      `json"id"`
	CreatedAt pgtype.Timestamp `json"created_at"`
	UpdatedAt pgtype.Timestamp `json"updated_at"`
	Name      string           `json"name"`
	APIKey    string           `json"api_key"`
}

func databseUserToUser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		APIKey:    dbuser.ApiKey,
	}
}

type Feed struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Name      string           `json:"name"`
	Url       string           `json:"url"`
	UserID    pgtype.UUID      `json:"user_id"`
}

func databseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databseFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbfeed := range dbFeed {
		feeds = append(feeds, databseFeedToFeed(dbfeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	UserID    pgtype.UUID      `json:"user_id"`
	FeedID    pgtype.UUID      `json"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feed_follows := []FeedFollow{}
	for _, dbFeedFollows := range dbFeedFollows {
		feed_follows = append(feed_follows, databaseFeedFollowToFeedFollow(dbFeedFollows))
	}
	return feed_follows
}

type Post struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Title       string           `json:"title"`
	Description *string          `json:"description"`
	PublishedAt pgtype.Timestamp `json:"published_at"`
	Url         string           `json:"url"`
	FeedID      pgtype.UUID      `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}
