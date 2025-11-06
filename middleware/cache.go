package middleware

import (
	"hifi/types"
	"sync"
)

var (
	songMap = make(map[string]types.SubsonicSong)
	songMu  sync.RWMutex

	albumYearMap = make(map[string]string)
	albumYearMu  sync.RWMutex

	artistCache = make(map[string]map[int]types.SubsonicArtist)
	artistMu    sync.RWMutex

	artistInfoCache = make(map[string]types.SubsonicArtistInfo)
	artistInfoMu    sync.RWMutex
)
