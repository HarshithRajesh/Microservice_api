package model

type Users struct{
  ID           uint `gorm:"primaryKey;autoIncrement" json:"id"`
  FullName     string `gorm:"not null" json:"full_name"`
  Email        string `gorm:"unique;not null" json:"email"`
  Password string `gorm:"not null" json:"password"`
  CreatedAt    string `gorm:"default:now()" json:"created_at"`
  UpdatedAt    string `gorm:"default:now()" json:"updated_at"`

}

type SignUp struct{
  FullName     string `json:"full_name" binding:"required"`
  Email        string `json:"email" binding:"required,email"`
  Password     string `json:"password" binding:"required"`

}
