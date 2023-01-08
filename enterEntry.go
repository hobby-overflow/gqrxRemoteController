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
	freq := setFreq(e.Text)
	labelPtr.SetText(freq)
	response := SendMessage("F " + freq)
	if response == "RPRT 1\n" {
		labelPtr.SetText("ERROR")
	}
	e.clearEntry()
}

func (e *enterEntry) onA() {
	labelPtr.SetText("Set AM")
	SendMessage("M AM")
}

func (e *enterEntry) onF() {
	labelPtr.SetText("Set FM")
	SendMessage("M FM")
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
