package repository

import (
	"context"
	"go-mongo/model/user/model"
	"go-mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context

type userRepoImpl struct {
	Connection *mongo.Database
}

func (u *userRepoImpl) FindAll(name string, limit int64, offset int64) ([]*model.Users, error) {
	var (
		user         *model.Users
		users        []*model.Users
		filerOptions = options.Find()
		csr          *mongo.Cursor
		errCsr       error
	)

	filerOptions.SetLimit(limit)
	filerOptions.SetSkip(offset)

	if name != "" {
		csr, errCsr = u.Connection.Collection("users").Find(ctx, bson.M{"name": name}, filerOptions)
		if !utils.GlobalErrorDatabaseException(errCsr) {
			return nil, errCsr
		}
	}

	for csr.Next(ctx) {
		err := csr.Decode(&user)
		if !utils.GlobalErrorDatabaseException(err) {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (u *userRepoImpl) Count() (int64, error) {
	count, err := u.Connection.Collection("users").CountDocuments(ctx, bson.M{}, nil)
	if !utils.GlobalErrorDatabaseException(err) {
		return 0, err
	}
	return count, nil
}

func (u *userRepoImpl) FindById(id string) (*model.Users, error) {
	var (
		user      *model.Users
		userId, _ = primitive.ObjectIDFromHex(id)
		filer     = bson.M{"_id": userId}
	)

	err := u.Connection.Collection("users").FindOne(ctx, filer).Decode(&user)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return user, nil
}

func (u *userRepoImpl) Save(payload *model.CreateUser) error {
	_, err := u.Connection.Collection("users").InsertOne(ctx, payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func (u *userRepoImpl) Update(id string, payload *model.UpdateUser) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectID,
	}
	updateField := bson.M{
		"$set": bson.M{
			"name":    payload.Name,
			"age":     payload.Age,
			"address": payload.Address,
		},
	}
	_, err := u.Connection.Collection("users").UpdateOne(ctx, filter, updateField)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func (u *userRepoImpl) Delete(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objectId,
	}
	_, err := u.Connection.Collection("users").DeleteOne(ctx, filter)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func NewUserRepository(connection *mongo.Database) Repository {
	return &userRepoImpl{
		Connection: connection,
	}
}
