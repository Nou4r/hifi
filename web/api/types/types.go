package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	OK  bool
	Err error
}

type CreateResult struct {
	Status int
	Body   []byte
	Err    error
}

type Ping struct {
	SubsonicResponse struct {
		Status string `json:"status"`
	} `json:"subsonic-response"`
}

type User struct {
	ID       string
	Username string
	Password string
}

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type RegisterResult struct {
	Success bool
	Message string
}
