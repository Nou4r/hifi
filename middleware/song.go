package middleware

import (
	"encoding/json"
	"hifi/types"
	"net/http"
)

func song(id string, w http.ResponseWriter) {
	var songMap = make(map[string]types.SubsonicSong)

	songMu.RLock()
	song := songMap[id]
	songMu.RUnlock()

	sub := types.MetaBanner()
	sub.Subsonic.Song = &song

	json.NewEncoder(w).Encode(sub)
}
