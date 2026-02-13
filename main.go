package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/SharathKumarK06/todo-app/internal/database"
  "github.com/SharathKumarK06/todo-app/internal/todo"
)

func main() {
  db, _ := database.Connect()
  // Auto migrate: Create table automatically
  db.AutoMigrate(&todo.Todo{})

  r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "pong"})
  })

  r.Run(":8080")

  fmt.Println("Database connected:", db)
}
