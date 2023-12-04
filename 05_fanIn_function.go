package main

func fanIn[T any](c1 <-chan T, c2 <-chan T) <-chan T {
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
