package routes

import (
	"api/middleware"
	"net/http"
)

func Handle() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", middleware.SigninUser)

	mux.HandleFunc("/signup", middleware.SignupUser)

	return mux
}
