package models

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title" binding:"required,min=3,max=100"`
    Description string `json:"description" binding:"max=500"`
    Completed   bool   `json:"completed"`
}
