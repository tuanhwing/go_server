package service

import (
	"goter.com.vn/server/entity"
	"goter.com.vn/server/repository"
)

type UserService interface {
	Profile(id entity.ID) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Profile(id entity.ID) (*entity.User, error) {
	return service.userRepository.Get(id)
}
