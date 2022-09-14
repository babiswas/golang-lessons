package main

import "fmt"

func main() {
	ch1 := make(chan string, 2)
	ch1 <- "hello"
	ch1 <- "bello"
	close(ch1)
	for val := range ch1 {
		fmt.Println(val)
	}
}
