package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type controller struct {
	*widget.Label
	*enterEntry
	*widget.Button
	freq string
}

func newController() *controller {
	c := &controller{}
	c.Label = widget.NewLabel(setFreq("120"))
	c.Button = widget.NewButton("Press Enter", c.onEnter)
	return c
}

type controllerWindow struct {
	fyne.Window
	*controller
}

func newControllerWindow(app fyne.App) *controllerWindow {
	controller := newController()
	w := app.NewWindow("Controller")
	w.Resize(fyne.NewSize(200, 30))
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			controller.Label,
			controller.enterEntry,
			controller.Button,
		),
	)
	return &controllerWindow{w, controller}
}
