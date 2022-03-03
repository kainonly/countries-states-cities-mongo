package app

import (
	"countries-states-cities-mongo/app/index"
	"countries-states-cities-mongo/common"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/weplanx/go/route"
)

var Provides = wire.NewSet(
	index.Provides,
	New,
)

func New(
	values *common.Values,
	index *index.Controller,
) *gin.Engine {
	r := globalMiddleware(gin.New(), values)
	r.POST("/event-invoke", route.Use(index.Run))
	return r
}
