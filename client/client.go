package main

import (
	"fmt"
	"github.com/rivo/tview"
	"log"
	"net"
	"pokeStore/btp"
	"pokeStore/cli"
)

type Client struct {
	socket net.Conn

	requests  chan *btp.Request
	responses chan *btp.Response
	exit      chan bool
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		socket:    conn,
		requests:  make(chan *btp.Request),
		responses: make(chan *btp.Response, 10),
	}
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
		payload := make([]byte, 1024)
		bytesRead, err := c.socket.Read(payload)
		if bytesRead < 1 {
			continue
		}
		if err != nil {
			panic(err.Error())
		}
		response, err := btp.ParseResponse(string(payload))
		c.responses <- response
	}
}

func (c *Client) Write(req *btp.Request) {
	if !req.IsValid() {
		fmt.Println("invalid request")
		return
	}
	_, err := c.socket.Write([]byte(req.String()))
	if err != nil {
		fmt.Println("server went down")
	}
}

type UserInterface struct {
	app         *tview.Application
	client      *Client
	optionsList *OptionsList
	pages       *tview.Pages
	streamBox   *StreamBox
}

func NewUserInterface(client *Client) *UserInterface {
	ui := &UserInterface{client: client}

	app := tview.NewApplication()
	pages := tview.NewPages()

	createBasketForm := NewCreateBasketForm(client.Write, ui.SwitchToLanding)
	addItemBasketForm := NewAddItemBasketForm(client.Write, ui.SwitchToLanding)
	checkoutBasketForm := NewCheckoutBasketForm(client.Write, ui.SwitchToLanding)
	dropBasketForm := NewDropBasketForm(client.Write, ui.SwitchToLanding)

	createBasketForm.SetCancelFunc(ui.SwitchToLanding)
	addItemBasketForm.SetCancelFunc(ui.SwitchToLanding)
	checkoutBasketForm.SetCancelFunc(ui.SwitchToLanding)
	dropBasketForm.SetCancelFunc(ui.SwitchToLanding)

	optionsList := NewOptionsList()
	optionsList.SetBorder(true)
	optionsList.AddLink('1', "Create basket", ui.GetSwitchToPage("create-basket"))
	optionsList.AddLink('2', "Add item to basket", ui.GetSwitchToPage("add-item-to-basket"))
	optionsList.AddLink('3', "Checkout basket", ui.GetSwitchToPage("checkout-basket"))
	optionsList.AddLink('4', "Drop checkout", ui.GetSwitchToPage("drop-basket"))
	optionsList.AddLink('0', "Quit", ui.Quit)

	pages.AddPage("landing", optionsList, true, true)
	pages.AddPage("create-basket", createBasketForm, true, false)
	pages.AddPage("add-item-to-basket", addItemBasketForm, true, false)
	pages.AddPage("checkout-basket", checkoutBasketForm, true, false)
	pages.AddPage("drop-basket", dropBasketForm, true, false)

	resultBox := tview.NewBox()
	resultBox.SetBorder(true).SetTitle("Events")
	streamBox := NewStreamBox(resultBox, nil, func() {
		app.Draw()
	})

	flex := tview.NewFlex().
		AddItem(pages, 0, 1, false).
		AddItem(streamBox, 0, 1, true)

	app.SetRoot(flex, true).SetFocus(pages)

	ui.app = app
	ui.pages = pages
	ui.streamBox = streamBox
	ui.optionsList = optionsList

	return ui
}

func (ui *UserInterface) GetSwitchToPage(target string) func() {
	return func() {
		ui.pages.SwitchToPage(target)
	}
}

func (ui *UserInterface) SwitchToLanding() {
	ui.pages.SwitchToPage("landing")
}

func (ui *UserInterface) Quit() {
	ui.app.Stop()
}

func (ui *UserInterface) Run() error {
	go func() {
		for {
			select {
			case resp := <-ui.client.responses:
				ui.streamBox.WriteResponse(resp)
			}
		}
	}()
	go ui.client.Read()
	return ui.app.Run()
}

func main() {
	connData := cli.GetConnectionData()
	conn, err := net.Dial("tcp", connData.Url())
	if err != nil {
		log.Fatal(err)
	}
	client := NewClient(conn)
	defer client.Exit()
	ui := NewUserInterface(client)

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
