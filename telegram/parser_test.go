package telegram_test

import (
	"reflect"
	"testing"

	"github.com/justblackbird/tg-poller/telegram"
)

func TestParseUpdates(t *testing.T) {
	type testcase struct {
		name            string
		input           []byte
		expectedUpdates []string
		expectedLastID  int
	}

	tests := []testcase{
		{
			"Empty result",
			[]byte(`{"ok":true,"result":[]}`),
			make([]string, 0),
			0,
		},
		{
			"Single update",
			[]byte(`{"ok":true,"result":[{"update_id":123}]}`),
			[]string{`{"update_id":123}`},
			123,
		},
		{
			"Multiple updates",
			[]byte(`{"ok":true,"result":[{"update_id":123}, {"update_id":456}]}`),
			[]string{`{"update_id":123}`, `{"update_id":456}`},
			456,
		},
		{
			"Corrupted json",
			[]byte(`{"ok":true,"result":[{"update_id":123,"date":1592747949`),
			make([]string, 0),
			0,
		},
	}

	for _, test := range tests {
		updates, lastID := telegram.ParseUpdates(test.input)

		if !reflect.DeepEqual(updates, test.expectedUpdates) {
			t.Errorf("ParseUpdates FAILED on case \"%s\". Expected updates to be %v got %v", test.name, test.expectedUpdates, updates)
		}

		if lastID != test.expectedLastID {
			t.Errorf("ParseUpdates FAILED on case \"%s\". Expected ID to be %v got %v", test.name, test.expectedLastID, lastID)
		}
	}
}
