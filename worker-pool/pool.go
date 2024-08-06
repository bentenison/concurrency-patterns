package workerpool

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"math/rand"
)

func WorkerPool() {
	ch := make(chan string)
	g := runtime.NumCPU()
	log.Printf("number of  employees are %d\n", g)
	// wait for task
	for e := 0; e < g; e++ {
		go func(emp int) {
			for p := range ch {
				log.Printf("employee %d recieved a task %s\n", emp, p)
				time.Sleep(time.Duration(rand.Intn(500) * int(time.Millisecond)))
			}
			log.Printf("employee %d recieved a shutdownn signal\n", emp)
		}(e)
	}
	const work = 100
	// manager sends task
	for w := 0; w < work; w++ {
		ch <- fmt.Sprintf("task %d", w)
	}
	close(ch) // signalling for employee without data to stop processing tasks and shutdown
}
