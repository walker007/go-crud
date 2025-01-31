package models

type User struct {
	ID        uint64 `json:"id"`
	FirstName string
	LastName  string
	Biography string
}

func (u *User) SetId(id uint64) {
	u.ID = id
}

func (u *User) GetId() uint64 {
	return u.ID
}
