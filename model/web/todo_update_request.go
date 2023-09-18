package web

type TodoUpdateRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,max=30,min=1"`
	Description string `validate:"required"`
	Status      string `validate:"required"`
}
