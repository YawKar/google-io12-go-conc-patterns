package main

func FanInViaSelect[T any](c1, c2 <-chan T) <-chan T {
	c := make(chan T)
	go func() {
		for {
			select {
			case v := <-c1:
				c <- v
			case v := <-c2:
				c <- v
			}
		}
	}()
	return c
}
