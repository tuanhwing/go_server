package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"goter.com.vn/server/dto"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/repository"
)

// interface use case need to do for authen
type AuthService interface {
	VerifyCredential(email string, password string) *entity.User
	CreateUser(user dto.RegisterDTO) (entity.ID, error)
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) *entity.User {
	user := service.userRepository.FindByEmail(email)
	if user != nil {
		comparedPassword := comparePassword(user.Password, []byte(password))
		if user.Email == email && comparedPassword {
			return user
		}
	}

	return nil
}

func (service *authService) CreateUser(user dto.RegisterDTO) (entity.ID, error) {
	e, err := entity.NewUser(user.Email, user.Password, user.Name)
	if err != nil {
		log.Println(err)
	}
	return service.userRepository.Create(e)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	user := service.userRepository.FindByEmail(email)
	if user != nil {
		return true
	}
	return false
}

func comparePassword(hashedPwd string, planPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, planPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
