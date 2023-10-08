package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/database"
)

func SetDatabaseConnection(conn *database.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("conn", conn)
		c.Next()
	}
}

func CastToDatabaseConnection(conn any) *database.Connection {
	return conn.(*database.Connection)
}
