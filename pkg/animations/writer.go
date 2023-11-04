package animations

import (
	"context"
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

type Writer interface {
	Loading(context.Context, string)
	Clear()
}

type PlainWriter struct{}

func NewPlainWriter() *PlainWriter {
	return &PlainWriter{}
}

func (w *PlainWriter) Loading(ctx context.Context, message string) {
	go func() {

		chars := []rune{
			'⡿', '⣟', '⣯',
			'⣷', '⣾', '⣽',
			'⣻', '⢿',
		}
		state := 0
		limit := len(chars)

		framesPerSecond := 24
		frameDuration := time.Second / time.Duration(framesPerSecond)

		for {
			select {
			// FIXME: loading new line
			case <-ctx.Done():
				return

			default:
				cursor.StartOfLine()
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
	}()
}

func (w *PlainWriter) Clear() {
	cursor.ClearLine()
	cursor.StartOfLine()
	cursor.Show()
}
