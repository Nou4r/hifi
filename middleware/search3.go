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

func search3(search string, user string, w http.ResponseWriter) {

	var tidalSearch types.TidalSearchResponse

	if search != "" {
		queryMu.Lock()
		query[user] = search
		queryMu.Unlock()
	}

	queryMu.RLock()
	qu := query[user]
	queryMu.RUnlock()

	// Tidal search URL
	tidalURL := &url.URL{
		Scheme: config.Scheme,
		Host:   config.TidalHost,
		Path:   "/v1/search/tracks",
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
		http.Error(w, fmt.Sprintf("tidal error: %v", err), http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Tidal returned %s", resp.Status), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read Tidal response", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &tidalSearch); err != nil {
		http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
		return
	}

	sub := types.MetaBanner()
	sub.Subsonic.SearchResult3 = &types.SubsonicSearchResult{}

	artistMap := make(map[int]bool)
	albumMap := make(map[int]bool)
	albumDurations := make(map[int]int)

	for _, item := range tidalSearch.Items {

		albumID := fmt.Sprint(item.Album.ID)
		songID := fmt.Sprint(item.ID)

		coverUUID := item.Album.Cover

		coverMu.Lock()
		coverMap[albumID] = coverUUID
		coverMap[songID] = coverUUID
		coverMu.Unlock()

		albumDurations[item.Album.ID] += item.Duration

		// Artist
		if !artistMap[item.Artist.ID] {
			sub.Subsonic.SearchResult3.Artist = append(sub.Subsonic.SearchResult3.Artist, types.SubsonicArtist{
				ID:       fmt.Sprint(item.Artist.ID),
				Name:     item.Artist.Name,
				CoverArt: item.Artist.Picture,
			})
			artistMap[item.Artist.ID] = true
		}

		// Album
		if !albumMap[item.Album.ID] {
			sub.Subsonic.SearchResult3.Album = append(sub.Subsonic.SearchResult3.Album, types.SubsonicAlbum{
				ID:       fmt.Sprint(item.Album.ID),
				Name:     item.Album.Title,
				Artist:   item.Artist.Name,
				CoverArt: item.Album.Cover,
				Year:     item.StreamStartDate[0:4],
				IsDir:    true,
			})
			albumMap[item.Album.ID] = true

		}

		// Song
		song := types.SubsonicSong{
			ID:          fmt.Sprint(item.ID),
			Title:       item.Title,
			Album:       item.Album.Title,
			Artist:      item.Artist.Name,
			Duration:    item.Duration,
			CoverArt:    item.Album.Cover,
			Type:        "music",
			IsVideo:     false,
			ContentType: "audio/flac",
			Suffix:      "flac",
			ArtistID:    fmt.Sprint(item.Artist.ID),
			AlbumID:     fmt.Sprint(item.Album.ID),
		}

		sub.Subsonic.SearchResult3.Song = append(sub.Subsonic.SearchResult3.Song, song)

		songMu.Lock()
		songMap[songID] = song
		songMu.Unlock()

	}

	for i, alb := range sub.Subsonic.SearchResult3.Album {
		id64, _ := strconv.ParseInt(alb.ID, 10, 64)
		alb.Duration = albumDurations[int(id64)]
		sub.Subsonic.SearchResult3.Album[i] = alb
	}

	// Write response
	json.NewEncoder(w).Encode(sub)

}
