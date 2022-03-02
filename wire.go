//go:build wireinject
// +build wireinject

package main

import (
	"countries-states-cities-mongo/app"
	"countries-states-cities-mongo/bootstrap"
	"countries-states-cities-mongo/common"
	"github.com/google/wire"
)

func App(value *common.Values) (*app.App, error) {
	wire.Build(
		wire.Struct(new(common.Inject), "*"),
		bootstrap.Provides,
		app.Provides,
	)
	return &app.App{}, nil
}
