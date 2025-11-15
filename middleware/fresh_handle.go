package middleware

import (
	"encoding/json"
	"hifi/config"
	"hifi/types"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

var (
	freshCache    []types.ExploreItem
	freshCacheMu  sync.RWMutex
	freshCacheExp time.Time
)

func getFreshCachedItems() []types.ExploreItem {
	now := time.Now()

	freshCacheMu.RLock()
	if freshCache != nil && now.Before(freshCacheExp) {
		items := freshCache
		freshCacheMu.RUnlock()
		return items
	}
	freshCacheMu.RUnlock()

	items := GetNewAndTopItems()

	freshCacheMu.Lock()
	freshCache = items
	freshCacheExp = now.Add(1 * time.Minute)
	freshCacheMu.Unlock()

	return items
}

func FreshHandler(w http.ResponseWriter, r *http.Request) {
	items := getFreshCachedItems()

	if len(items) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set(config.HeaderCacheControl, "public, max-age=86400")
	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)

	if err := json.NewEncoder(w).Encode(items); err != nil {
		slog.Error("failed to encode items", "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
