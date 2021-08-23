package user

import (
	"time"

	"github.com/wongpinter/wongflix/internal/entity/user"
)

type UserModel struct {
	ID        user.UUID `gorm:"primary_key;column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (um *UserModel) TableName() string {
	return "users"
}

func (um *UserModel) ToEntity() *user.User {
	return &user.User{
		ID:        um.ID,
		Name:      um.Name,
		Email:     um.Email,
		Password:  um.Password,
		CreatedAt: um.CreatedAt,
		UpdatedAt: um.UpdatedAt,
	}
}
