package wait

import (
	"log"
	"math/rand"
	"time"
)

func WaitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
	}()
	p := <-ch
	log.Println("recieved data from employee", p)
	// select {}
}

func WaitForTask() {
	ch := make(chan string)
	go func() {
		p := <-ch
		log.Println("employee recieved a signal from manager", p)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}()
	ch <- "paper"
}
