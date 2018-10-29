package main

import (
	"fmt"
	"github.com/learnergo/simplequeue/queue"
	"golang.org/x/net/context"
	"time"
)

var chanEnqueue = make(chan int)

func genkey(ctx context.Context) {
	i := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("genkey exit!")
				return
			default:
				chanEnqueue <- i
				i++
			}

		}
	}()
}

func enqueue(ctx context.Context, queue queue.Queue) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("enqueue exit")
				return
			case value := <-chanEnqueue:
				queue.Enqueue(value)
			default:
			}
		}
	}()
}

func dequeue(ctx context.Context, queue queue.Queue) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("dequeue exit")
				return
			default:
				if result, ok := queue.Dequeue(); ok {
					fmt.Println(result)
				}
			}

		}
	}()
}

func main() {
	queue := queue.NewQueue()

	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	ctx3, cancel3 := context.WithCancel(context.Background())
	genkey(ctx1)
	enqueue(ctx2, queue)
	dequeue(ctx3, queue)

	time.Sleep(2 * time.Second)
	cancel1()
	cancel2()
	cancel3()
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}
