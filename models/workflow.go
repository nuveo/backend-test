package models

import "encoding/json"

type Workflow struct {
	UUID   string          `json:"uuid"`
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
	Steps  []string        `json:"steps"`
}
