package types

type SignupRequest struct {
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
