package main

import (
	"encoding/json"
	"fmt"
	"github.com/Alpha-Knight-Zero/rss-aggregator/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Request could not be parsed as JSON: %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create feed: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get any feeds: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
