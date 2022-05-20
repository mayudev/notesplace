package auth

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	Cost int
}

func NewHasher() *Hasher {
	return &Hasher{
		Cost: bcrypt.DefaultCost,
	}
}

// HashPassword hashes a password
func (h *Hasher) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), h.Cost)

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
