package main

import (
	"github.com/itxor/tgsite/internal/tg"
	"os"
)

func main() {
	tgChannelService, err := tg.NewTelegramChannelService()
	if err != nil {
		os.Exit(1)
	}

	updatesChannel, err := tgChannelService.GetUpdates()
	if err != nil {
		os.Exit(1)
	}

	for update := range updatesChannel {
		err := tgChannelService.HandleMessage(update)
		if err != nil {
			os.Exit(1)
		}
	}
}

