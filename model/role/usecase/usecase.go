package usecase

import "go-mongo/model/role/model"

type UseCase interface {
	CountAllRoles() (int64, error)
	FindAllRoles(limit int64, page int64) ([]*model.Role, error)
	FindRoleById(id string) (*model.Role, error)
	CreateNewRole(payload *model.CreateRole) error
	UpdateRole(id string, payload *model.UpdateRole) error
	Delete(id string) error
}
