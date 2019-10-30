package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"pokeStore/btp"
	"strings"
	"time"
)

var requestParser = btp.NewRequestParser()

type Client struct {
	Id     uint64
	socket net.Conn
	store  *Store

	requests  chan *btp.Request
	responses chan *btp.Response
}

func NewClient(id uint64, conn net.Conn, store *Store) *Client {
	return &Client{
		Id:        id,
		socket:    conn,
		store:     store,
		requests:  make(chan *btp.Request),
		responses: make(chan *btp.Response),
	}
}

func (c *Client) Join() {
	c.store.Join <- c
}

func (c *Client) Leave() {
	c.store.Leave <- c
}

func (c *Client) Exit() {
	err := c.socket.Close()
	if err != nil {
		fmt.Println("Error when closing socket", err.Error())
	}
	close(c.requests)
	close(c.responses)
}

func (c *Client) Read() {
	for {
		netData, err := bufio.NewReader(c.socket).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// TODO: Handle client disconnection
			}
			return
		}

		request := requestParser.Parse(netData)
		c.printRequest(request)

		if request.IsValid() {
			c.store.Orders <- NewOrder(request, c)
		} else {
			c.responses <- btp.NewResponse(request.Error.Code, request.Error.Message)
		}
	}
}

func (c *Client) Write() {
	for response := range c.responses {
		_, err := c.socket.Write([]byte(response.String()))
		if err != nil {
			c.printError("server hang out")
			break
		}
	}
}

func (c *Client) printRequest(req *btp.Request) {
	c.print(fmt.Sprintf("[%s][client:%d] - %s\n",
		time.Now(), c.Id, strings.TrimSpace(req.String())))
}

func (c *Client) print(data string) {
	_, _ = io.WriteString(os.Stdout, data)
}

func (c *Client) printError(data string) {
	_, _ = io.WriteString(os.Stderr, data)
}
