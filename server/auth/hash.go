package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// ComparePassword compares a hash and a password, returning true on match
func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
