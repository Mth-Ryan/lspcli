package animations

import (
	"context"
	"fmt"
	"time"
)

type Writer interface {
	Loading(context.Context, string)
}

type PlainWriter struct{}

func NewPlainWriter() *PlainWriter {
	return &PlainWriter{}
}

func screenClear() {
	fmt.Print("\r")
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func (w *PlainWriter) Loading(ctx context.Context, message string) {
	chars := []rune{
		'⡿', '⣟', '⣯',
		'⣷', '⣾', '⣽',
		'⣻', '⢿',
	}
	state := 0
	limit := len(chars)

	framesPerSecond := 24
	frameDuration := time.Second / time.Duration(framesPerSecond)

	hideCursor()
	for {
		select {
		// FIXME: loading new line
		case <-ctx.Done():
			screenClear()
			showCursor()
			fmt.Println()
			return

		default:
			screenClear()
			fmt.Printf("%c %s", chars[state], message)

			timeToSleep := frameDuration - time.Since(time.Now())

			if timeToSleep > 0 {
				time.Sleep(timeToSleep)
			}

			if state+1 < limit {
				state += 1
			} else {
				state = 0
			}
			continue
		}
	}
}
