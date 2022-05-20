package auth_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/auth"
	"github.com/stretchr/testify/assert"
)

func TestIssueJWT(t *testing.T) {
	t.Run("issues and validates a valid JWT", func(t *testing.T) {
		issuer := auth.NewIssuer("unsafe_key")

		token, err := issuer.Issue("notebook_id")
		assert.NoError(t, err)

		validated, err := issuer.Validate(token)
		assert.NoError(t, err)
		assert.Equal(t, validated.NotebookID, "notebook_id")
	})

	t.Run("does not validate an invalid JWT", func(t *testing.T) {
		issuer := auth.NewIssuer("unsafe_key")

		token, err := issuer.Issue("notebook_id")
		assert.NoError(t, err)

		token += "temper_temper"

		_, err = issuer.Validate(token)
		assert.Error(t, err)
	})
}

func TestValidate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		issuer := auth.NewIssuer("unsafe_key")

		token, err := issuer.Issue("notebook_id")
		assert.NoError(t, err)

		valid := issuer.ValidateNotebook(token, "notebook_id")
		assert.True(t, valid)
	})

	t.Run("does not validate if token grants access to a different notebook", func(t *testing.T) {
		issuer := auth.NewIssuer("unsafe_key")

		token, err := issuer.Issue("different_notebook")
		assert.NoError(t, err)

		valid := issuer.ValidateNotebook(token, "notebook_id")
		assert.False(t, valid)
	})
}
