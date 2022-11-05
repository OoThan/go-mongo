package repository

import "go-mongo/model/role/model"

type Repository interface {
	Count() (int64, error)
	FindAll(limit int64, page int64) ([]*model.Role, error)
	FindById(id string) (*model.Role, error)
	Save(payload *model.CreateRole) error
	Update(id string, payload *model.UpdateRole) error
	Delete(id string) error
}
