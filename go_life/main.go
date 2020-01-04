package main

import "fmt"

import "time"

func main() {
	fmt.Println("main start")
	defer fmt.Println("main quit")

	go func() {
		fmt.Println("go func start")
		defer fmt.Println("go func quit")

		for count := 0; ; count++ {
			fmt.Printf("count: %v\n", count)
			time.Sleep(500 * time.Millisecond)
			if count >= 10 {
				break
			}
		}
	}()

	// try sleep 0 / 2 / 6 second to check life circle
	time.Sleep(2 * time.Second)
}
