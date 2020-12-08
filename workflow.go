package main

const (
	INSERTED Status = "inserted"
	CONSUMED Status = "consumed"
)

type Status string

func (s Status) IsValid() bool {
	return s == INSERTED || s == CONSUMED
}

type Workflow struct {
	UUID   string      `json:"uuid"`
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
	Steps  interface{} `json:"steps"`
}
