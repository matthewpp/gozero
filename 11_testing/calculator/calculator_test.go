package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := Add(2, 3)

		assert.Equal(t, 5, result)
	})

	t.Run("fail", func(t *testing.T) {
		result := Add(2, 2)

		assert.NotEqual(t, 5, result)
	})
}

func TestMinus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := Minus(5, 3)

		assert.Equal(t, 2, result)
	})

	t.Run("fail", func(t *testing.T) {
		result := Minus(5, 2)

		assert.NotEqual(t, 2, result)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := Multiply(2, 3)

		assert.Equal(t, 6, result)
	})

	t.Run("fail", func(t *testing.T) {
		result := Multiply(2, 2)

		assert.NotEqual(t, 6, result)
	})
}

func TestDivide(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := Divide(6, 3)

		assert.Nil(t, err)
		assert.Equal(t, 2.0, result)
	})

	t.Run("fail divide by zero", func(t *testing.T) {
		result, err := Divide(6, 0)

		assert.NotNil(t, err)
		assert.Equal(t, ErrDivideByZero, err)
		assert.Equal(t, 0.0, result)
	})
}
