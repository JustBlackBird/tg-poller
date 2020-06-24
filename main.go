package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/justblackbird/tg-poller/app"
	"github.com/justblackbird/tg-poller/telegram"
)

func main() {
	host := os.Getenv("TELEGRAM_API_HOST")
	if host == "" {
		host = "api.telegram.org"
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		panic("Env variable \"TELEGRAM_TOKEN\" must be defined")
	}

	updateHook := os.Getenv("UPDATE_HOOK")
	if updateHook == "" {
		panic("Env variable \"UPDATE_HOOK\" must be defined")
	}

	fmt.Println("App is running with env:")
	fmt.Printf("TELEGRAM_API_HOST = %v\n", host)
	fmt.Printf("TELEGRAM_TOKEN = %v\n", token)
	fmt.Printf("UPDATE_HOOK = %v\n", updateHook)

	poller := app.NewPoller(telegram.NewAPI(host, token))

	sender := app.NewSender(updateHook)
	sender.SendUpdatesFrom(poller.Updates)

	poller.Start()

	fmt.Println("App is started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	// Wait for a signal to stop.
	<-quit

	poller.Stop()

	fmt.Println("App is stopped")
}
