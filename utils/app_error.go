package utils

type AppError struct {
  StatusCode int
  Message    string
}

func (e *AppError) Error() string {
  return e.Message
}

func NewAppError(status int, message string) *AppError {
  return &AppError{
    StatusCode: status,
    Message: message,
  }
}
