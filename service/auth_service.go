package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/repository"
)

// interface use case need to do for authen
type AuthService interface {
	VerifyCredential(email, password string) *entity.User
	IsDuplicateEmail(email string) bool
	IsExistPhoneNumber(dialCode, phoneNumber string) (bool, error)
	SendVerifyCode(dialCode, phoneNumber string) (entity.ID, error)
	CodeVerification(verificationID, code string) (*entity.User, error)
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
	// user := service.userRepository.FindByEmail(email)
	// if user != nil {
	// 	comparedPassword := comparePassword(user.Password, []byte(password))
	// 	if user.Email == email && comparedPassword {
	// 		return user
	// 	}
	// }

	return nil
}

func (service *authService) IsDuplicateEmail(email string) bool {
	user := service.userRepository.FindByEmail(email)
	if user != nil {
		return true
	}
	return false
}

func (service *authService) IsExistPhoneNumber(dialCode, phoneNumber string) (bool, error) {
	e, err := entity.NewPhoneNumber(dialCode, phoneNumber)
	if err != nil {
		return false, err
	}

	user := service.userRepository.FindByPhoneNumber(e)
	if user != nil {
		return true, nil
	}
	return false, nil
}

func (service *authService) CodeVerification(verificationID, code string) (*entity.User, error) {
	id, err := entity.StringToID(verificationID)
	if err != nil {
		return nil, err
	}

	phone := service.userRepository.FindVerificationByID(&id)

	if phone != nil {
		//User already exists
		user := service.userRepository.FindByPhoneNumber(&entity.PhoneNumber{
			DialCode:        phone.Phone.DialCode,
			PhoneNumber:     phone.Phone.PhoneNumber,
			FullPhoneNumber: phone.Phone.DialCode + phone.Phone.PhoneNumber,
		})

		if user != nil {
			service.userRepository.DeleleVerificationID(&id)
			return user, nil
		}

		//Create new user
		user, err = entity.NewUser(phone.Phone.DialCode, phone.Phone.PhoneNumber, "")
		if err != nil {
			return nil, err
		}

		_, err := service.userRepository.Create(user)
		if err != nil {
			return nil, err
		}
		service.userRepository.DeleleVerificationID(&id)
		return user, nil
	}

	return nil, entity.ErrNotFound
}

func (service *authService) SendVerifyCode(dialCode, phoneNumber string) (entity.ID, error) {
	e, err := entity.NewPhoneVerification(dialCode, phoneNumber)
	if err != nil {
		log.Println(err)
	}
	return service.userRepository.SaveVerificationID(e)
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
