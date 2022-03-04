package index

import (
	"context"
	"countries-states-cities-mongo/common"
	"countries-states-cities-mongo/model"
	"encoding/csv"
	jsoniter "github.com/json-iterator/go"
	"github.com/panjf2000/ants/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Service struct {
	*common.Inject
}

func (x *Service) SyncCountries(ctx context.Context) (err error) {
	client := http.DefaultClient
	url := baseURL("/csv/countries.csv")
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}
	var resp *http.Response
	if resp, err = client.Do(req.WithContext(ctx)); err != nil {
		return
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	first := true
	var countries []interface{}
	for {
		var record []string
		if record, err = r.Read(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
		if first {
			first = false
			continue
		}
		var timezones []map[string]interface{}
		tzRaw := strings.ReplaceAll(record[14], `\/`, `/`)
		var vRegex *regexp.Regexp
		if vRegex, err = regexp.Compile(`:'([^,|}]+)'([,|}])`); err != nil {
			return
		}
		var kRegex *regexp.Regexp
		if kRegex, err = regexp.Compile(`([a-zA-Z]+):`); err != nil {
			return
		}
		toJSON := kRegex.ReplaceAll(
			vRegex.ReplaceAll([]byte(tzRaw), []byte(`:"$1"$2`)),
			[]byte(`"$1":`),
		)
		if err = jsoniter.Unmarshal(toJSON, &timezones); err != nil {
			return
		}
		latitude := float64(0)
		if record[15] != "" {
			if latitude, err = strconv.ParseFloat(record[15], 64); err != nil {
				return
			}
		}
		longitude := float64(0)
		if record[16] != "" {
			if longitude, err = strconv.ParseFloat(record[16], 64); err != nil {
				return
			}
		}
		countries = append(countries, model.Country{
			Name:           record[1],
			Iso3:           record[2],
			Iso2:           record[3],
			NumberCode:     record[4],
			PhoneCode:      record[5],
			Capital:        record[6],
			Currency:       record[7],
			CurrencyName:   record[8],
			CurrencySymbol: record[9],
			Tld:            record[10],
			Native:         record[11],
			Region:         record[12],
			Subregion:      record[13],
			Timezones:      timezones,
			Latitude:       latitude,
			Longitude:      longitude,
			Emoji:          record[17],
			EmojiU:         record[18],
		})
	}
	if err = x.Db.Collection("countries").
		Drop(ctx); err != nil {
		return
	}
	if _, err = x.Db.Collection("countries").
		InsertMany(ctx, countries); err != nil {
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
				Keys:    bson.M{"number_code": 1},
				Options: options.Index().SetName("uk_number_code").SetUnique(true),
			},
		}); err != nil {
		return
	}
	return
}

func (x *Service) SyncStates(ctx context.Context) (err error) {
	client := http.DefaultClient
	url := baseURL("/csv/states.csv")
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}
	var resp *http.Response
	if resp, err = client.Do(req.WithContext(ctx)); err != nil {
		return
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	first := true
	var states []interface{}
	for {
		var record []string
		if record, err = r.Read(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
		if first {
			first = false
			continue
		}
		latitude := float64(0)
		if record[7] != "" {
			if latitude, err = strconv.ParseFloat(record[7], 64); err != nil {
				return
			}
		}
		longitude := float64(0)
		if record[8] != "" {
			if longitude, err = strconv.ParseFloat(record[8], 64); err != nil {
				return
			}
		}
		states = append(states, model.State{
			Name:        record[1],
			CountryCode: record[3],
			StateCode:   record[5],
			Type:        record[6],
			Latitude:    latitude,
			Longitude:   longitude,
		})
	}
	if err = x.Db.Collection("states").
		Drop(ctx); err != nil {
		return
	}
	if _, err = x.Db.Collection("states").
		InsertMany(ctx, states); err != nil {
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

func (x *Service) SyncCities(ctx context.Context) (err error) {
	client := http.DefaultClient
	url := baseURL("/csv/cities.csv")
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}
	var resp *http.Response
	if resp, err = client.Do(req.WithContext(ctx)); err != nil {
		return
	}
	defer resp.Body.Close()
	if err = x.Db.Collection("cities").
		Drop(ctx); err != nil {
		return
	}
	var wg sync.WaitGroup
	var p *ants.PoolWithFunc
	if p, err = ants.NewPoolWithFunc(100, func(i interface{}) {
		if v, ok := i.([]interface{}); ok {
			if _, err := x.Db.Collection("cities").
				InsertMany(ctx, v); err != nil {
				log.Fatalln(err)
			}
		}
		wg.Done()
	}); err != nil {
		return
	}
	defer p.Release()
	r := csv.NewReader(resp.Body)
	first := true
	var cities []interface{}
	for {
		var record []string
		if record, err = r.Read(); err != nil {
			if err == io.EOF {
				wg.Add(1)
				_ = p.Invoke(cities)
				break
			} else {
				log.Fatalln(err)
			}
		}
		if first {
			first = false
			continue
		}
		latitude := float64(0)
		if record[8] != "" {
			if latitude, err = strconv.ParseFloat(record[8], 64); err != nil {
				return
			}
		}
		longitude := float64(0)
		if record[9] != "" {
			if longitude, err = strconv.ParseFloat(record[9], 64); err != nil {
				return
			}
		}
		cities = append(cities, model.City{
			Name:        record[1],
			CountryCode: record[6],
			StateCode:   record[3],
			Latitude:    latitude,
			Longitude:   longitude,
		})
		if len(cities) == 5000 {
			wg.Add(1)
			_ = p.Invoke(cities)
			cities = []interface{}{}
		}
	}
	wg.Wait()
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
