package slices

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		// Given

		numbers := []int{1, 2, 3, 4, 5, 6}
		expected := []int{2, 4, 6}

		// Then

		result := Filter(numbers, func(v int) bool {
			return v%2 == 0
		})

		// When

		assert.Equal(t, expected, result)
	})
}

func TestReduce(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		// Given

		numbers := []int{1, 2, 3}

		// Then
		sum := Reduce(numbers, func(acc, current int) int {
			return acc + current
		}, 0)

		// When
		assert.Equal(t, 6, sum)
	})
}

func TestMap(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		// Given

		numbers := []float64{4, 9}
		expected := []float64{2, 3}

		// Then

		newNumbers := Map(numbers, math.Sqrt)

		// When
		assert.Equal(t, expected, newNumbers)
	})
}
