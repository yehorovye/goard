package listeners

import (
	"fmt"

	"github.com/disgoorg/disgo/events"
)

func HandleReactionAdd(event *events.GuildMessageReactionAdd) {
	fmt.Printf("recieved reaction: %v", event)
}
