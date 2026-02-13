package database

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
  dsn := "host=localhost user=todo password=pass dbname=todo_db port=5432 sslmode=disable"
  return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}


