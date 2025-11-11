package types

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
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	} `json:"subsonic-response"`
}
