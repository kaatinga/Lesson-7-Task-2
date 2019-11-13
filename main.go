package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int, 2)
	squares := make(chan int, 2)
	limit := 20
	var currentNumberGen, currentNumberSquare, currentNumberPrint int

	fmt.Println("The limit is set to", limit)

	// генерация
	go func() {
		for x := 1; x <= limit; x++ {
			fmt.Println("Generating a number:", x)
			naturals <- x
			currentNumberGen++
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for {
			if currentNumberSquare != limit {
				x := <-naturals
				fmt.Println("Multiplying a number by itself:", x)
				squares <- x * x
				currentNumberSquare++
				continue
			}
			close(squares)
			break
		}
	}()

	// печать
	for ; currentNumberPrint < limit; {
		fmt.Println("Reading the result: ", <-squares)
		currentNumberPrint++
	}
}
