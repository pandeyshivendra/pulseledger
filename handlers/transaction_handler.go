package handlers

import (
	"pulseledger/dto"
	"pulseledger/services"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/transactions", h.CreateTransaction)
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var req dto.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success:      false,
			ErrorMessage: "Invalid request body",
		})
	}

	if err := h.service.CreateTransaction(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.APIResponse{
		Success: true,
	})
}
