package dto

type LoginRequestDto struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
