package middleware

import (
	"encoding/json"
	"hifi/config"
	"hifi/types"
	"io"
	"log/slog"
	"net/http"
)

func extractIDs(t *types.TidalNew, moduleIndex int) []int {
	var ids []int

	for index, row := range t.Rows {
		if index != moduleIndex {
			continue
		}

		for _, module := range row.Modules {
			for _, item := range module.PagedList.Items {
				id := item.ID
				ids = append(ids, id)

				go func(id int, title, cover string) {
					slog.Info("Tidal item",
						"title", title,
						"id", id,
						"cover", cover,
					)
				}(id, item.Title, item.Cover)
			}
		}
	}

	return ids
}

func GetNew() []int {

	var tidalNew types.TidalNew

	tidalURL := QueryBuild(config.TidalHost, "/v1/pages/explore_new_music")

	q := tidalURL.Query()
	q.Set("countryCode", "US")
	q.Set("deviceType", "BROWSER")
	tidalURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+TidalAuth())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to send request to Tidal", "error", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read Tidal response", "error", err)
		return nil
	}

	if err := json.Unmarshal(body, &tidalNew); err != nil {
		slog.Error("failed to parse Tidal response", "error", err)
		return nil
	}

	return extractIDs(&tidalNew, 2)
}

func GetTop() []int {
	var tidalNew types.TidalNew

	tidalURL := QueryBuild(config.TidalHost, "/v1/pages/explore_top_music")

	q := tidalURL.Query()
	q.Set("countryCode", "US")
	q.Set("deviceType", "BROWSER")
	tidalURL.RawQuery = q.Encode()

	req, _ := http.NewRequest(config.MethodGet, tidalURL.String(), nil)
	req.Header.Set("Authorization", "Bearer "+TidalAuth())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to send request to Tidal (top)", "error", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read Tidal response (top)", "error", err)
		return nil
	}

	if err := json.Unmarshal(body, &tidalNew); err != nil {
		slog.Error("failed to parse Tidal response (top)", "error", err)
		return nil
	}

	return extractIDs(&tidalNew, 2)
}

func GetNewAndTop() []int {
	ch := make(chan []int, 2)

	go func() {
		ch <- GetNew()
	}()

	go func() {
		ch <- GetTop()
	}()

	var all []int
	for range 2 {
		ids := <-ch
		all = append(all, ids...)
	}
	close(ch)

	return all
}
