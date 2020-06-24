package telegram

import (
	"encoding/json"
)

// ParseUpdates extracts updates from telegram API response
func ParseUpdates(data []byte) ([]string, int) {
	var parsed interface{}

	err := json.Unmarshal(data, &parsed)
	if err != nil {
		// Silently ignore json errors
		return make([]string, 0), 0
	}

	updates, ok := parsed.(map[string]interface{})["result"]
	if !ok {
		// By some reasons there is no updates in the message
		return make([]string, 0), 0
	}

	res := make([]string, 0)
	lastID := 0

	for _, upd := range updates.([]interface{}) {
		if rawID, ok := upd.(map[string]interface{})["update_id"]; ok {
			str, _ := json.Marshal(upd)
			res = append(res, string(str))

			if id := int(rawID.(float64)); id > lastID {
				lastID = id
			}
		}
	}

	return res, lastID
}
