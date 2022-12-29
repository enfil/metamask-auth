package main

import (
	"github.com/enfil/metamask-auth/app/config"
	delivery "github.com/enfil/metamask-auth/app/delivery/http"
	"github.com/enfil/metamask-auth/app/delivery/http/handler"
	"github.com/enfil/metamask-auth/app/delivery/http/middleware"
	"github.com/enfil/metamask-auth/app/reader"
	"github.com/enfil/metamask-auth/app/repository/user/memory"
	"github.com/enfil/metamask-auth/app/service"
	usecase2 "github.com/enfil/metamask-auth/usecase"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var (
	settings    = config.LoadAndStoreConfig(".env")
	userRepo    = memory.NewMemoryRepository()
	jwtProvider = service.NewJwt(
		settings.JWT.Secret,
		settings.JWT.Issuer,
		time.Minute*time.Duration(settings.JWT.TTL),
	)
	nonceProvider = service.Nonce{}
	signProvider  = service.Sign{}
	authHandler   = handler.Auth{
		TokenProvider: jwtProvider,
		Registrar:     usecase2.Registrar{Repo: userRepo, NonceProvider: &nonceProvider},
		SignIn:        usecase2.SignIn{Repo: userRepo, NonceProvider: &nonceProvider, SignProvider: signProvider},
		UserReader:    reader.User{ReadModel: userRepo},
	}
	authMiddleware = middleware.Auth{
		TokenProvider: jwtProvider,
		UserReader:    reader.User{ReadModel: userRepo},
	}
)

func main() {
	// set up the endpoints
	r := chi.NewRouter()
	//  Just allow all for the reference implementation
	r.Use(cors.AllowAll().Handler)

	delivery.RegisterRoutes(r, authHandler, authMiddleware)

	err := http.ListenAndServe("localhost:8001", r)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
