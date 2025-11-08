package routes

import (
	"api/middleware"
	"net/http"
)

func Handle() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", middleware.SignupUser)

	return mux
}
