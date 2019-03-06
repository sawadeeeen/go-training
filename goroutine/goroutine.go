package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started.")

	// Command taking one minutes
	log.Print("sleep1 time.Second.")
	time.Sleep(1 * time.Second)
	log.Print("sleep1 finished.")

	// Command taking two minutes
	log.Print("sleep2 started.")
	time.Sleep(2 * time.Second)
	log.Print("sleep2 finished.")

	// Command taking three minutes
	log.Print("sleep3 started")
	time.Sleep(3 * time.Second)
	log.Print("sleep3 finished.")

	log.Print("all finished.")
}
