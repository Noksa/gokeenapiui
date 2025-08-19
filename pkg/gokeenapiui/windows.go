package gokeenapiui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

var (
	mainWindow fyne.Window
)

func WelcomeWindow() fyne.Window {
	_ = app.NewWithID(uuid.New().String())
	mainWindow = fyne.CurrentApp().NewWindow("Gokeenapi")
	mainWindow.CenterOnScreen()
	exitButton := widget.NewButton("Выход", func() {
		fyne.CurrentApp().Quit()
	})
	exitButton.SetIcon(theme.CancelIcon())
	addAwgButton := widget.NewButton("Создать AWG соединение", func() {
		mainWindow.SetContent(Containers.CreateAwgContainer())
	})
	rawUrl, _ := url.Parse("https://github.com/Noksa/gokeenapi")
	mainWindow.SetContent(container.NewVBox(
		widget.NewLabel(`Gokeenapi - утилита для работы с роутерами Keenetic через REST API

(!) Данная версия является GUI версией и содержит не весь функционал

Если Вам нужен весь функционал, воспользуйтесь CLI версией`), widget.NewHyperlink("CLI версия", rawUrl),
		addAwgButton, exitButton))
	return mainWindow
}
