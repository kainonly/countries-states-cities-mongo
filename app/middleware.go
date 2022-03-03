package app

import (
	"countries-states-cities-mongo/common"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func globalMiddleware(r *gin.Engine, values *common.Values) *gin.Engine {
	r.SetTrustedProxies(values.TrustedProxies)
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(gin.CustomRecovery(catchError))
	return r
}

func catchError(c *gin.Context, err interface{}) {
	c.AbortWithStatusJSON(500, gin.H{
		"message": err,
	})
}
