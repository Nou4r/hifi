package middleware

import (
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/types"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func getArtists(search string, user string, w http.ResponseWriter) {

	var tidalArtist types.TidalArtistResponse

	if search != "" {
		queryMu.Lock()
		query[user] = search
		queryMu.Unlock()
	}

	queryMu.RLock()
	qu := query[user]
	queryMu.RUnlock()

	resultCh := make(chan *types.SubsonicWrapper)
	errorCh := make(chan error)

	go func() {
		tidalURL := &url.URL{
			Scheme: config.Scheme,
			Host:   config.TidalHost,
			Path:   "/v1/search/artists",
		}
		q := tidalURL.Query()
		q.Set("query", qu)
		q.Set("limit", "10000") // Max limit = 10K
		q.Set("offset", "0")
		q.Set("countryCode", "US")
		tidalURL.RawQuery = q.Encode()

		req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
		req.Header.Set("Authorization", "Bearer "+TidalAuth())

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			errorCh <- fmt.Errorf("tidal error: %v", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorCh <- fmt.Errorf("failed to read response: %v", err)
			return
		}

		if err := json.Unmarshal(body, &tidalArtist); err != nil {
			errorCh <- fmt.Errorf("parse error: %v", err)
			return
		}

		sub := buildSubsonicResponse(tidalArtist)
		resultCh <- sub
	}()

	select {
	case sub := <-resultCh:
		if sub == nil {
			http.Error(w, "no result", http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(sub)
	case err := <-errorCh:
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}

func buildSubsonicResponse(tidalArtist types.TidalArtistResponse) *types.SubsonicWrapper {
	sub := types.MetaBanner()
	sub.Subsonic.Artists = &types.SubsonicArtists{}

	artists := make([]types.SubsonicArtist, 0, len(tidalArtist.Items))
	for _, a := range tidalArtist.Items {
		artists = append(artists, types.SubsonicArtist{
			ID:       strconv.Itoa(a.ID),
			Name:     a.Name,
			CoverArt: a.Picture,
		})
	}

	sub.Subsonic.Artists.Index = []types.SubsonicArtistIndexItem{
		{
			Artist: artists,
		},
	}

	return &sub
}
