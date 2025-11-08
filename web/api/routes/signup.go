package routes

import (
	"api/middleware"
	"api/types"
	"net/http"
)

func Handle() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/signup", &types.Routes{User: middleware.SignupUser()})
	return mux
}
