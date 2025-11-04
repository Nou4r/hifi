package middleware

import (
	"fmt"
	"net/http"
)

func getArtistInfo(user string, id string, w http.ResponseWriter) {
	fmt.Println("getArtistInfo called with id:", id, "for user:", user)
	fmt.Println()
}
