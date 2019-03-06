package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")

	// channel
	// sleep1_finished := make(chan bool)
	// sleep2_finished := make(chan bool)
	// sleep3_finished := make(chan bool)
	sleep_finished := make(chan bool)

	go func() {
		log.Print("sllep1 started ")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 finished")
		sleep_finished <- true
	}()

	go func() {
		log.Print("sleep2 started")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 finished")
		sleep_finished <- true
	}()

	go func() {
		// command taking three minutes
		log.Print("sleep3 started.")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 finished.")
		sleep_finished <- true
	}()

	// need count for concurrent function
	<-sleep_finished
	<-sleep_finished
	<-sleep_finished

	log.Print("all finished")

}
