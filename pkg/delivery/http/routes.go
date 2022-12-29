package delivery

import (
	"github.com/enfil/metamask-auth/pkg/delivery/http/handler"
	"github.com/enfil/metamask-auth/pkg/delivery/http/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func RegisterRoutes(r *chi.Mux, authHandler handler.Auth, authMiddleware middleware.Auth) {
	//  Just allow all for the reference implementation
	r.Use(cors.AllowAll().Handler)

	r.Post("/register", authHandler.RegistrationHandler())
	r.Get("/users/{address:^0x[a-fA-F0-9]{40}$}/nonce", authHandler.UserNonceHandler())
	r.Post("/signin", authHandler.SignInHandler())
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.Handle())
		r.Get("/welcome", authHandler.WelcomeHandler())
	})
}
