package main

func MapConnect[T, U any](left chan<- U, right <-chan T, mapper func(T) U) {
	left <- mapper(<-right)
}
