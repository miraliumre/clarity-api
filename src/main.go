package main

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/database"
	"github.com/miraliumre/clarity-api/middleware"
	"github.com/miraliumre/clarity-api/routes"
	"github.com/miraliumre/clarity-api/settings"
)

func main() {
	s := settings.Init()
	r := gin.New()

	r.Use(ginzap.Ginzap(s.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(s.Logger, true))
	r.Use(middleware.SetSettings(s))

	if s.EnableDocuments {
		conn, err := database.Init(s)
		if err != nil {
			panic(err)
		}
		r.Use(middleware.SetDatabaseConnection(conn))
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/discovery", routes.Discovery)
		v1.GET("/documents/:hash", routes.Documents)
		v1.GET("/certs", routes.Certs)
	}

	r.NoRoute(routes.NoRoute)
	r.Run(s.BindAddr)
}
