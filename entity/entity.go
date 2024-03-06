package entity

type Status int

const (
	TaskCompleted  Status = 1
	TaskIncomplete Status = 0
	IdJsonName     string = "id"
	NameJsonName   string = "name"
	StatusJsonName string = "status"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}
