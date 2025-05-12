package validator

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type CreateAccountRequest struct {
	DocumentNumber uint64 `validate:"gt=0"`
}

func TestParseValidationErrors_KnownKey(t *testing.T) {
	validate := validator.New()
	req := CreateAccountRequest{DocumentNumber: 0}

	err := validate.Struct(req)
	msg := ParseValidationErrors(err)

	assert.Equal(t, "Document number must be greater than 0.", msg)
}

func TestParseValidationErrors_UnknownKey(t *testing.T) {
	type Dummy struct {
		Value string `validate:"required"`
	}
	validate := validator.New()
	req := Dummy{}

	err := validate.Struct(req)
	msg := ParseValidationErrors(err)

	assert.Contains(t, msg, "Value")
}

func TestParseValidationErrors_NonValidationError(t *testing.T) {
	err := errors.New("some other error")

	msg := ParseValidationErrors(err)

	assert.Equal(t, "some other error", msg)
}

func TestParseValidationErrors_NilError(t *testing.T) {
	msg := ParseValidationErrors(nil)

	assert.Equal(t, "", msg)
}
