package main

func FanIn[T any](c1, c2 <-chan T) <-chan T {
	resc := make(chan T)
	go func() {
		for {
			resc <- <-c1
		}
	}()
	go func() {
		for {
			resc <- <-c2
		}
	}()
	return resc
}
