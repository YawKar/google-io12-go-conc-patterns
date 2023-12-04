package main

import (
	"fmt"
	"time"
)

var (
	Web1   = fakeSearch("Web1")
	Web2   = fakeSearch("Web2")
	Image1 = fakeSearch("Image1")
	Image2 = fakeSearch("Image2")
	Video1 = fakeSearch("Video1")
	Video2 = fakeSearch("Video2")
)

func Google3_0(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}
