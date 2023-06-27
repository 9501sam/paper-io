package main

type Client struct {
	hub *Hub

	send chan []byte
}
