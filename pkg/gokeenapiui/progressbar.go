package gokeenapiui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ProgressBar(text string) *fyne.Container {
	progressBar := widget.NewProgressBarInfinite()
	progressBar.Start()
	c := container.NewVBox(widget.NewLabel(text), progressBar)
	return c
}
