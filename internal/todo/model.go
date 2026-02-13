package todo

import "time"

type Todo struct {
  ID        uint  `gorm:"primary key"`
  Title     string `gorm:"not null"`
  Completed bool `gorm:"default:false"`
  CreatedAt time.Time
  UpdatedAt time.Time
}
