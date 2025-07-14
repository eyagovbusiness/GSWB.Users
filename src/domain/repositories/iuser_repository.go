package repositories

import (
	"context"

	"github.com/eyagovbusiness/GSWB.Users/src/domain/entities"
)

type IUserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	List(ctx context.Context) ([]*entities.User, error)
}
