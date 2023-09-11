package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// WriteToConsole middleware learning
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hit Page")
		next.ServeHTTP(writer, request)

	})
}

// NoSurf adds CSRF Token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad save and load Session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
