package index

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Controller struct {
	Service *Service
}

func (x *Controller) Run(c *gin.Context) interface{} {
	ctx := c.Request.Context()
	log.Println("开始同步国家/地区信息")
	if err := x.Service.SyncCountries(ctx); err != nil {
		return err
	}
	log.Println("同步国家/地区信息已完成")
	log.Println("开始同步省/州信息")
	if err := x.Service.SyncStates(ctx); err != nil {
		return err
	}
	log.Println("同步省/州信息已完成")
	log.Println("开始同步城市信息")
	if err := x.Service.SyncCities(ctx); err != nil {
		return err
	}
	log.Println("同步城市信息已完成")
	return gin.H{"sync": time.Now()}
}
