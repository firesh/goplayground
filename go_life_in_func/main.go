package main

import "fmt"

import "time"

func main() {
	fmt.Println("main start")
	defer fmt.Println("main quit")

	func1()

	// try sleep 0 / 2 / 6 second to check life circle
	time.Sleep(6 * time.Second)
}

func func1() {
	fmt.Println("go func1 start")
	defer fmt.Println("go func1 quit")

	// func2 will continue when func1 quit
	go func2()
	time.Sleep(2 * time.Second)
}

func func2() {
	fmt.Println("go func2 start")
	defer fmt.Println("go func2 quit")

	for count := 0; ; count++ {
		fmt.Printf("count: %v\n", count)
		time.Sleep(500 * time.Millisecond)
		if count >= 10 {
			break
		}
	}
}