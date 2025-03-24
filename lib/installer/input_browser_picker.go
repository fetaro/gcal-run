package installer

import (
	"errors"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
)

type Browser struct {
	Name string
	Path string
}
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

	var browserList []Browser
	if common.IsWindows() {
		browserList = []Browser{
			{"Google Chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"},
			{"Google Chrome", "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"},
			{"Edge", "C:\\Program Files\\Microsoft\\Edge\\Application\\msedge.exe"},
			{"Edge", "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"},
			{"Edge", filepath.Join(os.Getenv("HOME"), "AppData\\Local\\Microsoft\\msedge.exe")},
		}
	} else {
		browserList = []Browser{
			{"Google Chrome", "/Applications/Google Chrome.app"},
			{"Safari", "/Applications/Safari.app"},
		}
	}
	listView := tview.NewList().
		ShowSecondaryText(true).
		SetSelectedFocusOnly(true)
	for _, b := range browserList {
		if common.FileExists(b.Path) {
			listView.AddItem(b.Name, b.Path, 0, nil)
		}
	}
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
