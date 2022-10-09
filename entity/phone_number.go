package entity

import "regexp"

type PhoneNumber struct {
	DialCode        string
	PhoneNumber     string
	FullPhoneNumber string
}

//NewPhoneNumber create a new user
func NewPhoneNumber(dialCode, phoneNumber string) (*PhoneNumber, error) {
	p := &PhoneNumber{
		DialCode:        dialCode,
		PhoneNumber:     phoneNumber,
		FullPhoneNumber: dialCode + phoneNumber,
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

//Validate  data
func (p *PhoneNumber) Validate() error {
	re := regexp.MustCompile(`\+[1-9][0-9]{6,12}`)
	phoneNumber := p.DialCode + p.PhoneNumber
	if re.MatchString(phoneNumber) {
		return nil
	}
	return ErrInvalidEntity
}
