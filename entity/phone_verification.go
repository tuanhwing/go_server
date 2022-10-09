package entity

import "time"

type PhoneVerification struct {
	ID        ID
	Phone     PhoneNumber
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewPhoneNumber create a new user
func NewPhoneVerification(dialCode, phoneNumber string) (*PhoneVerification, error) {
	p := &PhoneVerification{
		ID: NewID(),
		Phone: PhoneNumber{
			DialCode:        dialCode,
			PhoneNumber:     phoneNumber,
			FullPhoneNumber: dialCode + phoneNumber,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return p, nil
}
