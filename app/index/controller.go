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
	log.Println("Sync Countries")
	result, err := x.Service.SyncCountries(ctx)
	if err != nil {
		return err
	}
	log.Println("Insert", len(result.InsertedIDs))
	log.Println("Sync States")
	if result, err = x.Service.SyncStates(ctx); err != nil {
		return err
	}
	log.Println("Insert", len(result.InsertedIDs))
	log.Println("Sync Cities")
	if result, err = x.Service.SyncCities(ctx); err != nil {
		return err
	}
	log.Println("Insert", len(result.InsertedIDs))
	return gin.H{"sync": time.Now()}
}
