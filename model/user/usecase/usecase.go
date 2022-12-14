package usecase

import "go-mongo/model/user/model"

type UseCase interface {
	FindAllUsers(name string, limit int64, page int64) ([]*model.Users, error)
	CountAllUsers() (int64, error)
	FindUserById(id string) (*model.Users, error)
	CreateNewUser(payload *model.CreateUser) (*model.CreateUser, error)
	UpdateUser(id string, payload *model.UpdateUser) error
	DeleteUser(id string) error
}
