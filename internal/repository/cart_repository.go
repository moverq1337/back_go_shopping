package repository

import (
	"product_api/internal/models"

	"gorm.io/gorm"
)

type CartRepository interface {
    AddToCart(cart *models.Cart) error
    GetCartByUserID(userID int) ([]models.Cart, error)
    UpdateCartItem(cartID int, quantity int) error
    RemoveFromCart(cartID int) error
}

type cartRepository struct {
    db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
    return &cartRepository{db}
}

func (r *cartRepository) AddToCart(cart *models.Cart) error {
    return r.db.Create(cart).Error
}

func (r *cartRepository) GetCartByUserID(userID int) ([]models.Cart, error) {
    var cartItems []models.Cart
    // Preload для загрузки связанных данных (Product)
    err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error
    return cartItems, err
}


func (r *cartRepository) UpdateCartItem(cartID int, quantity int) error {
    return r.db.Model(&models.Cart{}).Where("id = ?", cartID).Update("quantity", quantity).Error
}

func (r *cartRepository) RemoveFromCart(cartID int) error {
    return r.db.Delete(&models.Cart{}, cartID).Error
}
