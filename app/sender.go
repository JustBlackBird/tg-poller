package app

import (
	"bytes"
	"fmt"
	"net/http"
)

// Sender represents a gateway to send updates to backend.
type Sender struct {
	client *http.Client
	url    string
}

// NewSender creates a new sender struct
func NewSender(url string) *Sender {
	return &Sender{
		client: http.DefaultClient,
		url:    url,
	}
}

// SendUpdatesFrom takes all the updates from the passed in channel and sends
// them to the backend.
//
// The method is non-blocking and uses goroutines for any network activity.
func (s *Sender) SendUpdatesFrom(updates chan string) {
	go func() {
		for upd := range updates {
			go s.send(upd)
		}
	}()
}

func (s *Sender) send(update string) {
	res, err := s.client.Post(s.url, "application/json", bytes.NewBuffer([]byte(update)))
	if err != nil {
		fmt.Printf("Cannot send update to backend %v\n", err)
	}
	if res.StatusCode != 200 {
		fmt.Printf("Backend responded with code %v expected 200\n", res.StatusCode)
	}
}
