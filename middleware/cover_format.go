package middleware

import (
	"hifi/config"
	"strings"
)

func NormalizeCover(id string) string {
	if id == "" {
		return ""
	}

	if strings.HasPrefix(id, "https://") {
		return id
	}

	return config.Scheme + "://" + config.TidalStaticHost + "/images/" + FormatCoverID(id)
}
