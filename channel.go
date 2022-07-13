package common

import "context"

func SendInNewRT[T any](ctx context.Context, t T, c chan<- T) {
	go func() {
		select {
		case <-ctx.Done():
			return
		case c <- t:
			return
		}
	}()
}
