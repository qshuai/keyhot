package gui

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var (
	attrstr *ui.AttributedString

	target string
)

type window struct {
	
}

type areaHandler struct {
	window *ui.Window
}

func (ah areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String:      attrstr,
		DefaultFont: &ui.FontDescriptor{},
		Width:       p.AreaWidth,
	})
	defer tl.Free()
	p.Context.Text(tl, 0, 0)
}

func (ah areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// do nothing
}

func (ah areaHandler) MouseCrossed(a *ui.Area, left bool) {
	if left {
		ah.window.Destroy()
		ui.Quit()
	}
	// do nothing
}

func (ah areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (ah areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// reject all keys
	return false
}

func setupUI() {
	mainwin := ui.NewWindow(target, 300, 200, true)
	mainwin.SetMargined(false)
	mainwin.SetBorderless(false)
	mainwin.OnClosing(func(*ui.Window) bool {
		mainwin.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(false)
	mainwin.SetChild(hbox)

	area := ui.NewArea(areaHandler{mainwin})

	hbox.Append(area, true)

	mainwin.Show()
}

func ShowTranslation(word, transaction string)  {
	attrstr = ui.NewAttributedString(transaction)
	target = word

	err := ui.Main(setupUI)
	if err != nil {
		// todo
	}
}
