package middleware

import (
	"net/http"

	"github.com/SharathKumarK06/todo-api/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Next()

    if len(c.Errors) > 0 {
      err := c.Errors.Last().Err

      if appErr, ok := err.(*utils.AppError); ok {
        c.JSON(appErr.StatusCode, gin.H{
          "success": false,
          "error":   appErr.Message,
        })
        return
      }

      c.JSON(http.StatusInternalServerError, gin.H{
        "success": false,
        "error":   "Internal server error",
      })
    }
  }
}
