package telegram

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// API represents telegram API with some handy methods to retrieve data.
type API struct {
	Host         string
	Token        string
	lastUpdateID int
	client       *http.Client
}

// NewAPI Creates new API struct
func NewAPI(host string, token string) *API {
	api := &API{Host: host, Token: token}

	api.client = http.DefaultClient

	return api
}

// GetUpdates retrive all pending updates from Telegram API.
// The function returns a channel which will be used once updates are loaded
// and a function which can be used by the client code to cancel the request.
func (api *API) GetUpdates(timeout int) (chan []string, func()) {
	// Increase timeout by one second to make sure API retrieves the answer
	// before network connection is timmed out.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout+1)*time.Second)

	url := api.updatesURL(timeout, api.lastUpdateID)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	res := make(chan []string)
	go func() {
		resp, err := api.client.Do(req)
		if err != nil {
			// Just ignore the errors. It could be temporary problems that will
			// be fixed up for the next request.
			fmt.Printf("Telegram API error: %v\n", err)
			res <- make([]string, 0)

			return
		}

		body, _ := ioutil.ReadAll(resp.Body)

		updates, lastID := ParseUpdates(body)
		api.lastUpdateID = lastID

		res <- updates
		close(res)
	}()

	return res, cancel
}

func (api *API) updatesURL(timeout int, lastUpdate int) string {
	url := "https://" + api.Host + "/bot" + api.Token + "/getUpdates?timeout=" + strconv.Itoa(timeout)

	if lastUpdate != 0 {
		url = url + "&offset=" + strconv.Itoa(lastUpdate+1)
	}

	return url
}
