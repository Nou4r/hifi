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

	for _, item := range tidalArtistAlbums.Items {
		if item.ModuleId == "ARTIST_ALBUMS" {
			for _, albumItem := range item.Items {
				data := albumItem.Data
				id := fmt.Sprint(data.ID)
				title := data.Title
				duration := data.Duration
				year := data.ReleaseDate[:4]
				cover := data.Cover
				fmt.Println("Album:", id, title, duration, year, cover)
			}
		}
	}

	artistData := tidalArtistAlbums.Item.Data

	artist := types.SubsonicArtistWithAlbums{
		ID:         fmt.Sprint(artistData.ID),
		Name:       artistData.Name,
		CoverArt:   firstNonEmpty(artistData.Picture, artistData.SelectedAlbumCoverFallback),
		AlbumCount: len(tidalArtistAlbums.Items[1].Items),
	}

	sub := types.MetaBanner()
	sub.Subsonic.Artist = &artist

	_ = json.NewEncoder(w).Encode(sub)
}
