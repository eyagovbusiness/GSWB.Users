//go:build wireinject
// +build wireinject

package di

import (
	"github.com/eyagovbusiness/GSWB.Users/src/application/UseCases/user"
	"github.com/eyagovbusiness/GSWB.Users/src/infrastructure/repositories"
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http"
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http/handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/wire"
)

func InitializeServer() *http.Server {
	wire.Build(
		ProvideDatabase,
		repositories.NewGormUserRepository,

		user.NewCreateUserUseCase,
		user.NewListUsersUseCase,

		handler.NewUserHandler,
		http.NewServer,
	)

	return &http.Server{}
}

func ProvideDatabase() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres dbname=users password=secret sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
