package handlers

import (
	"net/http"
	"product_api/internal/models"
	"product_api/internal/repository"
	"product_api/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
    cartRepo repository.CartRepository
}

func NewCartHandler(cartRepo repository.CartRepository) *CartHandler {
    return &CartHandler{cartRepo}
}

// AddToCart добавляет товар в корзину
func (h *CartHandler) AddToCart(c *gin.Context) {
    userID, _ := c.Get("userId")

    var req struct {
        ProductID int `json:"productId"`
        Quantity  int `json:"quantity"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверные данные")
        return
    }

    if req.Quantity < 1 {
        utils.JSONResponse(c, http.StatusBadRequest, "Количество должно быть больше нуля")
        return
    }

    cartItem := &models.Cart{
        UserID:    userID.(int),
        ProductID: req.ProductID,
        Quantity:  req.Quantity,
    }

    if err := h.cartRepo.AddToCart(cartItem); err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка добавления товара в корзину")
        return
    }

    utils.JSONResponse(c, http.StatusCreated, "Товар добавлен в корзину")
}

// GetCart получает товары в корзине пользователя
func (h *CartHandler) GetCart(c *gin.Context) {
    userID, _ := c.Get("userId")

    cartItems, err := h.cartRepo.GetCartByUserID(userID.(int))
    if err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка получения корзины")
        return
    }

    utils.JSONResponse(c, http.StatusOK, cartItems)
}

// UpdateCartItem обновляет количество товара в корзине
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
    cartIDStr := c.Param("id")
    cartID, err := strconv.Atoi(cartIDStr) // Преобразование строки в int
    if err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверный формат ID")
        return
    }

    var req struct {
        Quantity int `json:"quantity"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверные данные")
        return
    }

    if req.Quantity < 1 {
        utils.JSONResponse(c, http.StatusBadRequest, "Количество должно быть больше нуля")
        return
    }

    if err := h.cartRepo.UpdateCartItem(cartID, req.Quantity); err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка обновления товара в корзине")
        return
    }

    utils.JSONResponse(c, http.StatusOK, "Количество товара обновлено")
}

// RemoveFromCart удаляет товар из корзины
func (h *CartHandler) RemoveFromCart(c *gin.Context) {
    cartIDStr := c.Param("id")
    cartID, err := strconv.Atoi(cartIDStr) // Преобразование строки в int
    if err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверный формат ID")
        return
    }

    if err := h.cartRepo.RemoveFromCart(cartID); err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка удаления товара из корзины")
        return
    }

    utils.JSONResponse(c, http.StatusOK, "Товар удален из корзины")
}
