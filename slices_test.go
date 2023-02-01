package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterOK(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		// Given

		numbers := []int{1, 2, 3, 4, 5, 6}
		expected := []int{2, 4, 6}

		// then

		result := Filter(numbers, func(v int) bool {
			return v%2 == 0
		})

		// when

		assert.Equal(t, expected, result)
	})
}
