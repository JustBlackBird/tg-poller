package app

import (
	"fmt"

	"github.com/justblackbird/tg-poller/telegram"
)

// Poller can be used to poll updates from Telegram API
//
// After the poller is started it will watch for updates and pass them
// through Poller.Updates channel.
//
// Communication with telegram API is done via separated goroutine so
// the poller is truly no-blocking.
type Poller struct {
	Updates   chan string
	api       *telegram.API
	timeout   int
	done      chan bool
	cancelReq func()
}

// NewPoller creates new poller struct
func NewPoller(api *telegram.API) *Poller {
	return &Poller{
		api:     api,
		timeout: 10,
		Updates: make(chan string),
		done:    make(chan bool),
	}
}

// Start the poller watch for updates from Telegram API.
//
// The method does not block the process but uses a goroutine to watch updates.
//
// The poller can be stopped via Poller.Stop method.
func (p *Poller) Start() {
	go p.runLoop()
}

// Stop all the goroutins created by the poller and cleans up the memory.
func (p *Poller) Stop() {
	if p.cancelReq != nil {
		// Cancel request to the API
		p.cancelReq()
	}

	p.done <- true
}

func (p *Poller) runLoop() {
	for {
		select {
		default:
			updates, cancel := p.api.GetUpdates(p.timeout)

			p.cancelReq = cancel

			for _, upd := range <-updates {
				fmt.Printf("New update: %v\n", upd)
				p.Updates <- upd
			}
		case <-p.done:
			close(p.Updates)
			return
		}
	}
}
