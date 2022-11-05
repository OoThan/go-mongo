package usecase

import (
	"go-mongo/model/role/model"
	"go-mongo/model/role/repository"
	"go-mongo/utils"
)

type roleUseCaseImpl struct {
	roleRepository repository.Repository
}

func (u *roleUseCaseImpl) CountAllRoles() (int64, error) {
	count, err := u.roleRepository.Count()
	if !utils.GlobalErrorDatabaseException(err) {
		return 0, err
	}
	return count, nil
}

func (u *roleUseCaseImpl) FindAllRoles(limit int64, page int64) ([]*model.Role, error) {
	var pages int64
	if page == 1 {
		pages = page
	} else {
		pages = page * 10
	}
	roles, err := u.roleRepository.FindAll(limit, pages)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return roles, nil
}

func (u *roleUseCaseImpl) FindRoleById(id string) (*model.Role, error) {
	role, err := u.roleRepository.FindById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return role, nil
}

func (u *roleUseCaseImpl) CreateNewRole(payload *model.CreateRole) error {
	err := u.roleRepository.Save(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func (u *roleUseCaseImpl) UpdateRole(id string, payload *model.UpdateRole) error {
	_, err := u.roleRepository.FindById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	err = u.roleRepository.Update(id, payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func (u *roleUseCaseImpl) Delete(id string) error {
	_, err := u.roleRepository.FindById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	err = u.roleRepository.Delete(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	return nil
}

func NewRoleUseCase(repo repository.Repository) UseCase {
	return &roleUseCaseImpl{
		roleRepository: repo,
	}
}
