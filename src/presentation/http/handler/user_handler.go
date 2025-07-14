package handler

import (
	"net/http"

	"github.com/eyagovbusiness/GSWB.Users/pkg/logger"
	"github.com/eyagovbusiness/GSWB.Users/src/application/UseCases/user"
	"github.com/eyagovbusiness/GSWB.Users/src/application/dtos"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	CreateUserUseCase *user.CreateUserUseCase
	ListUsersUseCase  *user.ListUsersUseCase
}

func NewUserHandler(createUC *user.CreateUserUseCase, listUC *user.ListUsersUseCase) *UserHandler {
	return &UserHandler{
		CreateUserUseCase: createUC,
		ListUsersUseCase:  listUC,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var input dtos.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	user, err := h.CreateUserUseCase.Execute(ctx, input)
	if err != nil {
		logger.Logger.Error("failed to create user", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandler) ListUsers(ctx *gin.Context) {
	users, err := h.ListUsersUseCase.Execute(ctx)
	if err != nil {
		logger.Logger.Error("failed to list users", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
