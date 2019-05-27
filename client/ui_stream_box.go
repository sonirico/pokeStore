package main

import (
	"github.com/rivo/tview"
	"pokeStore/btp"
)

type StreamBox struct {
	*tview.TextView
	responses chan *btp.Response
}

func NewStreamBox(container *tview.Box, responses chan *btp.Response, drawCallBack func()) *StreamBox {
	tv := new(StreamBox)
	tv.TextView = tview.NewTextView().
		SetRegions(false).
		SetDynamicColors(false).
		SetScrollable(true).
		SetWordWrap(true).
		SetWrap(true).
		SetChangedFunc(drawCallBack)
	tv.TextView.Box = container
	tv.responses = responses
	return tv
}

func (sb *StreamBox) WriteResponse(resp *btp.Response) {
	_, err := sb.Write([]byte(resp.String()))
	sb.ScrollToEnd()
	if err != nil {
		panic(err.Error())
	}
}
