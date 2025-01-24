package repository

import (
	"product_api/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
    GetAll(filters map[string]interface{}) ([]models.Product, error)
    GetByID(id int) (*models.Product, error)
    Create(product *models.Product) error
    Update(product *models.Product) error
    Delete(id int) error
}

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{db}
}

func (r *productRepository) GetAll(filters map[string]interface{}) ([]models.Product, error) {
    var products []models.Product
    query := r.db

    if sex, ok := filters["sex"]; ok {
        query = query.Where("sex = ?", sex)
    }

    if isNew, ok := filters["is_new"]; ok {
        query = query.Where("is_new = ?", isNew)
    }

    result := query.Find(&products)
    return products, result.Error
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
    var product models.Product
    result := r.db.First(&product, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &product, nil
}

func (r *productRepository) Create(product *models.Product) error {
    return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
    return r.db.Save(product).Error
}

func (r *productRepository) Delete(id int) error {
    return r.db.Delete(&models.Product{}, id).Error
}
