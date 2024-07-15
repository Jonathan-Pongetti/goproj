package main

import (

	"fmt"
	"net/http"

	"github.com/Jonathan-Pongetti/goproj/internal/auth"
	"github.com/Jonathan-Pongetti/goproj/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error ", err))
			return
		}
	
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("couldnt get user ", err))
			return
		}

		handler(w, r, user)
	}
}