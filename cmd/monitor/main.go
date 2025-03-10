package main

import (
	"log"
	"time"

	"github.com/xuerbujia/clipboard"
)

func main() {
	changes := make(chan string, 10)
	stopCh := make(chan struct{})

	go clipboard.Monitor(time.Second, stopCh, changes)

	// Watch for changes
	for {
		select {
		case <-stopCh:
			break
		default:
			change, ok := <-changes
			if ok {
				log.Printf("change received: '%s'", change)
			} else {
				log.Printf("channel has been closed. exiting..")
			}
		}
	}
}
