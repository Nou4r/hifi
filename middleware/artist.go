package middleware

import (
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/types"
	"io"
	"net/http"
	"net/url"
)

func getArtist(id string, w http.ResponseWriter) {
	var tidalArtistAlbums types.TidalArtistAlbumsResponse

	tidalURL := &url.URL{
		Scheme: config.Scheme,
		Host:   config.TidalHost,
		Path:   fmt.Sprintf("/v2/artist/%s", id),
	}
	q := tidalURL.Query()
	q.Set("locale", "en_US")
	q.Set("countryCode", "US")
	q.Set("deviceType", "BROWSER")
	q.Set("platform", "WEB")
	tidalURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+TidalAuth())
	req.Header.Set("x-tidal-client-version", "2025.11.3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Tidal error: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read Tidal response", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &tidalArtistAlbums); err != nil {
		http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Number of albumItem.Items: %d\n", len(tidalArtistAlbums.Items))

	artistData := tidalArtistAlbums.Item.Data

	artist := types.SubsonicArtistWithAlbums{
		ID:         fmt.Sprint(artistData.ID),
		Name:       artistData.Name,
		CoverArt:   firstNonEmpty(artistData.Picture, artistData.SelectedAlbumCoverFallback),
		AlbumCount: len(tidalArtistAlbums.Items),
	}

	sub := types.MetaBanner()
	sub.Subsonic.Artist = &artist

	_ = json.NewEncoder(w).Encode(sub)
}
