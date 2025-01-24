package handlers

import (
	"net/http"
	"product_api/internal/models"
	"product_api/internal/repository"
	"product_api/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
    repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
    return &ProductHandler{repo}
}

// GetAllProducts обработчик для получения всех продуктов с фильтрацией
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
    filters := make(map[string]interface{})

    // Получение параметров фильтрации из query
    if sexStr := c.Query("sex"); sexStr != "" {
        // Преобразование строки в boolean
        sex, err := strconv.ParseBool(sexStr)
        if err != nil {
            utils.JSONResponse(c, http.StatusBadRequest, "Неверный формат параметра 'sex'")
            return
        }
        filters["sex"] = sex
    }

    if isNewStr := c.Query("isNew"); isNewStr != "" {
        isNew, err := strconv.ParseBool(isNewStr)
        if err != nil {
            utils.JSONResponse(c, http.StatusBadRequest, "Неверный формат параметра 'isNew'")
            return
        }
        filters["is_new"] = isNew
    }

    products, err := h.repo.GetAll(filters)
    if err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка получения продуктов")
        return
    }
    utils.JSONResponse(c, http.StatusOK, products)
}

// GetProductByID обработчик для получения продукта по ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверный ID продукта")
        return
    }

    product, err := h.repo.GetByID(id)
    if err != nil {
        utils.JSONResponse(c, http.StatusNotFound, "Продукт не найден")
        return
    }

    utils.JSONResponse(c, http.StatusOK, product)
}

// CreateProduct обработчик для создания нового продукта
func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверные данные")
        return
    }

    if err := h.repo.Create(&product); err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка при создании продукта")
        return
    }

    utils.JSONResponse(c, http.StatusCreated, product)
}

// UpdateProduct обработчик для обновления существующего продукта
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверный ID продукта")
        return
    }

    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверные данные")
        return
    }

    // Убедитесь, что ID продукта совпадает
    if product.ID != id {
        utils.JSONResponse(c, http.StatusBadRequest, "ID продукта в URL и теле запроса не совпадают")
        return
    }

    existingProduct, err := h.repo.GetByID(id)
    if err != nil {
        utils.JSONResponse(c, http.StatusNotFound, "Продукт не найден")
        return
    }

    // Обновление полей продукта
    existingProduct.Name = product.Name
    existingProduct.Description = product.Description
    existingProduct.Imageurl = product.Imageurl
    existingProduct.Sex = product.Sex
    existingProduct.IsNew = product.IsNew
    existingProduct.Price = product.Price

    if err := h.repo.Update(existingProduct); err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка при обновлении продукта")
        return
    }

    utils.JSONResponse(c, http.StatusOK, existingProduct)
}

// DeleteProduct обработчик для удаления продукта
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        utils.JSONResponse(c, http.StatusBadRequest, "Неверный ID продукта")
        return
    }

    // Проверка существования продукта
    _, err = h.repo.GetByID(id)
    if err != nil {
        utils.JSONResponse(c, http.StatusNotFound, "Продукт не найден")
        return
    }

    if err := h.repo.Delete(id); err != nil {
        utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка при удалении продукта")
        return
    }

    utils.JSONResponse(c, http.StatusOK, "Продукт успешно удален")
}
