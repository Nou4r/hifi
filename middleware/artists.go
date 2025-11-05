package middleware

import (
	"encoding/json"
	"hifi/types"
	"net/http"
)

func getArtists(user string, w http.ResponseWriter) {
	sub := types.MetaBanner()
	sub.Subsonic.Artists = &types.SubsonicArtists{}

	artistMu.RLock()
	userArtists := artistCache[user]
	artistMu.RUnlock()

	artists := make([]types.SubsonicArtist, 0, len(userArtists))
	for _, a := range userArtists {
		artists = append(artists, a)
	}

	sub.Subsonic.Artists.Index = []types.SubsonicArtistIndexItem{
		{Artist: artists},
	}

	json.NewEncoder(w).Encode(sub)
}
