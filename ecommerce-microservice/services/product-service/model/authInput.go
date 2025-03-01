package model

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null"`
  CreatedAt   string    `gorm:"default:now()" json:"created_at"`
  UpdatedAt   string    `gorm:"default:now()" json:"updated_at"`

}

type CreateProduct struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}
