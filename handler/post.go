package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/racsoJosu/rest-ws/models"
	"github.com/racsoJosu/rest-ws/repository"
	"github.com/racsoJosu/rest-ws/server"
	"github.com/segmentio/ksuid"
)

type PostRequest struct {
	PostContent string `json:"post_content"`
}
type PostResponse struct {
	ID          string `json:"id"`
	PostContent string `json:"post_content"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var request = PostRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return

			}

			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := models.Post{
				ID:          id.String(),
				PostContent: request.PostContent,
				UserId:      claims.UserId,
			}

			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				ID:          post.ID,
				PostContent: post.PostContent,
			})
			return

			// user, err := repository.GetUserById(r.Context(), claims.UserId)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	return
			// }
			// w.Header().Set("Content-Type", "Application/json")
			// json.NewEncoder(w).Encode(user)

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
