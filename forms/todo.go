package forms

type ToDoData struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsDone      bool   `json:"is_done" binding:"required"`
}
