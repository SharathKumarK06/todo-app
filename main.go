package main

import (
  "github.com/SharathKumarK06/todo-api/config"
  "github.com/SharathKumarK06/todo-api/models"
  "github.com/SharathKumarK06/todo-api/routes"
  "github.com/SharathKumarK06/todo-api/middleware"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.Use(middleware.ErrorHandler())

  config.ConnectDB()
  config.DB.AutoMigrate(&models.Todo{})

  routes.SetupRouter(r)

  r.Run(":8080")
}
