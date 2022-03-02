package common

import "go.mongodb.org/mongo-driver/mongo"

type Inject struct {
	Values      *Values
	MongoClient *mongo.Client
	Db          *mongo.Database
}

type Values struct {
	Database Database `envPrefix:"DATABASE_"`
}

type Database struct {
	Uri    string `env:"URI"`
	DbName string `env:"DBNAME"`
}
