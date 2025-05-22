package model

// Password represents the format of the password change request
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
