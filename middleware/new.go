package middleware

import (
	"encoding/json"
	"fmt"
	"hifi/config"
	"hifi/types"
	"io"
	"log/slog"
	"net/http"
)

func PrintTidalItems(t *types.TidalNew) {
	for _, row := range t.Rows {
		for _, module := range row.Modules {
			for _, item := range module.PagedList.Items {
				fmt.Printf("Title: %s, ID: %s, Cover: %s\n",
					item.Title, item.ID, item.Cover)
			}
		}
	}
}

func GetNew() {

	var tidalNew types.TidalNew

	// Build Tidal request
	tidalURL := QueryBuild(config.TidalHost, "/v1/pages/explore_new_music")

	q := tidalURL.Query()
	q.Set("countryCode", "US")
	tidalURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+TidalAuth())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to send request to Tidal", "error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		slog.Error("failed to read Tidal response", "error", err)
		return
	}

	if err := json.Unmarshal(body, &tidalNew); err != nil {
		slog.Error("failed to parse Tidal response", "error", err)
		return
	}

	PrintTidalItems(&tidalNew)
}
