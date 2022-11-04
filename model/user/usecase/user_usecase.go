package usecase

import (
	"go-mongo/model/user/model"
	"go-mongo/model/user/repository"
	"go-mongo/utils"
)

type useUseCaseImpl struct {
	userRepository repository.Repository
}

func (u *useUseCaseImpl) FindAllUsers(name string, limit int64, page int64) ([]*model.Users, error) {
	var pages int64
	if page == 1 {
		pages = page
	} else {
		pages = pages * 10
	}
	users, err := u.userRepository.FindAll(name, limit, pages)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return users, nil
}

func (u *useUseCaseImpl) CountAllUsers() (int64, error) {
	count, err := u.userRepository.Count()
	if !utils.GlobalErrorDatabaseException(err) {
		return 0, err
	}
	return count, nil
}

func (u *useUseCaseImpl) FindUserById(id string) (*model.Users, error) {
	user, err := u.userRepository.FindById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return user, nil
}

func (u *useUseCaseImpl) CreateNewUser(payload *model.CreateUser) (*model.CreateUser, error) {
	err := u.userRepository.Save(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return nil, err
	}
	return payload, nil
}

func (u *useUseCaseImpl) UpdateUser(id string, payload *model.UpdateUser) error {
	_, err := u.userRepository.FindById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}
	errUpdate := u.userRepository.Update(id, payload)
	if !utils.GlobalErrorDatabaseException(errUpdate) {
		return errUpdate
	}
	return nil
}

func (u *useUseCaseImpl) DeleteUser(id string) error {
	_, err := u.userRepository.FindById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return err
	}

	errDel := u.userRepository.Delete(id)
	if !utils.GlobalErrorDatabaseException(errDel) {
		return errDel
	}
	return nil
}

func NewUserUseCase(userRepo repository.Repository) UseCase {
	return &useUseCaseImpl{
		userRepository: userRepo,
	}
}
