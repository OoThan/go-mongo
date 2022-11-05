package repository

import (
	"context"
	"go-mongo/model/role/model"
	"go-mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context

type roleRepoImpl struct {
	connection *mongo.Database
}

func (r *roleRepoImpl) Count() (int64, error) {
	count, err := r.connection.Collection("role").CountDocuments(ctx, bson.M{}, nil)
	if !utils.GlobalErrorDatabaseException(err) {
		return 0, err
	}
	return count, nil
}

func (r *roleRepoImpl) FindAll(limit int64, page int64) ([]*model.Role, error) {
	var (
		role          *model.Role
		roles         []*model.Role
		csr           *mongo.Cursor
		filterOptions = options.Find()
	)

	filterOptions.SetSkip(page)
	filterOptions.SetLimit(limit)

	csr, err := r.connection.Collection("role").Find(ctx, bson.M{}, filterOptions)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}

	for csr.Next(ctx) {
		err = csr.Decode(&role)
		if !utils.GlobalErrorDatabaseException(err) {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *roleRepoImpl) FindById(id string) (*model.Role, error) {
	var (
		role      *model.Role
		roleId, _ = primitive.ObjectIDFromHex(id)
		filter    = bson.M{"_id": roleId}
	)
	err := r.connection.Collection("role").FindOne(ctx, filter).Decode(&role)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return role, nil
}

func (r *roleRepoImpl) Save(payload *model.CreateRole) error {
	_, err := r.connection.Collection("role").InsertOne(ctx, payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func (r *roleRepoImpl) Update(id string, payload *model.UpdateRole) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	updatedFiled := bson.M{
		"$set": bson.M{
			"role": payload.RoleName,
		},
	}
	_, err := r.connection.Collection("role").UpdateOne(ctx, filter, updatedFiled)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func (r *roleRepoImpl) Delete(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	_, err := r.connection.Collection("role").DeleteOne(ctx, filter)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func NewRoleRepository(connection *mongo.Database) Repository {
	return &roleRepoImpl{
		connection: connection,
	}
}
