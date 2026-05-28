package entities

import "time"

type Task struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	ProjectID   *int       `json:"project_id"` // Puede ser null
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`   // Puede ser null
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
