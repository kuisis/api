package main

import (
	"errors"
	"github.com/danilopolani/gocialite"
	"github.com/patrickmn/go-cache"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
	"strings"
	"time"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

var gocial = gocialite.NewDispatcher()
var VerifyStorage = cache.New(24*time.Hour, 2*time.Hour)
var PassResetStorage = cache.New(10*time.Minute, 5*time.Minute)
var Rnd = renderer.New()

func init() {
	Router.HandleFunc("/logout", handlerLogout).Methods("GET")
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
	clientId := r.URL.Query().Get("client_id")
	redirectUriRequest := r.URL.Query().Get("redirect_uri")
	c := &Client{Id: clientId}
	err, redirectUri := c.check()

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if authCodeConfig.ValidateRedirectURI == true {
		if len(redirectUriRequest) > 0 && !strings.HasPrefix(redirectUriRequest, redirectUri) {
			err = errors.New("access denied")
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if len(redirectUriRequest) > 0 {
		redirectUri = redirectUriRequest
	}

	ClearSession(w)

	http.Redirect(w, r, redirectUri, 302)
}
