package model

import (
	"github.com/mayudev/notesplace/server/auth"
	"github.com/mayudev/notesplace/server/util"
)

type NotebookCreate struct {
	// Name of the notebook. Optional.
	Name string `json:"name"`
	// Protection level.
	//
	// Enum - possible values: "none" | "readonly" | "protected" (0 | 1 | 2) (db: enum)
	//
	// None | 0 - everyone has access
	//
	// ReadOnly | 1 - read only. password needed to write
	//
	// Protected | 2 - protected. password needed to read
	//
	// Required.
	ProtectionLevel auth.ProtectionLevel `json:"protection_level"`
	// Password (unencoded)
	Password string `json:"password"`
}

type NotebookCreateResponse struct {
	// Assigned ID
	ID string `json:"id"`
	// Response information
	util.Response
}
