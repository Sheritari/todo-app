package main

import (
    "log"
    "todo-app/db"
    "todo-app/handlers"
    "github.com/gin-gonic/gin"
)

func main() {

    if err := db.InitDB(); err != nil {
        log.Fatal("Failed to initialize database:", err)
    }

    r := gin.Default()
    r.POST("/tasks", handlers.AddTask)
    r.GET("/tasks", handlers.GetTasks)
    r.GET("/tasks/:id", handlers.GetTask)
    r.PUT("/tasks/:id", handlers.UpdateTask)
    r.DELETE("/tasks/:id", handlers.DeleteTask)

    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}
