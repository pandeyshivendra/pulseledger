package dto

type CreateAccountRequest struct {
	DocumentNumber uint64 `json:"document_number" validate:"gt=0"`
}

type GetAccountByIDRequest struct {
	AccountID uint64 `json:"account_id" validate:"gt=0"`
}

type GetAccountByIDResponse struct {
	AccountId      uint64 `json:"account_id"`
	DocumentNumber uint64 `json:"document_number"`
}
