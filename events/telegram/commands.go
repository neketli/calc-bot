package telegram

import (
	"log"
	"regexp"
	"strings"
)

const (
	StartCommand = "/start"
	HelpCommand  = "/help"
)

func (p *Processor) handle(text string, ChatID int, username string) {
	text = strings.TrimSpace(text)

	log.Printf("DEBUG: new command '%s' from '%s'", text, username)

	if isMathOperation(text) {
		// ...
	}

	switch text {
	case HelpCommand:
	case StartCommand:
	default:
	}
}

func isMathOperation(text string) bool {
	// Regular expression that show us if text is basic math expression
	return regexp.MustCompile(`[0-9]+[-+\/:*][0-9]+`).MatchString(text)
}
