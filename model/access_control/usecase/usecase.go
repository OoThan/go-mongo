package usecase

import "go-mongo/model/access_control/model"

type UseCase interface {
	FindAllUsers(limit int64, page int64) ([]*model.AccessControl, error)
	CountAllUsers() (int64, error)
	FindAccessControlById(id string) (*model.DetailAccessControl, error)
	CreateNewAccessControl(payload *model.CreateAccessControl) (*model.CreateAccessControl, error)
}
