package main

import (
	"flag"
	"fmt"
)

func main() {
	nworkers := flag.Int("workers", 1, "number of worker threads")
	flag.Parse()

	msg := fmt.Sprintf("Starting the Bauss + %d worker(s)", *nworkers)
	fmt.Println(msg)

	for i := 0; i < *nworkers; i++ {
		go worker()
	}
}

func worker() {
	for {
	}
}
