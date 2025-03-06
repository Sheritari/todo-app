package handlers

import (
    "net/http"
    "strconv"
    "todo-app/db"
    "todo-app/models"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

// POST /tasks
func AddTask(c *gin.Context) {
    var task models.Task

    if err := c.ShouldBindJSON(&task); err != nil {
        var validationErrors []string

        if ve, ok := err.(validator.ValidationErrors); ok {
            for _, fieldErr := range ve {
                validationErrors = append(validationErrors,
                    "Field " + fieldErr.Field() + " failed validation: " + fieldErr.Tag())
            }
        } else {
            validationErrors = append(validationErrors, err.Error())
        }

        c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
        return
    }

    result, err := db.DB.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)",
        task.Title, task.Description, task.Completed)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
        return
    }

    id, _ := result.LastInsertId()
    task.ID = int(id)
    c.JSON(http.StatusCreated, task)
}

// PUT /tasks/{id}
func UpdateTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var task models.Task

    if err := c.ShouldBindJSON(&task); err != nil {
        var validationErrors []string

        if ve, ok := err.(validator.ValidationErrors); ok {
            for _, fieldErr := range ve {
                validationErrors = append(validationErrors,
                    "Field " + fieldErr.Field() + " failed validation: " + fieldErr.Tag())
            }
        } else {
            validationErrors = append(validationErrors, err.Error())
        }

        c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
        return
    }

    result, err := db.DB.Exec("UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?",
        task.Title, task.Description, task.Completed, id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
        return
    }

    rowsAffected, _ := result.RowsAffected()

    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    task.ID = id
    c.JSON(http.StatusOK, task)
}

// GET /tasks
func GetTasks(c *gin.Context) {
    rows, err := db.DB.Query("SELECT id, title, description, completed FROM tasks")

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    defer rows.Close()

    var tasks []models.Task

    for rows.Next() {
        var task models.Task

        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        tasks = append(tasks, task)
    }

    c.JSON(http.StatusOK, tasks)
}

// GET /tasks/{id}
func GetTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var task models.Task
    err := db.DB.QueryRow("SELECT id, title, description, completed FROM tasks WHERE id = ?", id).
        Scan(&task.ID, &task.Title, &task.Description, &task.Completed)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, task)
}

// DELETE /tasks/{id}
func DeleteTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    result, err := db.DB.Exec("DELETE FROM tasks WHERE id = ?", id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    rowsAffected, _ := result.RowsAffected()

    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.Status(http.StatusNoContent)
}
