package util

import "fmt"

type Style int

const (
	StyleBlocks Style = iota
	StyleShaded
	StyleArrows
	StyleDots
	StyleMinimal
)

func Bar(current, total, width int, style Style) string {
	if total <= 0 {
		total = 1
	}
	if current > total {
		current = total
	}

	pct := float64(current) / float64(total)
	filled := int(pct * float64(width))
	empty := width - filled

	var filled_s, empty_s, open, close string

	switch style {
	case StyleBlocks:
		filled_s, empty_s, open, close = "█", "░", "[", "]"
	case StyleShaded:
		filled_s, empty_s, open, close = "▓", "░", "[", "]"
	case StyleArrows:
		filled_s, empty_s, open, close = "=", "-", "[", "]"
	case StyleDots:
		filled_s, empty_s, open, close = "●", "○", "|", "|"
	case StyleMinimal:
		filled_s, empty_s, open, close = "▉", " ", "▕", "▏"
	}

	bar := open
	if style == StyleArrows {
		if filled > 0 {
			for range filled - 1 {
				bar += filled_s
			}
			bar += ">"
		}
	} else {
		for range filled {
			bar += filled_s
		}
	}
	for range empty {
		bar += empty_s
	}
	bar += close

	return fmt.Sprintf("%s %d/%d (%d%%)", bar, current, total, int(pct*100))
}
