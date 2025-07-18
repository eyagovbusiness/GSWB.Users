package dtos

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
