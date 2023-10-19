package dmtask

import "time"

type Task struct {
	ID          int32     `json:"id"`
	UserID      int32     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
