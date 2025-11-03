package middleware

import "net/http"

func getArtistInfo(user string, id string, w http.ResponseWriter) {
	w.Write([]byte("getArtistInfo: " + id + " for user: " + user))
}
