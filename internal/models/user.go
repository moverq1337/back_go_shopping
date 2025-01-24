package models


type User struct {
    ID       int    `gorm:"primaryKey" json:"id"`
    Email    string `gorm:"unique;not null" json:"email"`
    Username string `gorm:"unique;not null" json:"username"`
    Password string `gorm:"not null" json:"password"`
    Role     string `gorm:"not null" json:"role"`
}


type Cart struct {
    ID        int `gorm:"primaryKey" json:"id"`
    UserID    int `gorm:"not null" json:"userId"`
    ProductID int `gorm:"not null" json:"productId"`
    Quantity  int `gorm:"default:1" json:"quantity"`
}
