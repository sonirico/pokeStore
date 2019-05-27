package main

import (
	"github.com/rivo/tview"
	"pokeStore/btp"
	"pokeStore/constants"
)

var itemTypes = []string{
	string(constants.REPELENTE),
	string(constants.CARAMELORARO),
	string(constants.SPORTS),
}

type requestPoster func(*btp.Request)

// CREATE basket

type CreateBasketForm struct {
	*tview.Form

	basketName string

	postRequest requestPoster
}

func (cbf *CreateBasketForm) onSubmit() {
	req := &btp.Request{
		Verb:     btp.Create,
		BasketId: cbf.basketName,
	}

	cbf.postRequest(req)
}

func (cbf *CreateBasketForm) onTextChanged(newText string) {
	cbf.basketName = newText
}

func NewCreateBasketForm(postRequest requestPoster, back func()) *CreateBasketForm {
	newForm := new(CreateBasketForm)
	newForm.basketName = ""
	newForm.postRequest = postRequest
	newForm.Form = tview.NewForm()
	newForm.Form.
		AddInputField("Basket name", "", 20, nil, newForm.onTextChanged).
		SetBorder(true).
		SetTitle("Create a basket").
		SetTitleAlign(tview.AlignLeft)
	newForm.AddButton("Create", newForm.onSubmit)
	newForm.AddButton("Cancel", back)
	return newForm
}

// ADD item to basket

type AddItemBasketForm struct {
	*tview.Form

	basketName  string
	productName string

	postRequest requestPoster
}

func (self *AddItemBasketForm) onProductChanged(option string, optionIndex int) {
	self.productName = option
}

func (self *AddItemBasketForm) onTextChanged(newText string) {
	self.basketName = newText
}

func (self *AddItemBasketForm) onSubmit() {
	req := &btp.Request{
		BasketId: self.basketName,
		ItemType: self.productName,
		Verb:     btp.Add,
	}

	self.postRequest(req)
}

func NewAddItemBasketForm(postRequest func(*btp.Request), back func()) *AddItemBasketForm {
	newForm := new(AddItemBasketForm)
	newForm.basketName = ""
	newForm.postRequest = postRequest
	newForm.Form = tview.NewForm()
	newForm.Form.AddDropDown("Product", itemTypes, 0, newForm.onProductChanged).
		AddInputField("Basket name", "", 20, nil, newForm.onTextChanged).
		SetBorder(true).
		SetTitle("Add item to basket").
		SetTitleAlign(tview.AlignLeft)
	newForm.AddButton("Add item", newForm.onSubmit)
	newForm.AddButton("Cancel", back)
	return newForm
}

// CHECKOUT basket

type CheckoutBasketForm struct {
	*tview.Form

	basketName string

	postRequest requestPoster
}

func (cbf *CheckoutBasketForm) onSubmit() {
	req := &btp.Request{
		Verb:     btp.Checkout,
		BasketId: cbf.basketName,
	}

	cbf.postRequest(req)
}

func (cbf *CheckoutBasketForm) onTextChanged(newText string) {
	cbf.basketName = newText
}

func NewCheckoutBasketForm(postRequest requestPoster, back func()) *CheckoutBasketForm {
	newForm := new(CheckoutBasketForm)
	newForm.basketName = ""
	newForm.postRequest = postRequest
	newForm.Form = tview.NewForm()
	newForm.Form.
		AddInputField("Basket name", "", 20, nil, newForm.onTextChanged).
		SetBorder(true).
		SetTitle("Perform basket checkout").
		SetTitleAlign(tview.AlignLeft)
	newForm.AddButton("Checkout", newForm.onSubmit)
	newForm.AddButton("Cancel", back)
	return newForm
}

// DROP basket

type DropBasketForm struct {
	*tview.Form

	basketName string

	postRequest requestPoster
}

func (cbf *DropBasketForm) onSubmit() {
	req := &btp.Request{
		Verb:     btp.Drop,
		BasketId: cbf.basketName,
	}

	cbf.postRequest(req)
}

func (cbf *DropBasketForm) onTextChanged(newText string) {
	cbf.basketName = newText
}

func NewDropBasketForm(postRequest requestPoster, back func()) *DropBasketForm {
	newForm := new(DropBasketForm)
	newForm.basketName = ""
	newForm.postRequest = postRequest
	newForm.Form = tview.NewForm()
	newForm.Form.
		AddInputField("Basket name", "", 20, nil, newForm.onTextChanged).
		SetBorder(true).
		SetTitle("Drop basket").
		SetTitleAlign(tview.AlignLeft)
	newForm.AddButton("Drop", newForm.onSubmit)
	newForm.AddButton("Cancel", back)
	return newForm
}
