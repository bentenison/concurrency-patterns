package faninout

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func FanIn() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	urls := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for letter := range fanin(ctx, 3, generator(ctx, urls...), generator(ctx, urls...)) {
		log.Println("recieved from fan in", letter)
	}

}
func fanin(ctx context.Context, maxWorkers int, in ...<-chan interface{}) <-chan interface{} {
	// NOTE: We can add the buffered channels here, but it is always better to decouple logic
	outStream := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(in))

	// Multiplex is just a function that processes each channels separately,
	// and sending them back to a single out channel
	multiplex := func(worker int, inStream <-chan interface{}) {
		defer wg.Done()

		for i := range inStream {
			// Fake processing time
			time.Sleep(time.Duration(rand.Intn(250)+250) * time.Millisecond)
			select {
			case <-ctx.Done():
				return
			case outStream <- i:
				log.Println("worker:", worker)
				// log.Printf("fanout queue: %d/%d\n", len(outStream), cap(outStream))
			}
		}
	}

	// For each channel, run a separate process simulating
	// multiple workers
	// for i := 0; i < maxWorkers; i++ {
	// 	go multiplex(i, in)
	// }
	for i, c := range in {
		go multiplex(i, c)
	}

	go func() {
		// Wait for all the workers to be completed before closing the stream
		defer close(outStream)
		wg.Wait()
	}()

	return outStream
}
func generator(ctx context.Context, in ...string) <-chan interface{} {
	log.Println(in)
	outStream := make(chan interface{})

	go func() {
		defer close(outStream)
		for _, v := range in {
			select {
			// Ensure there are no goroutines leak
			case <-ctx.Done():
				return
			case outStream <- v:
			}
		}
	}()

	return outStream
}
