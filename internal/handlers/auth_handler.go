package handlers

import (
	"log"
	"net/http"
	"product_api/internal/models"
	"product_api/internal/repository"
	"product_api/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo}
}

// Register обработчик регистрации
func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Неверные данные")
		return
	}

	// Хэширование пароля с использованием bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка создания пользователя")
		return
	}
	user.Password = string(hashedPassword)

	// Сохранение пользователя
	if err := h.userRepo.CreateUser(&user); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка создания пользователя")
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Пользователь успешно зарегистрирован")
}

// Login обработчик логина
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Неверные данные")
		return
	}

	// Поиск пользователя по email
	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "Неверный email или пароль")
		return
	}

	// Сравнение введенного пароля с хэшированным паролем
	log.Printf("Запрос на логин: email=%s, password=%s", req.Email, req.Password)
	log.Printf("Найден пользователь: %+v", user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Printf("Ошибка сравнения пароля: %v", err)
		utils.JSONResponse(c, http.StatusUnauthorized, "Неверный email или пароль")
		return
	}

	// Генерация JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Ошибка авторизации")
		return
	}

	utils.JSONResponse(c, http.StatusOK, gin.H{
		"token": token,
	})
}
