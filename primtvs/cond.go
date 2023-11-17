package primtvs

import (
	"fmt"
	"sync"
	"time"
)

var (
	c     = sync.NewCond(&sync.Mutex{})
	queue = make([]interface{}, 0, 10)
)

func EnqueueWithCond() {
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("addint to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func removeFromQueue(delay time.Duration) {
	time.Sleep(delay)
	c.L.Lock()
	queue = queue[1:]
	fmt.Println("removed from queue")
	c.L.Unlock()
	c.Signal()
}
