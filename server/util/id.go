package util

import uuid "github.com/satori/go.uuid"

func GenerateID() uuid.UUID {
	return uuid.NewV4()
}
