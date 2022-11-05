package repository

import "go-mongo/model/access_control/model"

type Repository interface {
	Count() (int64, error)
	FindAll(limit int64, page int64) ([]*model.AccessControl, error)
	FindById(id string) (*model.DetailAccessControl, error)
	Save(payload *model.CreateAccessControl) (*model.CreateAccessControl, error)
}
