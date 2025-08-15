package services

import (
	ws "github.com/nutgood/gome-assistant/internal/websocket"
	"github.com/nutgood/gome-assistant/types"
)

type Notify struct {
	conn *ws.WebsocketWriter
}

// Notify sends a notification. Takes a types.NotifyRequest.
func (ha *Notify) Notify(reqData types.NotifyRequest) error {
	req := NewBaseServiceRequest("")
	req.Domain = "notify"
	req.Service = reqData.ServiceName

	serviceData := map[string]any{}
	serviceData["message"] = reqData.Message
	serviceData["title"] = reqData.Title
	if reqData.Data != nil {
		serviceData["data"] = reqData.Data
	}

	req.ServiceData = serviceData
	return ha.conn.WriteMessage(req)
}
