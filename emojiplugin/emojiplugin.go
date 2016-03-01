package emojiplugin

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iopred/bruxism"
)

func emojiMessageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.Name() == bruxism.DiscordServiceName && !service.IsMe(message) {
		if bruxism.MatchesCommand(service, "emoji", message) {
			_, parts := bruxism.ParseCommand(service, message)
			if len(parts) == 1 {
				s := strings.TrimSpace(parts[0])
				for _, r := range s {
					if f, err := os.Open(fmt.Sprintf("emoji/twitter/%s.png", strconv.FormatInt(int64(r), 16))); err == nil {
						defer f.Close()
						service.SendFile(message.Channel(), "emjoi.png", f)
					}
					return
				}
			}

		}
	}
}

func emojiHelpFunc(bot *bruxism.Bot, service bruxism.Service, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "emoji", "<emoji>", "Returns a big version of an emoji.")
}

// NewEmojiPlugin creates a new emoji plugin.
func NewEmojiPlugin() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("Emoji")
	p.MessageFunc = emojiMessageFunc
	p.HelpFunc = emojiHelpFunc
	return p
}
