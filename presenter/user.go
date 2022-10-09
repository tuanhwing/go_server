package presenter

import "goter.com.vn/server/entity"

//User data
type User struct {
	ID    entity.ID `json:"id"`
	Phone Phone     `json:"phone"`
	Name  string    `json:"name"`
}
