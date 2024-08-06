package drop

import (
	"fmt"
	"log"
	"sync"
	"time"

	"math/rand"
)

func DropPattern() {
	const cap = 100
	ch := make(chan string, cap)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for p := range ch {
			log.Printf("employee recieved signal %s\n", p)
			time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
		}
	}()
	const work = 500
	for w := 0; w < work; w++ {
		select {
		case ch <- fmt.Sprintf("work--%d", w):
			log.Println("managet sent signal", fmt.Sprintf("work--%d", w))
		default:
			log.Println("manager dropped data", fmt.Sprintf("work--%d", w), "all employees have work already")
		}
	}
	close(ch)
	wg.Wait()
	log.Println("manager sent a shutdown signal to all employees")
}

func DropPatternWithBounded() {
	const cap = 100
	ch := make(chan string, cap)
	employees := 3
	var wg sync.WaitGroup
	wg.Add(employees)

	for i := 0; i < employees; i++ {
		// i = i
		go func(emp int) {
			defer wg.Done()
			for p := range ch {
				log.Printf("employee %d recieved signal %s\n", emp, p)
				time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
			}
		}(i)
	}
	const work = 500
	for w := 0; w < work; w++ {
		select {
		case ch <- fmt.Sprintf("work--%d", w):
			log.Println("managet sent signal", fmt.Sprintf("work--%d", w))
		default:
			log.Println("manager dropped data", fmt.Sprintf("work--%d", w), "all employees have work already")
		}
	}
	close(ch)
	wg.Wait()
	log.Println("manager sent a shutdown signal to all employees")
}
