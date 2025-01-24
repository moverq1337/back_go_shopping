package models

type User struct {
    ID       int    `gorm:"primaryKey" json:"id"`
    Email    string `gorm:"unique;not null" json:"email"`
    Username string `gorm:"unique;not null" json:"username"`
    Password string `gorm:"not null" json:"password"`
    Role     string `gorm:"not null" json:"role"`
    Cart     []Cart `gorm:"foreignKey:UserID" json:"cart"` // Связь с корзиной
}



type Cart struct {
    ID        int     `gorm:"primaryKey" json:"id"`
    UserID    int     `gorm:"not null" json:"userId"`
    ProductID int     `gorm:"not null" json:"productId"`
    Quantity  int     `gorm:"default:1" json:"quantity"`
    Product   Product `gorm:"foreignKey:ProductID" json:"product"` // Связь с продуктом
    User      User    `gorm:"foreignKey:UserID" json:"user"`       // Связь с пользователем
}
