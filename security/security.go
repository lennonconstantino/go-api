package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and hash it
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerifyPassword compare a password and a hash and return whether they are equal
func VerifyPassword(senhaString, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaString), []byte(senhaComHash))
}
