package installer

import (
	"errors"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InputBrowser struct {
	selectedName string
	selectedPath string
	err          error
}

func NewBrowserPicker() *InputBrowser {
	return &InputBrowser{}
}

func (f *InputBrowser) Run() (string, error) {
	app := tview.NewApplication()

	// UI components
	infoView := tview.NewTextView()
	infoView.SetBorder(true)
	infoView.SetText("ブラウザアプリケーションを選択してください")

	listView := tview.NewList().
		ShowSecondaryText(true).
		SetSelectedFocusOnly(true)
	listView.AddItem("Google Chrome", "/Applications/Google Chrome.app", 0, nil)
	listView.AddItem("Safari", "/Applications/Safari.app", 0, nil)

	footerView := tview.NewTextView().SetText("'q' を押すとインストールを中止します")

	// build UI
	pages := tview.NewPages()
	body := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(listView, 0, 1, true)
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(infoView, 3, 0, false).
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

	// if item selected
	listView.SetSelectedFunc(func(index int, main string, secondary string, code rune) {
		f.selectedName = main
		f.selectedPath = secondary
		app.Stop()
	})

	// -------------------
	// main
	// -------------------
	app.SetRoot(pages, true).Run()
	if f.err != nil {
		return "", f.err
	} else {
		return f.selectedPath, nil
	}
}
