package main

import (
	"fmt"
	"time"
)

func main() {
}

func main01() {
	boring("boring!")
}

func main02() {
	go boringRand("boring!")
}

func main02_1() {
	go boringRand("boring!")
	fmt.Println("I'm listening!")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

func main03() {
	c := make(chan string)
	go boringSendToChan("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func main04() {
	c := boringGenerator("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func main04_1() {
	joe := boringGenerator("Joe")
	ann := boringGenerator("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func main05() {
	fanned := fanIn(boringGenerator("joe"), boringGenerator("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-fanned)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
