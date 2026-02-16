package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/SharathKumarK06/todo-api/controllers"
)

func SetupRouter(r *gin.Engine) {
  r.POST("/todos", controllers.CreateTodo)
  r.GET("/todos", controllers.GetTodo)
  r.PUT("/todos/:id", controllers.UpdateTodo)
  r.DELETE("/todos/:id", controllers.DeleteTodo)
}
