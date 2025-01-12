package databases

import (
	"context"

	"github.com/wuhewuhe/bybit.go.api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository interface {
	Create(ctx context.Context, dataBaseName, collectionName string, data interface{}) error
}

type mongoDB struct {
	db *mongo.Client
}

func NewMongoDB(db *mongo.Client) MongoDBRepository {
	return &mongoDB{db}
}

func (m *mongoDB) Create(ctx context.Context, dataBaseName, collectionName string, data interface{}) error {
	collection := m.db.Database(dataBaseName).Collection(collectionName)
	resdata := data.(models.CoinInfoResult)
	for _, d := range resdata.Rows {
		_, err := collection.InsertOne(context.TODO(), d)

		if err != nil {
			return err
		}
	}
	return nil
}
