package middlewares

import (
	"strings"

	"github.com/STREAM-BUSTER/stream-buster/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// defaultAllowedOrigins is the local Vite dev server, used when
// CORS_ALLOWED_ORIGINS is unset.
var defaultAllowedOrigins = []string{"http://localhost:5173"}

func CORS() gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:     allowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	}

	return cors.New(config)
}

// allowedOrigins reads CORS_ALLOWED_ORIGINS: a comma-separated list of exact
// origins, e.g. "https://streambuster.xyz,https://dev.streambuster.xyz". These
// routes set credentialed cookies, so a wildcard is not usable — every
// deployment names its own front ends.
func allowedOrigins() []string {
	raw := utils.GetEnvVariable("CORS_ALLOWED_ORIGINS")
	if raw == "" {
		return defaultAllowedOrigins
	}

	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))
	for _, origin := range parts {
		if origin = strings.TrimSpace(origin); origin != "" {
			origins = append(origins, origin)
		}
	}

	if len(origins) == 0 {
		return defaultAllowedOrigins
	}
	return origins
}
