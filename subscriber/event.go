package subscriber

import (
	"encoding/json"
	"fmt"

	"github.com/forestvpn/fvpn-api-go"
)

type EventRequest struct {
	ID             string `json:"id,omitempty"`
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

type Event struct {
	ID         string        `json:"id,omitempty"`
	Object     string        `json:"object,omitempty"`
	APIVersion string        `json:"version,omitempty"`
	Type       string        `json:"type,omitempty"`
	Request    *EventRequest `json:"request,omitempty"`
	Data       *EventData    `json:"data,omitempty"`
}

type EventData struct {
	Object any `json:"object,omitempty"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch e.Object {
	case "node_device":
		var obj fvpn.NodeDevice
		if err := json.Unmarshal(aux.Data.Object.([]byte), &obj); err != nil {
			return err
		}
		e.Data.Object = obj
	default:
		return fmt.Errorf("unsupported object type: %s", e.Object)
	}

	return nil
}
