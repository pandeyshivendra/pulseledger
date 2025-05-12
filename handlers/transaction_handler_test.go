package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pulseledger/dto"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTransactionService struct {
	mock.Mock
}

func (m *mockTransactionService) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func setupTransactionTestApp(handler *TransactionHandler) *fiber.App {
	app := fiber.New()
	handler.RegisterRoutes(app)
	return app
}

func TestCreateTransaction_Success(t *testing.T) {
	mockService := new(mockTransactionService)
	handler := NewTransactionHandler(mockService)

	app := setupTransactionTestApp(handler)

	payload := dto.CreateTransactionRequest{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          decimal.NewFromFloat(123.45),
	}
	mockService.On("CreateTransaction", mock.Anything, payload).Return(nil)

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	var res dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&res)
	assert.True(t, res.Success)
	assert.Empty(t, res.ErrorMessage)
	mockService.AssertExpectations(t)
}

func TestCreateTransaction_InvalidPayload(t *testing.T) {
	handler := NewTransactionHandler(nil)
	app := setupTransactionTestApp(handler)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	var res dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&res)
	assert.False(t, res.Success)
	assert.Equal(t, "Invalid request body", res.ErrorMessage)
}

func TestCreateTransaction_ValidationError(t *testing.T) {
	mockService := new(mockTransactionService)
	handler := NewTransactionHandler(mockService)

	app := setupTransactionTestApp(handler)

	payload := dto.CreateTransactionRequest{
		AccountID:       0,
		OperationTypeID: 0,
		Amount:          decimal.NewFromFloat(0),
	}
	mockService.On("CreateTransaction", mock.Anything, payload).Return(errors.New("validation error"))

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	var res dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&res)
	assert.False(t, res.Success)
	assert.Equal(t, "validation error", res.ErrorMessage)
	mockService.AssertExpectations(t)
}
