package yeast_test

import (
	"testing"

	"github.com/mayudev/notesplace/server/util/yeast"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	t.Run("returns the same result as the nodejs counterpart", func(t *testing.T) {
		want := "O371ltp"
		got := yeast.Encode(1652606565875)

		assert.Equal(t, got, want)
	})
}

func TestDecode(t *testing.T) {
	t.Run("can be reversed", func(t *testing.T) {
		var seed int64 = 1652606565875

		encoded := yeast.Encode(seed)
		decoded := yeast.Decode(encoded)

		assert.Equal(t, seed, decoded)
	})

	t.Run("does give different results", func(t *testing.T) {
		var seed int64 = 55555555

		encoded := yeast.Encode(seed)
		decoded := yeast.Decode(encoded)

		assert.Equal(t, seed, decoded)
	})
}
