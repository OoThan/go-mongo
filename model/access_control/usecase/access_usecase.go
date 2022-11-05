package usecase

import (
	"go-mongo/model/access_control/model"
	"go-mongo/model/access_control/repository"
	"go-mongo/utils"
)

type accessUseCaseImpl struct {
	accessRepository repository.Repository
}

func (u *accessUseCaseImpl) FindAllUsers(limit int64, page int64) ([]*model.AccessControl, error) {
	var pages int64
	if page == 1 {
		pages = page
	} else {
		pages = page * 10
	}
	acls, err := u.accessRepository.FindAll(limit, pages)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return acls, nil
}

func (u *accessUseCaseImpl) CountAllUsers() (int64, error) {
	count, err := u.accessRepository.Count()
	if !utils.GlobalErrorDatabaseException(err) {
		return 0, err
	}
	return count, nil
}

func (u *accessUseCaseImpl) FindAccessControlById(id string) (*model.DetailAccessControl, error) {
	acl, err := u.accessRepository.FindById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return acl, nil
}

func (u *accessUseCaseImpl) CreateNewAccessControl(payload *model.CreateAccessControl) (*model.CreateAccessControl, error) {
	_, err := u.accessRepository.Save(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return payload, nil
}

func NewAccessControlUseCase(accessRepo repository.Repository) UseCase {
	return &accessUseCaseImpl{
		accessRepository: accessRepo,
	}
}
