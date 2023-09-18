package web

type TodoResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
