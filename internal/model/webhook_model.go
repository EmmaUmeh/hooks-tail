package webhook

import "time"

type WebhookDelivery struct {
    ID          string    `json:"id"`
    EndpointID  string    `json:"endpoint_id"`
    OrderID     string    `json:"order_id"`
    Event       string    `json:"event"`
    Payload     any       `json:"payload"`
    Status      string    `json:"status"`
    retry    int       `json:"retry"`
    CreatedAt   time.Time `json:"created_at"`
}