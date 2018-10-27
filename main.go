package main

import (
	"fmt"
	"github.com/learnergo/simplequeue/queue"
	"time"
)

var chanEnqueue = make(chan int)

func genkey() {
	i := 0
	go func() {
		for {
			chanEnqueue <- i
			i++
		}
	}()
}

func enqueue(queue queue.Queue) {
	go func() {
		for {
			select {
			case value := <-chanEnqueue:
				queue.Enqueue(value)
			default:
			}
		}
	}()
}

func dequeue(queue queue.Queue) {
	go func() {
		for {
			if result, ok := queue.Dequeue(); ok {
				fmt.Println(result)
			}
		}
	}()
}

func main() {
	queue := queue.NewQueue()

	genkey()
	enqueue(queue)
	dequeue(queue)

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
