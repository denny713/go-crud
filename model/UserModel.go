package model

type User struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

func (u *User) TableName() string {
	return "user"
}
