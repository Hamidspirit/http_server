package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Hamidspirit/http_server.git/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	newUUID := uuid.New()
	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        pgtype.UUID{Bytes: newUUID, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UserID:    user.ID,
		FeedID:    pgtype.UUID{Bytes: params.FeedID, Valid: true},
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating feed_follow: %v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *apiConfig) handlerGettingFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error Getting feed follows: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feed_follows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdString := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIdString)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parse feedFollowID : %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     pgtype.UUID{Bytes: feedFollowID, Valid: true},
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't delete feed follow  : %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
