package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Aksh-Bansal-dev/starterm/internal/config"
	"github.com/Aksh-Bansal-dev/starterm/internal/tab"
)

const art = `



  _________     ___.                 
 /   _____/____ \_ |__   ___________ 
 \_____  \\__  \ | __ \_/ __ \_  __ \
 /        \/ __ \| \_\ \  ___/|  | \/
/_______  (____  /___  /\___  |__|   
        \/     \/    \/     \/       
`

func main() {
	app := tview.NewApplication()
	view := tview.NewFlex()
	heading := tview.NewTextView().
		SetText(art).SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorBlue)

	list := tview.NewList().
		AddItem("Tmux", "", '1', func() { tab.OpenTab(0) }).
		AddItem("Starterm", "", '2', func() { tab.OpenTab(1) }).
		AddItem("Competitive Programming", "", '3', func() { tab.OpenTab(2) }).
		AddItem("New", "", 'q', func() { app.Stop() })
	config := config.GetConfig()
	for _, item := range config.ConfigItems {
		list.AddItem(item.Name, item.Description, []rune(item.Key)[0], func() { tab.OpenTab(1) })
	}
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			app.Stop()
			return nil
		}
		return event
	})

	view.SetDirection(tview.FlexRow).
		AddItem(heading, 0, 2, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(list, 0, 1, true).
			AddItem(tview.NewBox(), 0, 1, false),
			0, 3, true)

	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}
