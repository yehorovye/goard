package listeners

import (
	"log/slog"

	"github.com/disgoorg/disgo/events"
)

func HandleReady(event *events.Ready) {
	slog.Info("goard bot is now ready")
}
