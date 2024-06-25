package middleware

import (
	"awesomeProject/internal/pkg/util/jwter"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func New(log *logrus.Logger /*, authJwt *jwter.JWTer*/) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("start auth middleware")
			header := r.Header
			value := header.Get("Authorization")
			jwtToken := strings.Replace(value, "Bearer ", "", 1)
			if jwtToken == "" {
				log.Error("jwt token is empty")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			username, err := jwter.DecodeToken(jwtToken)
			if err != nil {
				log.Error("jwt token is invalid")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Printf("username is %s", username)
			ctx := context.WithValue(r.Context(), "username", username)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
