package validation_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mayudev/notesplace/server/model"
	"github.com/mayudev/notesplace/server/util"
	"github.com/mayudev/notesplace/server/validation"
)

func TestValidateNotebookCreate(t *testing.T) {
	cases := []struct {
		Name          string
		Notebook      model.NotebookCreate
		ExpectedError error
	}{
		{
			Notebook: model.NotebookCreate{
				Name:            "Correct",
				ProtectionLevel: 0,
			},
			ExpectedError: nil,
		},
		{
			Notebook: model.NotebookCreate{
				Name:            "Correct",
				ProtectionLevel: 3,
			},
			ExpectedError: fmt.Errorf(util.InvalidProtectionLevel),
		},
		{
			Notebook: model.NotebookCreate{
				Name:            strings.Repeat("Long", 256),
				ProtectionLevel: 0,
			},
			ExpectedError: fmt.Errorf(util.InvalidTitle),
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			got := validation.ValidateNotebookCreate(tt.Notebook)

			if got != nil && got.Error() != tt.ExpectedError.Error() {
				t.Errorf("got incorrect error, got %v want %v", got, tt.ExpectedError)
			}

			if got == nil && tt.ExpectedError != nil {
				t.Errorf("got incorrect error, got %v want %v", got, tt.ExpectedError)
			}
		})
	}
	return
}
