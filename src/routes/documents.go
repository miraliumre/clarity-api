package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/database"
	"github.com/miraliumre/clarity-api/middleware"
	"github.com/miraliumre/clarity-api/strings"
	"github.com/redis/go-redis/v9"
)

func Documents(c *gin.Context) {
	val, exists := c.Get("s")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			strings.ErrorKey: strings.ErrorMessageInternalServerError,
		})
		return
	}

	s := middleware.CastToSettings(val)
	if !s.EnableDocuments {
		c.JSON(http.StatusNotImplemented, gin.H{
			strings.ErrorKey: strings.ErrorMessageFeatureNotEnabled,
		})
		return
	}

	val, exists = c.Get("conn")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			strings.ErrorKey: strings.ErrorMessageInternalServerError,
		})
		s.Logger.Error("database connection does not exist in context")
		return
	}

	hash := c.Param("hash")
	conn := middleware.CastToDatabaseConnection(val)
	doc, err := database.GetDocument(conn, hash)

	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			strings.ErrorKey: strings.ErrorMessageNotFound,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			strings.ErrorKey: strings.ErrorMessageInternalServerError,
		})
		s.Logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, doc)
}
