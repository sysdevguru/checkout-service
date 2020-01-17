package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAmount(t *testing.T) {
	items := []string{"PEN", "MUG", "TSHIRT", "TSHIRT", "PEN", "TSHIRT", "MUG"}
	assert.Equal(t, "60.00", CalculateAmount(items))

	items1 := []string{"PEN", "PEN", "TSHIRT", "MUG", "PEN", "MUG", "TSHIRT", "TSHIRT", "PEN", "TSHIRT", "MUG"}
	assert.Equal(t, "92.50", CalculateAmount(items1))

	items2 := []string{"PEN", "TSHIRT", "MUG"}
	assert.Equal(t, "32.50", CalculateAmount(items2))
}