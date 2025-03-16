package handler

import (
	"encoding/json"
	"errors"
	"example-sqlc-postgre/internal/db/sqlcdb"
	"example-sqlc-postgre/orderservice"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"net/http"
)

type OrderHandler struct {
	logger       *zerolog.Logger
	orderService orderservice.OrderService
}

func NewOrderHandler(logger *zerolog.Logger, orderService orderservice.OrderService) *OrderHandler {
	return &OrderHandler{
		logger:       logger,
		orderService: orderService,
	}
}

type OrderCreateRequest struct {
	AccountId string          `json:"account_id"`
	Symbol    string          `json:"symbol"`
	Quantity  decimal.Decimal `json:"quantity"`
	Type      string          `json:"type"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var req OrderCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error().Err(err).Msg("Failed to decode request body")
		handleError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate request (using validator package or custom validation)
	if err := validateOrderRequest(&req); err != nil {
		h.logger.Warn().Err(err).
			Str("accountId", req.AccountId).
			Str("symbol", req.Symbol).
			Str("type", req.Type).
			Msg("Request validation failed")
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	param := &orderservice.OrderCreateParam{
		AccountId: req.AccountId,
		Symbol:    req.Symbol,
		Quantity:  req.Quantity,
		Type:      sqlcdb.OrderType(req.Type),
	}

	if err := h.orderService.CreateOrder(r.Context(), param); err != nil {
		h.logger.Error().Err(err).
			Str("accountId", req.AccountId).
			Str("symbol", req.Symbol).
			Str("type", req.Type).
			Msg("Failed to create order")
		handleError(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
		Message: "Order created successfully",
	})

}

// validateOrderRequest performs basic validation on the request
func validateOrderRequest(req *OrderCreateRequest) error {
	if req.AccountId == "" {
		return errors.New("account_id is required")
	}
	if req.Symbol == "" {
		return errors.New("symbol is required")
	}
	if req.Quantity.IsZero() || req.Quantity.IsNegative() {
		return errors.New("quantity must be a positive number")
	}
	if req.Type == "" {
		return errors.New("type is required")
	}
	if req.Type != "market" && req.Type != "limit" {
		return errors.New("type must be either market or limit")
	}
	return nil
}

// handleError sends an error response
func handleError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
