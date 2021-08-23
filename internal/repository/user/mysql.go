package user

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	entity "github.com/wongpinter/wongflix/internal/entity/user"
)

type repository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) entity.Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Store(ctx context.Context, user *entity.User) (*entity.User, error) {

	newId := uuid.New()

	result := repo.db.Create(&UserModel{
		ID:        newId,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *repository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var user UserModel

	repo.db.First(&user, id)

	return user.ToEntity(), nil
}

func (repo *repository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user UserModel

	repo.db.Where("email = ?", email).First(&user)

	return user.ToEntity(), nil
}
