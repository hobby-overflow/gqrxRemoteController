package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func setFreq(freq string) string {
	ret := freq
	for i := 0; i <= 8-len(freq); i++ {
		ret += "0"
	}
	return ret
}

func main() {
	a := app.New()
	w := a.NewWindow("GqrxController")

	// input freq
	// show input freq
	// enter to send freq
	// telnet connect
	// send message

	freqView := widget.NewLabel("Hello Fyne")
	input := newEnterEntry()
	input.OnChanged = func(s string) {
		str := strings.ReplaceAll(s, "a", "")
		str = strings.ReplaceAll(str, "f", "")
		input.SetText(str)
	}

	w.Resize(fyne.NewSize(200, 30))
	w.SetContent(container.NewVBox(
		freqView,
		input,
		widget.NewButton("PressEnter", func() {
			input.onEnter()
		}),
	))

	windowPtr = &w
	labelPtr = freqView

	w.Canvas().Focus(input)
	w.ShowAndRun()
}
