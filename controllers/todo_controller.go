package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/SharathKumarK06/todo-app/config"
  "github.com/SharathKumarK06/todo-app/models"
  "github.com/SharathKumarK06/todo-app/utils"
)

func CreateTodo(c *gin.Context) {
  var todo models.Todo

  if err := c.ShouldBindJSON(&todo); err != nil {
    c.Error(utils.NewAppError(
      http.StatusBadRequest,
      "Invalid input data",
    ))
    c.Abort()
    return
  }

  config.DB.Create(&todo)
  c.JSON(http.StatusOK, todo)
}

func GetTodo(c *gin.Context) {
  var todos []models.Todo
  if err := config.DB.Find(&todos).Error; err != nil {
    c.Error(err)
    c.Abort()
    return
  }
  c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c *gin.Context) {
  id := c.Param("id")
  if err := config.DB.Delete(&models.Todo{}, id).Error; err != nil {
    c.Error(err)
    c.Abort()
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func UpdateTodo(c *gin.Context) {
  id := c.Param("id")

  var todo models.Todo
  if err := config.DB.First(&todo, id).Error; err != nil {
    c.Error(utils.NewAppError(
      http.StatusNotFound,
      "Todo not found",
    ))
    c.Abort()
    return
  }

  var input models.Todo
  if err := c.ShouldBindJSON(&input); err != nil {
    c.Error(err)
    c.Abort()
    return
  }

  todo.Title = input.Title
  todo.Completed = input.Completed

  config.DB.Save(&todo)
  c.JSON(http.StatusOK, todo)
}
