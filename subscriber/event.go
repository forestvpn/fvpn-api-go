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

		// Handle the case where Data.Object is map[string]interface{}
		switch v := aux.Data.Object.(type) {
		case map[string]interface{}:
			// Re-marshal the map to JSON and then unmarshal into the correct type
			b, err := json.Marshal(v)
			if err != nil {
				return fmt.Errorf("failed to marshal object data: %w", err)
			}
			if err := json.Unmarshal(b, &obj); err != nil {
				return fmt.Errorf("failed to unmarshal node_device: %w", err)
			}
			e.Data.Object = &obj
		case []byte:
			// Handle the original case for backward compatibility
			if err := json.Unmarshal(v, &obj); err != nil {
				return fmt.Errorf("failed to unmarshal node_device from bytes: %w", err)
			}
			e.Data.Object = obj
		default:
			return fmt.Errorf("unexpected type for object data: %T", aux.Data.Object)
		}
	default:
		return fmt.Errorf("unsupported object type: %s", e.Object)
	}

	return nil
}
