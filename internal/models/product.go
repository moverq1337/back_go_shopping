package models

type Product struct {
    ID          int     `gorm:"primaryKey" json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Imageurl    string  `json:"Imageurl"`
    Sex         bool    `json:"sex"`
    IsNew       bool    `json:"isNew"`
    Price       float64 `json:"price"`
}

// TableName возвращает имя таблицы для модели Product
func (Product) TableName() string {
    return "product"
}
