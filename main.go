package main

import (
	"fmt"
	"math/rand"
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

func main06() {
	joe := boringSequencingViaWaitChan("Joe")
	ann := boringSequencingViaWaitChan("Ann")
	fanned := fanIn(joe, ann)
	for i := 0; i < 5; i++ {
		msg1 := <-fanned
		fmt.Printf("msg1: %v\n", msg1.str)
		msg2 := <-fanned
		fmt.Printf("msg2: %v\n", msg2.str)
		// msg3 := <-fanned // will block
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring; I'm leaving.")
}

func main07() {
	fanned := fanInViaSelect(boringSequencingViaWaitChan("Joe"), boringSequencingViaWaitChan("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-fanned
		fmt.Printf("msg1: %v\n", msg1.str)
		msg2 := <-fanned
		fmt.Printf("msg2: %v\n", msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func main08() {
	joe := boringGenerator("joe")
	for {
		select {
		case v := <-joe:
			fmt.Println(v)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("You're too slow, Joe. I'm leaving.")
			return
		}
	}
}

func main09() {
	joe := boringGenerator("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-joe:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much. I'm leaving.")
			return
		}
	}
}

func main09_1() {
	c := fanInViaSelect(boringSequencingViaWaitChan("Joe"), boringSequencingViaWaitChan("Ann"))
	timeout := time.After(5 * time.Second)
	for {
		select {
		case m := <-c:
			fmt.Println(m.str)
			m.wait <- true
		case <-timeout:
			fmt.Println("You both talk too much. I'm leaving.")
			return
		}
	}
}

func main10() {
	quit := make(chan bool)
	joe := boringWithQuit("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-joe)
	}
	quit <- true
}

func main11() {
	quit := make(chan string)
	joe := boringWithQuitAndNotifyWhenCleanupIsDone("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-joe)
	}
	fmt.Println("Saying Joe to stop!")
	quit <- "You should stop!"
	fmt.Printf("Joe says: %s\n", <-quit)
}

func main12() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go mapConnect(left, right, func(v int) int { return v + 1 })
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

func main13() {
	start := time.Now()
	results := Google1_0("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func main14() {
	start := time.Now()
	results := Google2_0("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func main15() {
	start := time.Now()
	results := Google2_1("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func main16() {
	start := time.Now()
	results := First("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func main17() {
	start := time.Now()
	results := Google3_0("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
