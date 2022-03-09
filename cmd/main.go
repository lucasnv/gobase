package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("Start app")

	for {
		fmt.Println("Is where my deamons hide")
		time.Sleep(time.Second * 15)
	}
}
