package main

import "fmt"

func main() {
	resultStream := ChanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
