package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BoringSendToChan(msg string, c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
