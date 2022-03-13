package model

import (
	"encoding/json"
	"fmt"
)

type Text struct {
	UUID      string `json:"uuid"`
	Timestamp int64  `json:"timestamp"`
	Content   string `json:"content"`
}

func (t *Text) Stringify() string {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}
	return string(b)
}

func (t *Text) Parse(s string) {
	json.Unmarshal([]byte(s), t)
}
