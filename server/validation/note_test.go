package validation_test

import (
	"strings"
	"testing"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/validation"
)

func TestValidateNoteUpdate(t *testing.T) {
	cases := []struct {
		Name     string
		Note     model.Note
		Expected bool
	}{
		{
			Note: model.Note{
				Title:   "Correct",
				Content: strings.Repeat("a", 20001),
			},
			Expected: false,
		},
		{
			Note: model.Note{
				Title:   strings.Repeat("a", 101),
				Content: "Correct",
			},
			Expected: false,
		},
		{
			Note: model.Note{
				Title:   "Correct",
				Content: "An exemplary note",
			},
			Expected: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			got := validation.ValidateNoteUpdate(&tt.Note)

			if got != tt.Expected {
				t.Errorf("got incorrect validation, got %v want %v", got, tt.Expected)
			}
		})
	}
	return
}
