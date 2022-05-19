package util_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/util"
	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	t.Run("generates unique ids", func(t *testing.T) {
		id1 := util.GenerateID()
		id2 := util.GenerateID()

		assert.NotEqual(t, id1, id2)
	})
}
