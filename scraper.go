package main

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Hamidspirit/http_server.git/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration.", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error fetcing feeds;", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed fetched: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("error fetching feed %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		pubDate, err := time.Parse(time.RFC1123Z, item.PublishDate)
		if err != nil {
			log.Printf("Couldn't parse date %v with error %v", item.PublishDate, err)
			continue
		}

		newUUID := uuid.New()
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          pgtype.UUID{Bytes: newUUID, Valid: true},
			CreatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
			UpdatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
			Title:       item.Title,
			Description: pgtype.Text{String: item.Description},
			PublishedAt: pgtype.Timestamp{Time: pubDate, Valid: true},
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to crearte post:", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
