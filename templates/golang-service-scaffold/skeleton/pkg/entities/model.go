package entities

import (
  "time"
)

type Model struct {
  Id          string      `json:"id"`

  Fullname    string      `json:"fullname"`

  CreatedBy   string      `json:"created_by"`
  CreatedAt   time.Time   `json:"created_at"`
  UpdatedBy   string      `json:"updated_by"`
  UpdatedAt   time.Time   `json:"updated_at"`
  RemovedBy   string      `json:"removed_by"`
  RemovedAt   time.Time   `json:"removed_at"`
}
