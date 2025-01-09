package model

type UserModel struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	passwordHash string `json:"password_hash"`
}

type Credentials struct {
	UserName string `json:"user_name"`
	password string `json:"password"`
}
