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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAccountService struct {
	mock.Mock
}

func (m *mockAccountService) CreateAccount(ctx context.Context, req dto.CreateAccountRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *mockAccountService) GetAccountByID(ctx context.Context, req dto.GetAccountByIDRequest) (*dto.GetAccountByIDResponse, error) {
	args := m.Called(ctx, req)
	if res, ok := args.Get(0).(*dto.GetAccountByIDResponse); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func setupTestApp(handler *AccountHandler) *fiber.App {
	app := fiber.New()
	handler.RegisterRoutes(app)
	return app
}

func TestCreateAccount_Success(t *testing.T) {
	mockService := new(mockAccountService)
	handler := NewAccountHandler(mockService)

	app := setupTestApp(handler)

	payload := dto.CreateAccountRequest{DocumentNumber: 12345678900}
	mockService.On("CreateAccount", mock.Anything, payload).Return(nil)

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var apiResp dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	assert.True(t, apiResp.Success)

	mockService.AssertExpectations(t)
}

func TestCreateAccount_InvalidPayload(t *testing.T) {
	handler := NewAccountHandler(nil)
	app := setupTestApp(handler)

	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var apiResp dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	assert.False(t, apiResp.Success)
	assert.Equal(t, "Invalid request body", apiResp.ErrorMessage)
}

func TestCreateAccount_ValidationError(t *testing.T) {
	mockService := new(mockAccountService)
	handler := NewAccountHandler(mockService)

	app := setupTestApp(handler)

	payload := dto.CreateAccountRequest{DocumentNumber: 0}
	mockService.On("CreateAccount", mock.Anything, payload).Return(errors.New("validation error"))

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var apiResp dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	assert.False(t, apiResp.Success)
	assert.Equal(t, "validation error", apiResp.ErrorMessage)

	mockService.AssertExpectations(t)
}

func TestGetAccountByID_Success(t *testing.T) {
	mockService := new(mockAccountService)
	handler := NewAccountHandler(mockService)

	app := setupTestApp(handler)

	accountID := uint64(1)
	expected := &dto.GetAccountByIDResponse{
		AccountId:      accountID,
		DocumentNumber: 12345678900,
	}

	mockService.On("GetAccountByID", mock.Anything, dto.GetAccountByIDRequest{AccountID: accountID}).Return(expected, nil)

	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var apiResp dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	assert.True(t, apiResp.Success)

	data, ok := apiResp.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.EqualValues(t, expected.AccountId, uint64(data["account_id"].(float64)))
	assert.EqualValues(t, expected.DocumentNumber, uint64(data["document_number"].(float64)))

	mockService.AssertExpectations(t)
}

func TestGetAccountByID_InvalidID(t *testing.T) {
	handler := NewAccountHandler(nil)
	app := setupTestApp(handler)

	req := httptest.NewRequest(http.MethodGet, "/accounts/abc", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var apiResp dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	assert.False(t, apiResp.Success)
	assert.Equal(t, "Invalid account ID", apiResp.ErrorMessage)
}

func TestGetAccountByID_NotFound(t *testing.T) {
	mockService := new(mockAccountService)
	handler := NewAccountHandler(mockService)

	app := setupTestApp(handler)

	mockService.On("GetAccountByID", mock.Anything, dto.GetAccountByIDRequest{AccountID: 2}).Return(nil, errors.New("not found"))

	req := httptest.NewRequest(http.MethodGet, "/accounts/2", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	var apiResp dto.APIResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	assert.False(t, apiResp.Success)
	assert.Equal(t, "not found", apiResp.ErrorMessage)

	mockService.AssertExpectations(t)
}
