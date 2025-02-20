package model

type Users struct{
	ID           string `gorm:"type:uuid;primaryKey"`
	FullName     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt    string `gorm:"default:now()"`
	UpdatedAt    string `gorm:"default:now()"`

}
