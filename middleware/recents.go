package middleware

import (
	"fmt"
	"hifi/types"
	"log/slog"
)

func RecentAlbum() {

	var allAlbums []types.SubsonicAlbum

	user := Public
	ids := []string{
		"247415928",
		"463900363",
		"441821356",
		"462795478",
		"469227943",
		"466488538",
	}

	results := make(chan types.SubsonicAlbum, len(ids))

	for _, id := range ids {
		go func(albumID string) {
			album := fetchAndCacheAlbum(user, albumID)
			results <- album
		}(id)
	}

	for range ids {
		album := <-results
		fmt.Printf("[Album cached] %s — %s\n", album.ID, album.Title)
		allAlbums = append(allAlbums, album)
	}
	close(results)

	useralbumMu.Lock()
	if userAlbumCache[user] == nil {
		userAlbumCache[user] = make(map[string]types.SubsonicAlbum)
	}
	for _, album := range allAlbums {
		userAlbumCache[user][album.ID] = album
	}
	useralbumMu.Unlock()

	logger := slog.Default()
	logger.Info("Album refresh completed",
		slog.Int("total_albums_cached", len(allAlbums)),
		slog.String("user", user),
	)
}
