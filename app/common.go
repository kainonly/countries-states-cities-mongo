package app

import (
	"countries-states-cities-mongo/common"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
)

var Provides = wire.NewSet(New)

func New(i *common.Inject) *App {
	return &App{
		Inject: i,
		Client: resty.New().
			SetBaseURL(`https://raw.githubusercontent.com/dr5hn/countries-states-cities-database/master`),
	}
}
