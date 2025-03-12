package installer

import (
	"errors"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type InputCredPath struct {
	infoView         *tview.TextView
	footerView       *tview.TextView
	parentButtonView *tview.Flex
	listView         *tview.List
	currentPath      string
	selectedPath     string
	err              error
}

func NewInputCredPath() *InputCredPath {
	return &InputCredPath{}
}

func (f *InputCredPath) changeDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	f.listView.Clear()
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		str := file.Name()
		if file.IsDir() {
			str += "/"
		}
		f.listView = f.listView.AddItem(str, "", 0, nil)
	}
	f.currentPath = dir
	f.infoView.SetText("クレデンシャルファイルを選択してください\n現在のフォルダ: " + dir).SetBorder(true)
	return nil
}
func (f *InputCredPath) Run(startDir string) (string, error) {
	app := tview.NewApplication()

	// UI components
	f.infoView = tview.NewTextView()
	f.parentButtonView = tview.NewFlex()
	f.listView = tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFocusOnly(true)
	f.footerView = tview.NewTextView().SetText("'q' を押すとインストールを中止します")

	button := tview.NewButton(".. <親フォルダに移動>")
	button.Box = tview.NewBox()
	button.SetSelectedFunc(func() {
		f.err = f.changeDir(filepath.Dir(f.currentPath))
		if f.err != nil {
			app.Stop()
		}
	})
	f.parentButtonView.AddItem(button, 22, 0, true)

	// build UI
	pages := tview.NewPages()
	body := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(f.listView, 0, 1, true)
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(f.infoView, 4, 0, false).
		AddItem(f.parentButtonView, 1, 0, true).
		AddItem(body, 0, 1, true).
		AddItem(f.footerView, 1, 0, false)
	pages.AddPage("main", mainFlex, true, true)

	// if down pressed on header
	f.parentButtonView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyDown {
			app.SetFocus(f.listView)
			return nil
		}
		return event
	})

	// if up pressed on body
	body.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyUp {
			if f.listView.GetCurrentItem() == 0 {
				app.SetFocus(button)
				return nil
			}
		}
		return event
	})

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
	f.listView.SetSelectedFunc(func(index int, selectedStr string, secondary string, code rune) {
		if strings.HasSuffix(selectedStr, "/") {
			f.err = f.changeDir(filepath.Join(f.currentPath, selectedStr))
		} else {
			f.selectedPath = filepath.Join(f.currentPath, selectedStr)
			app.Stop()
		}
	})

	// -------------------
	// main
	// -------------------
	f.err = f.changeDir(startDir)
	if f.err != nil {
		return "", f.err
	}
	app.SetRoot(pages, true).Run()
	if f.err != nil {
		return "", f.err
	} else {
		return f.selectedPath, nil
	}
}
