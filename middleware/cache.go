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

	artistMapCache = make(map[int]types.SubsonicArtist)
	artistMu       sync.RWMutex
)
