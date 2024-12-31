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

func (apiCfg *apiConfig) handleCreatUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	newUUID := uuid.New()
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        pgtype.UUID{Bytes: newUUID, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating user: %v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, databseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
