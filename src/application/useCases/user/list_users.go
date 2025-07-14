package user

import (
	"context"

	"github.com/eyagovbusiness/GSWB.Users/src/application/dtos"
	"github.com/eyagovbusiness/GSWB.Users/src/domain/repositories"
)

type ListUsersUseCase struct {
	userRepository repositories.IUserRepository
}

func NewListUsersUseCase(userRepository repositories.IUserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		userRepository: userRepository,
	}
}

func (uc *ListUsersUseCase) Execute(ctx context.Context) ([]*dtos.UserOutput, error) {
	users, err := uc.userRepository.List(ctx)
	if err != nil {
		return nil, err // optionally wrap with app error
	}

	var result []*dtos.UserOutput
	for _, u := range users {
		result = append(result, &dtos.UserOutput{
			ID:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return result, nil
}
