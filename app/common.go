package app

import (
	"countries-states-cities-mongo/common"
	"github.com/google/wire"
)

var Provides = wire.NewSet(New)

func New(i *common.Inject) *App {
	return &App{
		Inject: i,
	}
}
