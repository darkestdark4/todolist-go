package web

type TodoCreateRequest struct {
	Name        string `validate:"required,max=30,min=1"`
	Description string `validate:"required"`
	Status      string `validate:"required"`
}
