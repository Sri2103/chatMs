package model

type UserModel struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	Role         Role   `json:"role"`
	PasswordHash string `json:"password_hash,omitempty"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Role string

const Admin Role = "admin"
const Participant Role = "participant"

type AuthModel struct {
	ID     string `json:"id"`
	UserId string `json:"user_id"`
	AuthId string `json:"auth_id"`
}

type AuthDetails struct {
	UserId string `json:"user_id"`
	AuthId string `json:"auth_id"`
}
