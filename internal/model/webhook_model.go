package webhook

import (
	"time"

	"gorm.io/datatypes"
)

type WebhookModel struct {
	ID         string         `gorm:"primaryKey" json:"id"`
	EndpointID string         `json:"endpoint_id"`
	OrderID    string         `gorm:"uniqueIndex" json:"order_id"`
	Event      string         `json:"event"`
	Payload    datatypes.JSON `json:"payload"`
	Status     string         `json:"status"`
	Retry      int            `json:"retry"`
	CreatedAt  time.Time      `json:"created_at"`
}