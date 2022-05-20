package auth_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	hasher := auth.Hasher{Cost: bcrypt.MinCost}
	t.Run("hashes a password and compares it", func(t *testing.T) {
		password := "supersecret"

		hashed, err := hasher.HashPassword(password)
		assert.NoError(t, err)

		success := auth.ComparePassword(hashed, password)
		assert.True(t, success)
	})
}

func TestComparePassword(t *testing.T) {
	hasher := auth.Hasher{Cost: bcrypt.MinCost}
	t.Run("fails when passwords do not match", func(t *testing.T) {
		hashed, err := hasher.HashPassword("qwerty")
		assert.NoError(t, err)

		success := auth.ComparePassword(hashed, "azerty")
		assert.False(t, success)
	})
}
