package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/settings"
)

func SetSettings(s *settings.Settings) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("s", s)
		c.Next()
	}
}

func CastToSettings(s any) *settings.Settings {
	return s.(*settings.Settings)
}
