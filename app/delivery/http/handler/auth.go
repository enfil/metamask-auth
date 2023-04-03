package handler

import (
	"errors"
	"github.com/enfil/metamask-auth/app/delivery/http/request"
	"github.com/enfil/metamask-auth/app/delivery/http/response"
	"github.com/enfil/metamask-auth/app/reader"
	contract "github.com/enfil/metamask-auth/contract/service"
	"github.com/enfil/metamask-auth/domain/user"
	"github.com/enfil/metamask-auth/usecase"
	command2 "github.com/enfil/metamask-auth/usecase/command"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

type Auth struct {
	TokenProvider contract.TokenProvider
	UserReader    reader.User
	Registrar     usecase.Registrar
	SignIn        usecase.SignIn
}

func (auth *Auth) RegistrationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c command2.Register
		if err := request.BindReqBody(r, &c); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := auth.Registrar.Handle(c)
		if err != nil {
			writeErrorHeaders(err, w)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (auth *Auth) UserNonceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address := chi.URLParam(r, "address")
		if !user.HexRegex.MatchString(address) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := auth.UserReader.ByAddress(strings.ToLower(address))
		if err != nil {
			writeErrorHeaders(err, w)
			return
		}

		resp := struct {
			Nonce string
		}{
			Nonce: u.Nonce(),
		}

		response.RenderJson(r, w, http.StatusOK, resp)
	}
}

func (auth *Auth) SignInHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c command2.SignIn
		if err := request.BindReqBody(r, &c); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c.CryptoAddress = strings.ToLower(c.CryptoAddress)
		err := auth.SignIn.Handle(c)
		if err != nil {
			writeErrorHeaders(err, w)
			return
		}
		signedToken, err := auth.TokenProvider.Create(c.CryptoAddress)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp := struct {
			AccessToken string `json:"access"`
		}{
			AccessToken: signedToken,
		}
		response.RenderJson(r, w, http.StatusOK, resp)
	}
}

func (auth *Auth) CheckAuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		resp := response.CheckAuthResponse{u.Uuid(), u.CryptoAddress()}
		response.RenderJson(r, w, http.StatusOK, resp)
	}
}

func (auth *Auth) WelcomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := request.GetUserFromReqContext(r)
		resp := struct {
			Msg string `json:"msg"`
		}{
			Msg: "Congrats " + u.CryptoAddress() + " you made it",
		}
		response.RenderJson(r, w, http.StatusOK, resp)
	}
}

func writeErrorHeaders(err error, w http.ResponseWriter) {
	switch err {
	case user.ErrUserNotExists:
		w.WriteHeader(http.StatusNotFound)
	case user.ErrUserExists:
		w.WriteHeader(http.StatusConflict)
	case user.ErrInvalidAddress:
		w.WriteHeader(http.StatusBadRequest)
	case user.ErrAuthError:
		w.WriteHeader(http.StatusUnauthorized)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
