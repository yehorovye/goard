package main

import (
	"context"
	"goard/bot/listeners"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
)

// i don't initialize .env files so remember to pass the token on command invokation lol
var token string = os.Getenv("DISCORD_TOKEN")

func main() {
	client, err := disgo.New(token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentMessageContent,
				gateway.IntentGuildMessageReactions,
			),
		),
		bot.WithEventListenerFunc(listeners.HandleReactionAdd),
		bot.WithEventListenerFunc(listeners.HandleReady),
	)

	if err != nil {
		panic(err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	slog.Info("goard process is running")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
