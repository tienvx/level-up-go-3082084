package main

import (
	"flag"
	"fmt"
	"log"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	ch := make(chan struct{})
	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Printf("Message '%s' from thread %d\n", message, i)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < n; i++ {
		<-ch
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	if *factor <= 0 {
		log.Fatalf("Factor need to be positive integer")
	}
	for _, m := range messages {
		log.Println(m)
		repeat(int(*factor), m)
	}
}
