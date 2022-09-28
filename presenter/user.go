package presenter

import "goter.com.vn/server/entity"

//User data
type User struct {
	ID    entity.ID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
