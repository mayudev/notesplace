package validation

import (
	"unicode/utf8"

	"github.com/mayudev/notesplace/server/model"
)

// ValidateNoteUpdate verifies if fields that will be updated are correct
func ValidateNoteUpdate(note *model.Note) bool {

	// Verify Content and Title lengths
	if utf8.RuneCountInString(note.Content) > 20000 {
		return false
	}

	if utf8.RuneCountInString(note.Title) > 100 {
		return false
	}

	return true
}
