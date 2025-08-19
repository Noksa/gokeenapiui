package gokeenapiui

import "fyne.io/fyne/v2/widget"

var (
	Forms = forms{}
)

type forms struct {
}

func (*forms) RouterUrlLoginPassword() *widget.Form {
	urlWidget := widget.NewEntryWithData(Bindings.RouterUrl)
	if urlWidget.Text == "" {
		urlWidget.SetText("http://192.168.1.1")
	}
	url := &widget.FormItem{
		Text:     "URL роутера",
		Widget:   urlWidget,
		HintText: "IP или DNS имя (протокол http/https) роутера",
	}
	loginWidget := widget.NewEntryWithData(Bindings.RouterLogin)
	if loginWidget.Text == "" {
		loginWidget.SetText("admin")
	}
	passwordWidget := widget.NewPasswordEntry()
	passwordWidget.Bind(Bindings.RouterPassword)
	f := widget.NewForm(url,
		&widget.FormItem{
			Text:     "Логин",
			Widget:   loginWidget,
			HintText: "Логин администратора",
		},
		&widget.FormItem{
			Text:     "Пароль",
			Widget:   passwordWidget,
			HintText: "Пароль администратора",
		})
	return f
}
