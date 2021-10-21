package model

type Todo struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

type CreateTodoRequest struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Task string `json:"task"`
}

type UpdateTodoRequest struct {
	Status bool `json:"status"`
}
