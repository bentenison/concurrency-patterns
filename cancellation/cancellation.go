package cancellation

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func CacellationPattern() {
	dur := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200) * int(time.Millisecond)))
		ch <- "work done"
	}()
	select {
	case w := <-ch:
		log.Printf("recieved work from employee %s\n", w)
	case <-ctx.Done():
		log.Println("This is not bearable, employee exceeds deadline")

	}

}
