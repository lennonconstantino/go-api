package entity

// Auth contains the token and the id of the authenticated user
type Auth struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
