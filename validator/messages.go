package validator

import "github.com/go-playground/validator/v10"

var validationMessages = map[string]string{
	"CreateAccountRequest.DocumentNumber.gt":            "Document number must be greater than 0.",
	"GetAccountByIDRequest.AccountID.gt":                "Account ID must be greater than 0.",
	"CreateTransactionRequest.AccountID.gt":             "Account ID must be greater than 0.",
	"CreateTransactionRequest.AccountID.required":       "Account ID is required.",
	"CreateTransactionRequest.OperationTypeID.required": "Operation type is required.",
	"CreateTransactionRequest.Amount.required":          "Amount is required.",
}

func ParseValidationErrors(err error) string {
	if err == nil {
		return ""
	}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrs {
			key := ve.StructNamespace() + "." + ve.Tag() // e.g., CreateAccountRequest.DocumentNumber.gt
			if msg, found := validationMessages[key]; found {
				return msg
			}
		}
		// fallback default message
		return validationErrs.Error()
	}
	return err.Error()
}
