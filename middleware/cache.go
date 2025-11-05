package middleware

import (
	"hifi/types"
	"sync"
)

var (
	songMap = make(map[string]types.SubsonicSong)
	songMu  sync.RWMutex
)
