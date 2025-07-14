package user

import (
	"context"

	"github.com/eyagovbusiness/GSWB.Users/src/application/dtos"
	"github.com/eyagovbusiness/GSWB.Users/src/application/errors"
	"github.com/eyagovbusiness/GSWB.Users/src/domain/entities"
	"github.com/eyagovbusiness/GSWB.Users/src/domain/repositories"
)

type CreateUserUseCase struct {
	userRepository repositories.IUserRepository
}

func NewCreateUserUseCase(userRepository repositories.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input dtos.CreateUserInput) (*dtos.UserOutput, error) {
	user, err := entities.NewUser(input.Name, input.Email)
	if err != nil {
		return nil, errors.NewValidationError(err.Error())
	}

	if err := uc.userRepository.Create(ctx, user); err != nil {
		return nil, errors.NewUnexpectedError("failed to persist user", err)
	}

	return &dtos.UserOutput{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
