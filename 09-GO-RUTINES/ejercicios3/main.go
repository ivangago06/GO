package main

import (
	"fmt"
	"math/rand"
	"time"
)

func productor(c chan<- int) {
	for {
		n := rand.Intn(100)
		c <- n
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func consumidor(c <-chan int) {
	for n := range c {
		fmt.Println(n)
	}
}

func main() {
	c := make(chan int)
	go productor(c)
	consumidor(c)
}
