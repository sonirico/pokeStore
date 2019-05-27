package main

import "github.com/rivo/tview"

type OptionsList struct {
	*tview.List
}

func NewOptionsList() *OptionsList {
	ol := new(OptionsList)
	ol.List = tview.NewList()
	return ol
}

func (ol *OptionsList) AddLink(shortcut rune, title string, callback func()) {
	ol.AddItem(title, "", shortcut, callback)
}
