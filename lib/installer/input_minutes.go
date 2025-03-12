package installer

import (
	"errors"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
)

type InputMinutes struct {
	selected int
	err      error
}

func NewInputMinutes() *InputMinutes {
	return &InputMinutes{}
}

func (f *InputMinutes) Run() (int, error) {
	app := tview.NewApplication()

	// UI components
	infoView := tview.NewTextView()
	infoView.SetBorder(true)
	infoView.SetText("会議の何分前にTV会議準備画面を起動するか指定してください\nおすすめは2分前です")

	listView := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFocusOnly(true)
	for i := 0; i < 6; i++ {
		listView.AddItem(fmt.Sprintf("%d分前", i), fmt.Sprintf("%d", i), 0, nil)
	}
	listView.SetSelectedFunc(func(index int, main string, secondary string, code rune) {
		// string to int
		num, err := strconv.Atoi(secondary)
		if err != nil {
			f.err = err
			app.Stop()
		}
		f.selected = num
		app.Stop()
	})
	listView.SetCurrentItem(2)

	footerView := tview.NewTextView().SetText("'q' を押すとインストールを中止します")

	// build UI
	pages := tview.NewPages()
	body := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(listView, 0, 1, true)
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(infoView, 4, 0, false).
		AddItem(body, 0, 1, true).
		AddItem(footerView, 1, 0, false)
	pages.AddPage("main", mainFlex, true, true)

	// if q pressed
	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case 'q':
				f.err = errors.New("qが押されました")
				app.Stop()
			}
		}
		return event
	})

	// -------------------
	// main
	// -------------------
	app.SetRoot(pages, true).Run()
	if f.err != nil {
		return 0, f.err
	} else {
		return f.selected, nil
	}
}
