package middleware

import (
	"context"
	"errors"
	contract "github.com/enfil/metamask-auth/internal/contract/service"
	"github.com/enfil/metamask-auth/internal/domain/user"
	"github.com/enfil/metamask-auth/pkg/reader"
	"net/http"
)

type Auth struct {
	TokenProvider contract.TokenProvider
	UserReader    reader.User
}

func (auth *Auth) Handle() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headerValue := r.Header.Get("Authorization")
			const prefix = "Bearer "
			if len(headerValue) < len(prefix) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			tokenString := headerValue[len(prefix):]
			if len(tokenString) == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			subj, err := auth.TokenProvider.VerifyAndGetSubject(tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			u, err := auth.UserReader.ByAddress(subj)
			if err != nil {
				if errors.Is(err, user.ErrUserNotExists) {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), "user", u)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
