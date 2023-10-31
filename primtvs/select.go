package primtvs

import (
	"fmt"
	"time"
)

func SimpleSelectStatement() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("blocking on read...")

	select {
	case <-c:
		fmt.Printf("unblocked %v later.\n", time.Since(start))
	}
}
