package main

import (
	"fmt"

	"draft/primtvs"
)

func main() {
	resultStream := primtvs.ChanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")

	primtvs.SimpleSelectStatement()

	primtvs.EnqueueWithCond()
}
