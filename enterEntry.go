package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var windowPtr *fyne.Window
var labelPtr *widget.Label

type enterEntry struct {
	widget.Entry
}

func (e *enterEntry) onEnter() {
	if e.Text == "" {
		return
	}
	labelPtr.SetText(setFreq(e.Text))
	e.clearEntry()
}

func (e *enterEntry) onA() {
	println("Set AM")
}

func (e *enterEntry) onF() {
	println("Set FM")
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEnter:
		e.onEnter()
	case fyne.KeyReturn:
		e.onEnter()
	case fyne.KeyA:
		e.onA()
	case fyne.KeyF:
		e.onF()
	case fyne.KeyEscape:
		(*windowPtr).Close()
	}
}

func (e *enterEntry) clearEntry() {
	e.SetText("")
}
