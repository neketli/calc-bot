package telegram

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	StartCommand = "/start"
	HelpCommand  = "/help"
)

func (p *Processor) handle(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("DEBUG: new command '%s' from '%s'", text, username)

	if isMathOperation(text) {
		switch {
		case strings.Contains(text, "+"):
			result, err := adderString(text)
			if err != nil {
				log.Printf("ERROR: %s", err)
				p.tg.SendMessage(chatID, msgUnknown)
				return err
			}
			// Trims .0000
			result = strings.TrimRight(strings.TrimRight(result, "0"), ".")
			return p.tg.SendMessage(chatID, result)
		case strings.Contains(text, "-"):
			result, err := suberString(text)
			if err != nil {
				log.Printf("ERROR: %s", err)
				p.tg.SendMessage(chatID, msgUnknown)
				return err
			}
			// Trims .0000
			result = strings.TrimRight(strings.TrimRight(result, "0"), ".")
			return p.tg.SendMessage(chatID, result)
		case strings.Contains(text, "*"):
			result, err := multerString(text)
			if err != nil {
				log.Printf("ERROR: %s", err)
				p.tg.SendMessage(chatID, msgUnknown)
				return err
			}
			// Trims .0000
			result = strings.TrimRight(strings.TrimRight(result, "0"), ".")
			return p.tg.SendMessage(chatID, result)
		case strings.Contains(text, "/"):
			result, err := diverString(text)
			if err != nil {
				log.Printf("ERROR: %s", err)
				p.tg.SendMessage(chatID, msgUnknown)
				return err
			}
			// Trims .0000
			result = strings.TrimRight(strings.TrimRight(result, "0"), ".")
			return p.tg.SendMessage(chatID, result)
		default:
			return p.tg.SendMessage(chatID, msgUnknown)
		}
	}

	switch text {
	case HelpCommand:
		return p.tg.SendMessage(chatID, msgHelp)
	case StartCommand:
		return p.tg.SendMessage(chatID, msgHello)
	default:
		return p.tg.SendMessage(chatID, msgUnknown)
	}
}

func adderString(text string) (string, error) {
	res := strings.Split(text, "+")
	x1, x2, err := argsParser(res)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", x1+x2), nil
}

func suberString(text string) (string, error) {
	res := strings.Split(text, "-")
	x1, x2, err := argsParser(res)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", x1-x2), nil
}

func multerString(text string) (string, error) {
	res := strings.Split(text, "*")
	x1, x2, err := argsParser(res)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", x1*x2), nil
}

func diverString(text string) (string, error) {
	res := strings.Split(text, "/")
	x1, x2, err := argsParser(res)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", x1/x2), nil
}

func argsParser(arr []string) (float64, float64, error) {
	x1, err := strconv.ParseFloat(arr[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse 1st argument: %w", err)
	}
	x2, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse 2st argument: %w", err)
	}
	return x1, x2, nil
}

func isMathOperation(text string) bool {
	// Regular expression that show us if text is basic math expression
	return regexp.MustCompile(`[0-9]+\.?[0-9]*[-+\/*][0-9]+\.?[0-9]*`).MatchString(text)
}
