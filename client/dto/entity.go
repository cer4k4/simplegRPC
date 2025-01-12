package dto

import "go.mongodb.org/mongo-driver/mongo"

type ServiceConfig struct {
	ZipkinEndpoint string
	ServerPort     string
	MongoClient    *mongo.Client
}
