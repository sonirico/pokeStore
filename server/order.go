package main

import "pokeStore/btp"

type Order struct {
	request *btp.Request
	client  *Client
}

func NewOrder(req *btp.Request, client *Client) *Order {
	return &Order{
		request: req,
		client:  client,
	}
}
