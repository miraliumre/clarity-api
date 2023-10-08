package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/app"
	"github.com/miraliumre/clarity-api/middleware"
	"github.com/miraliumre/clarity-api/strings"
)

func Discovery(c *gin.Context) {
	val, exists := c.Get("s")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			strings.ErrorKey: strings.ErrorMessageInternalServerError,
		})
		return
	}

	s := middleware.CastToSettings(val)
	lookup := make([]string, 0)

	if s.EnableDocuments {
		lookup = append(lookup, strings.LookupTypeDocuments)
	}

	if s.EnableCerts {
		lookup = append(lookup, strings.LookupTypeCerts)
	}

	c.JSON(http.StatusOK, gin.H{
		"vendor":  app.Vendor,
		"app":     app.Name,
		"version": app.Version,
		"url":     app.Url,
		"lookup":  lookup,
	})
}
