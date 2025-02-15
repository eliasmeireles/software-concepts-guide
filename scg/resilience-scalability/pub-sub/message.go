package pub_sub

import (
	"encoding/json"
)

type Message struct {
	Id   int
	Data string
}

func (m *Message) ToJson() (string, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// MessageFromJson creates a Message instance from a JSON string.
func MessageFromJson(jsonString string) (*Message, error) {
	var msg Message
	err := json.Unmarshal([]byte(jsonString), &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
