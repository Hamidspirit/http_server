package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hamidspirit/http_server.git/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	newUUID := uuid.New()
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        pgtype.UUID{Bytes: newUUID, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating feed: %v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, databseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get the feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databseFeedsToFeeds(feed))
}
