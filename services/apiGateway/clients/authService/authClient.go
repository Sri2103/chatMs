package authservice

type AuthDetails struct {
	AuthID string `json:"auth_id"`
	UserID string `json:"user_id"`
}

type Auth struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	AuthID string `json:"auth_id"`
}

type AutheService interface {
	FetchAuth(*AuthDetails) (*Auth, error)
	CreateAuth(*AuthDetails) (*Auth, error)
	DeleteAuth(*AuthDetails) error
}
