package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/strings"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		strings.ErrorKey: strings.ErrorMessageNotFound,
	})
}
