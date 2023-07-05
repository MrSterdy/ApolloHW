package bot

import (
	"os"

	"github.com/MrSterdy/ApolloHW/config"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var Bot *gotgbot.Bot

func Start() {
	bot, err := gotgbot.NewBot(os.Getenv("BOT_TOKEN"), nil)
	if err != nil {
		panic(err)
	}

	Bot = bot

	updater := ext.NewUpdater(nil)

	dispatcher := updater.Dispatcher
	dispatcher.AddHandler(handlers.NewCommand("start", func(bot *gotgbot.Bot, ctx *ext.Context) error {
		_, err = ctx.EffectiveMessage.Reply(bot, config.File.Messages.Commands.Start, nil)

		return err
	}))

	err = updater.StartPolling(Bot, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic(err)
	}
}
