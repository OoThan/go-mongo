package repository

import (
	"context"
	"go-mongo/model/access_control/model"
	"go-mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context

type accessRepoImpl struct {
	connection *mongo.Database
}

func (r *accessRepoImpl) Count() (int64, error) {
	count, err := r.connection.Collection("access_control").CountDocuments(ctx, bson.M{})
	if !utils.GlobalErrorDatabaseException(err) {
		return 0, err
	}
	return count, nil
}

func (r *accessRepoImpl) FindAll(limit int64, page int64) ([]*model.AccessControl, error) {
	var (
		acl           *model.AccessControl
		acls          []*model.AccessControl
		cur           *mongo.Cursor
		filterOptions = options.Find()
	)
	filterOptions.SetSkip(page)
	filterOptions.SetLimit(limit)
	cur, err := r.connection.Collection("access_control").Find(ctx, bson.M{}, filterOptions)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	for cur.Next(ctx) {
		errDecode := cur.Decode(&acl)
		if !utils.GlobalErrorDatabaseException(errDecode) {
			return nil, errDecode
		}
		acls = append(acls, acl)
	}
	return acls, nil
}

func (r *accessRepoImpl) FindById(id string) (*model.DetailAccessControl, error) {
	var (
		objectId, _ = primitive.ObjectIDFromHex(id)
		acl         *model.DetailAccessControl
		filter      = bson.M{"_id": objectId}
	)
	err := r.connection.Collection("access_control").FindOne(ctx, filter).Decode(&acl)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return acl, nil
}

func (r *accessRepoImpl) Save(payload *model.CreateAccessControl) (*model.CreateAccessControl, error) {
	_, err := r.connection.Collection("access_control").InsertOne(ctx, payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return payload, nil
}

func NewAccessControlRepository(Connection *mongo.Database) Repository {
	return &accessRepoImpl{
		connection: Connection,
	}
}
