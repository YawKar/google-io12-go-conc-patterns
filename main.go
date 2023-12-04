package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
}

func Main01() {
	Boring("boring!")
}

func Main02() {
	go BoringRand("boring!")
}

func Main02_1() {
	go BoringRand("boring!")
	fmt.Println("I'm listening!")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

func Main03() {
	c := make(chan string)
	go BoringSendToChan("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func Main04() {
	c := BoringGenerator("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func Main04_1() {
	joe := BoringGenerator("Joe")
	ann := BoringGenerator("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func Main05() {
	fanned := FanIn(BoringGenerator("joe"), BoringGenerator("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-fanned)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func Main06() {
	joe := BoringSequencingViaWaitChan("Joe")
	ann := BoringSequencingViaWaitChan("Ann")
	fanned := FanIn(joe, ann)
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

func Main07() {
	fanned := FanInViaSelect(BoringSequencingViaWaitChan("Joe"), BoringSequencingViaWaitChan("Ann"))
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

func Main08() {
	joe := BoringGenerator("joe")
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

func Main09() {
	joe := BoringGenerator("Joe")
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

func Main09_1() {
	c := FanInViaSelect(BoringSequencingViaWaitChan("Joe"), BoringSequencingViaWaitChan("Ann"))
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

func Main10() {
	quit := make(chan bool)
	joe := BoringWithQuit("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-joe)
	}
	quit <- true
}

func Main11() {
	quit := make(chan string)
	joe := BoringWithQuitAndNotifyWhenCleanupIsDone("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-joe)
	}
	fmt.Println("Saying Joe to stop!")
	quit <- "You should stop!"
	fmt.Printf("Joe says: %s\n", <-quit)
}

func Main12() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go MapConnect(left, right, func(v int) int { return v + 1 })
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

func Main13() {
	start := time.Now()
	results := Google1_0("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func Main14() {
	start := time.Now()
	results := Google2_0("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func Main15() {
	start := time.Now()
	results := Google2_1("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func Main16() {
	start := time.Now()
	results := First("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func Main17() {
	start := time.Now()
	results := Google3_0("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
