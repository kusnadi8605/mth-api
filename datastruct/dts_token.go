package datastruct

//User ..
type User struct {
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
}

//Token ..
type Token struct {
	Token string `json:"token,omitempty"`
}

//TokenReq ..
type TokenReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

//TokenResponse data
type TokenResponse struct {
	ResponseCode string  `json:"responseCode"`
	ResponseDesc string  `json:"responseDesc"`
	Payload      []Token `json:"payload,omitempty"`
}
