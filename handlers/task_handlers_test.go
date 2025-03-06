package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "todo-app/db"
    "todo-app/models"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.POST("/tasks", AddTask)
    r.GET("/tasks", GetTasks)
    r.GET("/tasks/:id", GetTask)
    r.PUT("/tasks/:id", UpdateTask)
    r.DELETE("/tasks/:id", DeleteTask)
    return r
}

func TestAddTask(t *testing.T) {
    db.InitDB()
    router := setupRouter()
    task := models.Task{Title: "Test Task", Description: "Test Desc", Completed: false}
    jsonStr, _ := json.Marshal(task)
    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
    assert.Contains(t, w.Body.String(), `"title":"Test Task"`)
}

func TestAddTaskValidation(t *testing.T) {
    db.InitDB()
    router := setupRouter()
    task := models.Task{Title: "ab", Description: "Test Desc"}
    jsonStr, _ := json.Marshal(task)
    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Contains(t, w.Body.String(), `Field Title failed validation: min`)
}

func TestGetTasks(t *testing.T) {
    db.InitDB()
    router := setupRouter()
    db.DB.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", "Task 1", "Desc 1", false)
    req, _ := http.NewRequest("GET", "/tasks", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), `"title":"Task 1"`)
}
