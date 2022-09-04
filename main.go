package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	for {
		select {
		case <-time.After(1 * time.Minute):
			err := SayHello()
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

func SayHello() error {
	_, err := fmt.Println("Hello, man!")
	return err
}
