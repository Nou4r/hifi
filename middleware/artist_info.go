package middleware

import (
	"encoding/json"
	"hifi/types"
	"net/http"
)

func getArtistInfo(id string, user string, w http.ResponseWriter) {

	info := types.SubsonicArtistInfo{
		Biography:      "No biography available.",
		SmallImageURL:  "/coverArt?id=" + id + "&size=200",
		MediumImageURL: "/coverArt?id=" + id + "&size=450",
		LargeImageURL:  "/coverArt?id=" + id + "&size=500",
	}

	sub := types.MetaBanner()
	sub.Subsonic.ArtistInfo = &info

	json.NewEncoder(w).Encode(sub)
}
