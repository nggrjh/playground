package datastructures

import (
	"fmt"
)

const (
	bracesOpening      = "{"
	bracesClosing      = "}"
	bracketsOpening    = "["
	bracketsClosing    = "]"
	parenthesesOpening = "("
	parenthesesClosing = ")"
)

var bracesOf = map[string]string{
	bracesOpening:      bracesClosing,
	bracketsOpening:    bracketsClosing,
	parenthesesOpening: parenthesesClosing,
}

func lint(command string) error {
	var openings []string

	for _, char := range command {
		switch syntax := string(char); syntax {
		case bracesOpening, bracketsOpening, parenthesesOpening:
			openings = append(openings, syntax)
		case bracesClosing, bracketsClosing, parenthesesClosing:
			if len(openings) < 1 {
				return fmt.Errorf("%s does not have an opening brace", syntax)
			}

			if bracesOf[openings[len(openings)-1]] != syntax {
				return fmt.Errorf("%s does not have an opening brace", syntax)
			}

			openings = append([]string{}, openings[:len(openings)-1]...)
		default:
			break
		}
	}

	if len(openings) > 0 {
		return fmt.Errorf("%s does not have a closing brace", openings[len(openings)-1])
	}

	return nil
}
