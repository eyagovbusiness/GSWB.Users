package repositories

import (
	"context"
	"fmt"

	"github.com/eyagovbusiness/GSWB.Users/src/domain/entities"
	"github.com/eyagovbusiness/GSWB.Users/src/domain/repositories"
	"github.com/eyagovbusiness/GSWB.Users/src/infrastructure/persistence/models"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) Create(ctx context.Context, user *entities.User) error {
	model := models.UserModel{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *GormUserRepository) List(ctx context.Context) ([]*entities.User, error) {
	var models []models.UserModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	var users []*entities.User
	for _, m := range models {
		users = append(users, &entities.User{
			ID:        m.ID,
			Name:      m.Name,
			Email:     m.Email,
			CreatedAt: m.CreatedAt,
		})
	}

	return users, nil
}
