package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Aksh-Bansal-dev/starterm/internal/config"
	"github.com/Aksh-Bansal-dev/starterm/internal/runner"
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
	config := config.GetConfig()
	app := tview.NewApplication()
	view := tview.NewFlex()

	headingText := config.Heading
	if headingText == "" {
		headingText = art
	}

	heading := tview.NewTextView().
		SetText(headingText).SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorBlue)

	list := tview.NewList().
		AddItem("New", "", 'q', func() { app.Stop() })
	for _, item := range config.ConfigItems {
		foo := item.Cmd
		list.AddItem(item.Name, item.Description, []rune(item.Key)[0], func() { runner.Run(foo) })
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
