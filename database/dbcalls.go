package database

import (
	"context"
	"urlShortenerMongo/constant"
	"urlShortenerMongo/types"

	"go.mongodb.org/mongo-driver/bson"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	res, err := instance.InsertOne(context.TODO(), data)
	return res.InsertedID, err
}

func (mgr *manager) GetUrlFromCode(code string, collectionName string) (resp types.UrlDao, err error) {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	err = instance.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	return resp, err
}
