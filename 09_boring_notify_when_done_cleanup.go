package main

import (
	"fmt"
	"time"
)

func BoringWithQuitAndNotifyWhenCleanupIsDone(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit:
				fmt.Println("Cleanup is done!")
				quit <- "I'm done!"
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return c
}
