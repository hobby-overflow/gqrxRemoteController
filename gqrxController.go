package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ziutek/telnet"
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

	t, err := telnet.Dial("tcp", "127.0.0.1:7356")
	checkErr(err)
	defer t.Close()

	freqView := widget.NewLabel("Hello Fyne")
	input := newEnterEntry()
	input.OnChanged = func(s string) {
		switch s {
		case "a":
			input.onA()
		case "f":
			input.onF()
		}
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
	telnetPtr = t

	w.Canvas().Focus(input)
	w.ShowAndRun()
}
