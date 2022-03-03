package common

import (
	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Inject struct {
	Values      *Values
	MongoClient *mongo.Client
	Db          *mongo.Database
	Client      *resty.Client
}

type Values struct {
	TrustedProxies []string `env:"TRUSTED_PROXIES"`
	Database       Database `envPrefix:"DATABASE_"`
}

type Database struct {
	Uri    string `env:"URI"`
	DbName string `env:"DBNAME"`
}
