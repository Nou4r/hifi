package middleware

import (
	"encoding/json"
	"hifi/types"
	"net/http"
)

func getArtist(id string, user string, w http.ResponseWriter) {
	sub := types.MetaBanner()
	sub.Subsonic.Artists = &types.SubsonicArtists{}

	artistMu.RLock()
	userArtists := artistCache[user]
	artistMu.RUnlock()

	var artists []types.SubsonicArtist
	for _, a := range userArtists {
		artists = append(artists, a)

	}

	sub.Subsonic.Artists.Index = []types.SubsonicArtistIndexItem{
		{Artist: artists},
	}

	_ = json.NewEncoder(w).Encode(sub)
}
