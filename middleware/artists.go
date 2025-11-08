package middleware

import (
	"encoding/json"
	"hifi/types"
	"maps"
	"net/http"
	"slices"
)

func getArtists(user string, w http.ResponseWriter) {
	sub := types.MetaBanner()
	sub.Subsonic.Artists = &types.SubsonicArtists{}

	artistsMu.RLock()
	userArtists := artistsCache[user]
	artistsMu.RUnlock()

	artists := slices.Collect(maps.Values(userArtists))

	sub.Subsonic.Artists.Index = []types.SubsonicArtistIndexItem{
		{Artist: artists},
	}

	_ = json.NewEncoder(w).Encode(sub)
}
