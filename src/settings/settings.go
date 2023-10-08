package settings

import (
	"github.com/gin-gonic/gin"
	"github.com/miraliumre/clarity-api/env"
	"go.uber.org/zap"
)

type Settings struct {
	BindAddr        string
	EnableDebug     bool
	EnableDocuments bool
	EnableCerts     bool
	RedisAddr       string
	RedisPassword   string
	CertsDir        string
	Logger          *zap.Logger
}

func Init() *Settings {
	s := &Settings{}

	s.BindAddr = env.ReadStr("CLARITY_BIND_ADDR", "0.0.0.0:8080")

	s.EnableDebug = env.ReadBool("CLARITY_DEBUG_MODE", false)
	s.EnableDocuments = env.ReadBool("CLARITY_DOCUMENTS", false)
	s.EnableCerts = env.ReadBool("CLARITY_CERTS", false)

	s.RedisAddr = env.ReadStr("CLARITY_REDIS_ADDR", "localhost:6379")
	s.RedisPassword = env.ReadStr("CLARITY_REDIS_PASSWORD", "")

	s.CertsDir = env.ReadStr("CLARITY_CERTS_DIR", "/usr/local/clarity/certs")

	if s.EnableDebug {
		s.Logger, _ = zap.NewDevelopment()
	} else {
		s.Logger, _ = zap.NewProduction()
		gin.SetMode(gin.ReleaseMode)
	}

	return s
}
