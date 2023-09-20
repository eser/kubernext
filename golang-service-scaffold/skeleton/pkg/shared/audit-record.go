package shared

import (
	"time"
)

type AuditRecord struct {
	RequestedBy string    `json:"by"`
	RequestedAt time.Time `json:"at"`
}
