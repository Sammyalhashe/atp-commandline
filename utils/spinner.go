package utils

import (
	"fmt"
	"time"

	spin "github.com/tj/go-spin"
)

// StartLoading starts a spinning loader that stops when the channel it recieves as one of its arguments sends a stop signal
func StartLoading(c <-chan bool) {
	s := spin.New()
	// s.Set("Spin5")
	for {
		select {
		case <-c:
			return
		default:
			fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
			time.Sleep(100 * time.Millisecond)
		}
	}
}
