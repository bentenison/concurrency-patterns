package faninout

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func FanoutSemaphore() {
	work := 2000
	//omly g goroutine runs at any given point
	ch := make(chan string, work)
	g := runtime.NumCPU()
	sem := make(chan struct{}, g)
	for e := 0; e < work; e++ {
		go func(emp int) {
			sem <- struct{}{}
			{
				time.Sleep(time.Duration(rand.Intn(500) * int(time.Millisecond)))
				ch <- "paper"
				log.Printf("employee %d done work %d\n", emp, work)
			}
			<-sem
		}(e)
	}
	for work > 0 {
		p := <-ch
		work--
		log.Printf("recieved %s\n", p)
	}
	close(ch)
}
func FanoutBounded() {
	//only g goroutine run for the entire lifetime of the program
	workPool := []string{"work", "work", "work", 10000: "work"}
	g := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)
	for e := 0; e < g; e++ {
		go func(emp int) {
			defer wg.Done()
			for p := range ch {
				log.Printf("employee %d recieved work %s\n", emp, p)
			}
		}(e)
	}
	for _, work := range workPool {
		ch <- work
	}
	close(ch)
	wg.Wait()
}
