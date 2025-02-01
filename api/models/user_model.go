package models

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"biography"`
}

func (u *User) SetId(id uint64) {
	u.ID = id
}

func (u *User) GetId() uint64 {
	return u.ID
}
