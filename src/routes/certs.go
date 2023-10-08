package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/certs"
	"github.com/miraliumre/clarity-api/middleware"
	"github.com/miraliumre/clarity-api/strings"
)

func Certs(c *gin.Context) {
	val, exists := c.Get("s")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			strings.ErrorKey: strings.ErrorMessageInternalServerError,
		})
		return
	}

	s := middleware.CastToSettings(val)
	if !s.EnableCerts {
		c.JSON(http.StatusNotImplemented, gin.H{
			strings.ErrorKey: strings.ErrorMessageFeatureNotEnabled,
		})
		return
	}

	contents, err := certs.ReadAll(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			strings.ErrorKey: strings.ErrorMessageInternalServerError,
		})
		s.Logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		strings.CertsKey: contents,
	})
}
