//go:build wireinject
// +build wireinject

package di

import (
	"fmt"
	"os"

	"github.com/eyagovbusiness/GSWB.Users/src/application/useCases/user"
	"github.com/eyagovbusiness/GSWB.Users/src/infrastructure/persistence/models"
	"github.com/eyagovbusiness/GSWB.Users/src/infrastructure/persistence/repositories"
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http"
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http/handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/wire"
)

func InitializeServer() (*http.Server, error) {
	wire.Build(
		ProvideDatabase,
		repositories.NewGormUserRepository,
		user.NewCreateUserUseCase,
		user.NewListUsersUseCase,
		handler.NewUserHandler,
		http.NewServer,
	)
	return nil, nil
}

func ProvideDatabase() (*gorm.DB, error) {
	// Read from environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Debug log to confirm what we read
	fmt.Println("[ENV] DB_HOST =", host)
	fmt.Println("[ENV] DB_PORT =", port)
	fmt.Println("[ENV] DB_USER =", user)
	fmt.Println("[ENV] DB_PASSWORD =", password)
	fmt.Println("[ENV] DB_NAME =", dbname)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		host, port, user, dbname, password)

	fmt.Println("[GORM] Using DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database, got error %w", err)
	}

	if err := db.AutoMigrate(&models.UserModel{}); err != nil {
		return nil, err
	}

	return db, nil
}
