package handlers

import (
	"strconv"

	"pulseledger/dto"
	"pulseledger/services"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	Service services.AccountService
}

func NewAccountHandler(service services.AccountService) *AccountHandler {
	return &AccountHandler{Service: service}
}

func (h *AccountHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/accounts", h.CreateAccount)
	router.Get("/accounts/:id", h.GetAccountByID)
}

func (h *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var req dto.CreateAccountRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success:      false,
			ErrorMessage: "Invalid request body",
		})
	}

	if err := h.Service.CreateAccount(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.APIResponse{
		Success: true,
	})
}

func (h *AccountHandler) GetAccountByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success:      false,
			ErrorMessage: "Invalid account ID",
		})
	}

	req := dto.GetAccountByIDRequest{AccountID: id}
	res, err := h.Service.GetAccountByID(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.APIResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    res,
	})
}
