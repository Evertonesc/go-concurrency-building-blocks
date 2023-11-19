package primtvs

import (
	"fmt"
	"sync"
)

func SimplePoolInitialization() {
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("creating new instance")
			return struct{}{}
		},
	}

	// in this first call there is no instance created in the pool
	pool.Get()

	// here the instance created above is returned
	instance := pool.Get()

	// after finish with the instance, we put it back to the pool
	pool.Put(instance)

	pool.Get()
}

func CalcsCreated() {
	var numCalsCreated int

	calcPool := &sync.Pool{
		New: func() any {
			numCalsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created", numCalsCreated)
}
