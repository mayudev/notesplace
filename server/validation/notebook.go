package validation

import (
	"fmt"
	"unicode/utf8"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
)

func ValidateNotebookCreate(notebook model.NotebookCreate) error {

	// Validate ProtectionLevel
	if notebook.ProtectionLevel > 2 {
		return fmt.Errorf(util.InvalidProtectionLevel)
	}

	// Validate Title
	length := utf8.RuneCountInString(notebook.Name)
	if length < 1 && length > 256 {
		return fmt.Errorf(util.InvalidTitle)
	}

	// Validate Password
	if len(notebook.Password) > 512 {
		return fmt.Errorf(util.PasswordTooLong)
	}

	if notebook.ProtectionLevel.WriteProtected() && len(notebook.Password) == 0 {
		return fmt.Errorf(util.PasswordTooShort)
	}

	return nil
}
