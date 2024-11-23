package main

import (
	"fmt"
	"github.com/Alpha-Knight-Zero/rss-aggregator/internal/auth"
	"github.com/Alpha-Knight-Zero/rss-aggregator/internal/database"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth Error: %s", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("User not found: %s", err))
			return
		}

		handler(w, r, user)
	}
}
