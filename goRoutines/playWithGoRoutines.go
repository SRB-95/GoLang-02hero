package goRoutines

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
		time.Sleep(time.Second * 2)
	}
}

func exploreGoRoutines() {
	// func main() {
	fmt.Println(time.Now().Format(time.RFC850))

	// doSomething("Oye Hoyee!!!")          // Blocking Function Call
	// doSomething("Ye Dil Maange More!!!") // Blocking Function Call

	go doSomething("Oye Hoyee!!!")          // Non Blocking Function Call
	go doSomething("Ye Dil Maange More!!!") // Non Blocking Function Call

	// someClosure := func( message string ) {
	// func( message string ) { // Blocking Function Call
	go func(message string) { // Non Blocking Function Call
		for i := 0; i < 3; i++ {
			fmt.Println(message, " : ", i)
			time.Sleep(time.Second * 2)
		}
		time.Sleep(time.Second * 5)
	}("Balleee Balleee!!")
	//}
	//go someClosure("Balleee Balleee!!")

	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println("Done: playWithGoRoutines")
}

func playWithChannels() {
	messages := make(chan string)
	var messageRead string

	// <- Symbol Behind Channel Means Writing
	// <- Symbol Before Channel Means Reading
	go func() {
		//	Writing To Channel messages
		messages <- "Ping"
		time.Sleep(time.Second * 2)

		messages <- "Pong"
		time.Sleep(time.Second * 2)

		messages <- "Ting"
		time.Sleep(time.Second * 2)

		messages <- "Tong"
		time.Sleep(time.Second * 2)
		fmt.Println("Done: Closure Go Routine")
	}()

	// Reading From Channel messages
	messageRead = <-messages
	fmt.Println(messageRead)

	messageRead = <-messages
	fmt.Println(messageRead)

	messageRead = <-messages
	fmt.Println(messageRead)

	messageRead = <-messages
	fmt.Println(messageRead)

	fmt.Println("Done: playWithChannels")
}

func sum(numbers []int, message chan int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	message <- sum
}

func sumOfArrays() {
	numbers := []int{1, 2, 9, 4}
	messages := make(chan int)

	go sum(numbers[:len(numbers)/2], messages)
	go sum(numbers[len(numbers)/2:], messages)

	// The data flows in direction of the arrow
	sum1, sum2 := <-messages, <-messages

	total := sum1 + sum2
	fmt.Println("Total Sum:", total)
}

// pings Is Write Only Channel
func ping(pings chan<- string, data string) {
	pings <- data
}

// pings Is Read Only Channel
// pongs Is Write Only Channel
func pong(pings <-chan string, pongs chan<- string) {
	data := <-pings
	pongs <- data
}

func playWithChannelWithReadWriteOnly() {
	pings := make(chan string, 1)
	fmt.Printf("\npings:::::->: %T", pings)

	pongs := make(chan string, 1)
	fmt.Println("pongs:::::->", pongs)

	ping(pings, "Ye Dil Maange More!!!")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}

func worker(done chan bool) {
	fmt.Println("Worker: Working...")
	time.Sleep(time.Second * 3)
	fmt.Println("Worker: Work Done!")

	done <- true
}

func playWithWorkers() {
	done := make(chan bool, 1)
	go worker(done)

	workerStatus := <-done
	fmt.Println(workerStatus)
}

func playWithClosingChannel() {
	messages := make(chan string, 2)
	messages <- "Ye Dil Maangee More!!!"
	messages <- "Oyee Hoyeee!!!"

	close(messages)

	// Reading From Channel Using For Loop
	for message := range messages {
		fmt.Println(message)
	}
}

func fibonacci(count int, fibos chan int) {
	x, y := 0, 1

	for i := 0; i < count; i++ {
		fibos <- x
		x, y = y, x+y
	}
	close(fibos)
}

func playWithFibonaccis() {
	fibos := make(chan int, 10)

	go fibonacci(cap(fibos), fibos)

	for fibo := range fibos {
		fmt.Println(fibo)
	}
}

func fibonacciAgain(fibos chan int, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case fibos <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Fibonacci Quit..!")
			return
		}
	}
}

func playWithfibonaccsiAgain() {
	fibos := make(chan int, 10)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-fibos)
		}
		quit <- 0
	}()
	fibonacciAgain(fibos, quit)
}

func playWithTimeTick() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func playWithSafeCounter() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

func PlayWithGoRoutines() {
	fmt.Println("Function: playWithStacks")
	playWithStacks()
	fmt.Println("---------------->")

	fmt.Println("Function: explore go routines")
	// exploreGoRoutines()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithChannels")
	// playWithChannels()
	// go playWithChannels()
	fmt.Println("---------------->")

	fmt.Println("Function: explore go routines")
	// exploreGoRoutines()
	fmt.Println("---------------->")

	fmt.Println("Function: Sum of Arrays")
	sumOfArrays()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithChannelWithReadWriteOnly")
	playWithChannelWithReadWriteOnly()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithWorkers")
	// playWithWorkers()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithClosingChannel")
	playWithClosingChannel()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithFibonaccis")
	playWithFibonaccis()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithfibonaccsiAgain")
	playWithfibonaccsiAgain()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithTimeTick")
	// playWithTimeTick()
	fmt.Println("---------------->")

	fmt.Println("Function: playWithSafeCounter")
	playWithSafeCounter()
	fmt.Println("---------------->")

	time.Sleep(time.Second * 4)
}
