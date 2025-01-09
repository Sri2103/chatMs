package model

type UserModel struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	Role         Role   `json:"role"`
	PasswordHash string `json:"password_hash,omitempty"`
}

type Credentials struct {
	UserName string `json:"user_name"`
	password string `json:"password"`
}

type Role string

const Admin Role = "admin"
const Participant Role = "participant"
