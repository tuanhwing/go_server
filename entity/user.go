package entity

import (
	"time"
)

type User struct {
	ID        ID
	Name      string
	Phone     PhoneNumber
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewUser create a new user
func NewUser(dialCode, phone, name string) (*User, error) {
	p, err := NewPhoneNumber(dialCode, phone)

	if err != nil {
		return nil, err
	}

	u := &User{
		ID:        NewID(),
		Name:      name,
		Phone:     *p,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//Validate  data
func (u *User) Validate() error {
	return nil
}

// //ValidatePassword validate user password
// func (u *User) ValidatePassword(p string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func generatePassword(raw string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hash), nil
// }

// func (user User) String() string {
// 	return fmt.Sprintf("%s | %s | %s | %s | %s", user.ID.String(), user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
// }
