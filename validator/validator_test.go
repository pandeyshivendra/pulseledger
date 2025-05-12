package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestInput struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=18"`
}

func TestValidateStruct_ValidInput(t *testing.T) {
	input := TestInput{Name: "John", Age: 25}
	err := ValidateStruct(input)
	assert.NoError(t, err)
}

func TestValidateStruct_InvalidInput(t *testing.T) {
	input := TestInput{Name: "", Age: 17}
	err := ValidateStruct(input)
	assert.Error(t, err)
}
