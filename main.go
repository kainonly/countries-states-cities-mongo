package main

import (
	"context"
	"countries-states-cities-mongo/app"
	"countries-states-cities-mongo/bootstrap"
	"countries-states-cities-mongo/common"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func Invoke(ctx context.Context) (err error) {
	var v *common.Values
	if v, err = bootstrap.SetValues(); err != nil {
		return err
	}
	var x *app.App
	if x, err = App(v); err != nil {
		return err
	}
	if err = x.Run(ctx); err != nil {
		return err
	}
	return nil
}

func main() {
	cloudfunction.Start(Invoke)
}
