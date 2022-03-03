package index

import (
	"context"
	"countries-states-cities-mongo/common"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	*common.Inject
}

func (x *Service) SyncCountries(ctx context.Context) (result *mongo.InsertManyResult, err error) {
	var res *resty.Response
	if res, err = x.Client.R().SetContext(ctx).
		Get("countries.json"); err != nil {
		return
	}
	var body []map[string]interface{}
	if err = jsoniter.Unmarshal(res.Body(), &body); err != nil {
		return
	}
	data := make([]interface{}, len(body))
	for k, v := range body {
		delete(v, "id")
		data[k] = v
	}
	if err = x.Db.Collection("countries").
		Drop(ctx); err != nil {
		return
	}
	if result, err = x.Db.Collection("countries").
		InsertMany(ctx, data); err != nil {
		return
	}
	if _, err = x.Db.Collection("countries").Indexes().
		CreateMany(ctx, []mongo.IndexModel{
			{
				Keys:    bson.M{"iso2": 1},
				Options: options.Index().SetName("uk_iso2").SetUnique(true),
			},
			{
				Keys:    bson.M{"iso3": 1},
				Options: options.Index().SetName("uk_iso3").SetUnique(true),
			},
			{
				Keys:    bson.M{"numeric_code": 1},
				Options: options.Index().SetName("uk_numeric_code").SetUnique(true),
			},
		}); err != nil {
		return
	}
	return
}

func (x *Service) SyncStates(ctx context.Context) (result *mongo.InsertManyResult, err error) {
	var res *resty.Response
	if res, err = x.Client.R().SetContext(ctx).
		Get("states.json"); err != nil {
		return
	}
	var body []map[string]interface{}
	if err = jsoniter.Unmarshal(res.Body(), &body); err != nil {
		return
	}
	data := make([]interface{}, len(body))
	for k, v := range body {
		delete(v, "id")
		delete(v, "country_id")
		data[k] = v
	}
	if err = x.Db.Collection("states").
		Drop(ctx); err != nil {
		return
	}
	if result, err = x.Db.Collection("states").
		InsertMany(ctx, data); err != nil {
		return
	}
	if _, err = x.Db.Collection("states").Indexes().
		CreateMany(ctx, []mongo.IndexModel{
			{
				Keys:    bson.M{"country_code": 1},
				Options: options.Index().SetName("idx_country_code"),
			},
			{
				Keys:    bson.M{"state_code": 1},
				Options: options.Index().SetName("idx_state_code"),
			},
		}); err != nil {
		return
	}
	return
}

func (x *Service) SyncCities(ctx context.Context) (result *mongo.InsertManyResult, err error) {
	var res *resty.Response
	if res, err = x.Client.R().SetContext(ctx).
		Get("cities.json"); err != nil {
		return
	}
	var body []map[string]interface{}
	if err = jsoniter.Unmarshal(res.Body(), &body); err != nil {
		return
	}
	data := make([]interface{}, len(body))
	for k, v := range body {
		delete(v, "id")
		delete(v, "country_id")
		delete(v, "state_id")
		data[k] = v
	}
	if err = x.Db.Collection("cities").
		Drop(ctx); err != nil {
		return
	}
	if result, err = x.Db.Collection("cities").
		InsertMany(ctx, data); err != nil {
		return
	}
	if _, err = x.Db.Collection("cities").Indexes().
		CreateMany(ctx, []mongo.IndexModel{
			{
				Keys:    bson.M{"country_code": 1},
				Options: options.Index().SetName("idx_country_code"),
			},
			{
				Keys:    bson.M{"state_code": 1},
				Options: options.Index().SetName("idx_state_code"),
			},
			{
				Keys:    bson.M{"name": 1},
				Options: options.Index().SetName("idx_name"),
			},
		}); err != nil {
		return
	}
	return
}
