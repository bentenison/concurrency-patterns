package barrier

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	count int
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(count int) *Barrier {
	b := &Barrier{
		count: count,
	}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count--
	if b.count == 0 {
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}

func BarrierPattern() {
	items := []int{1, 2, 3, 4, 5}     // Items to process
	barrier := NewBarrier(len(items)) // Create a barrier for the number of items

	var wg sync.WaitGroup
	wg.Add(len(items))

	for _, item := range items {
		go func(i int) {
			defer wg.Done()
			processItem(i)
			barrier.Wait() // Wait for all items to be processed
			fmt.Printf("Item %d processed\n", i)
		}(item)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All items processed. Moving to next stage.")
}

func processItem(item int) {
	time.Sleep(1 * time.Second) // Simulate processing time
	fmt.Printf("Processing item %d\n", item)
}
