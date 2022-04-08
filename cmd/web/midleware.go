package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// // WriteToConsole midleware custom anonymus fucntion example
// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("midleware test")
// 		next.ServeHTTP(w, r)
// 	})
// 	//could do any custom, the only condition is to return handler
// }

// NoSurf provide CSRF protection for all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad ensure to load session on each request
func SessioLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
