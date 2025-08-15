package services

import (
	ws "github.com/nutgood/gome-assistant/internal/websocket"
)

type Number struct {
	conn *ws.WebsocketWriter
}

func (ib Number) SetValue(entityId string, value float32) error {
	req := NewBaseServiceRequest(entityId)
	req.Domain = "number"
	req.Service = "set_value"
	req.ServiceData = map[string]any{"value": value}

	return ib.conn.WriteMessage(req)
}

func (ib Number) MustSetValue(entityId string, value float32) {
	if err := ib.SetValue(entityId, value); err != nil {
		panic(err)
	}
}
