package main

import (
  "github.com/SharathKumarK06/todo-app/config"
  "github.com/SharathKumarK06/todo-app/models"
  "github.com/SharathKumarK06/todo-app/routes"
  "github.com/SharathKumarK06/todo-app/middleware"

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
