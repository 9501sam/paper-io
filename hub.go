package main

import "fmt"

type Hub struct {
	str chan string
}

func newHub() *Hub {
	return &Hub{
		str: make(chan string),
	}
}

func (h *Hub) run() {
	fmt.Println("haha")
}
