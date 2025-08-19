package gokeenapiui

import (
	"fmt"
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/noksa/gokeenapi/pkg/config"
	"github.com/noksa/gokeenapi/pkg/gokeenrestapi"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

var (
	Containers = containers{}
)

const (
	AwgFileConf = "AWG конф. файл"
)

type containers struct {
}

func (*containers) CreateAwgContainer() *fyne.Container {
	creater := func(form *widget.Form) *fyne.Container {
		form.SubmitText = "Создать"
		form.CancelText = "Выход"
		form.OnCancel = func() {
			fyne.CurrentApp().Quit()
		}
		form.OnSubmit = func() {
			Common.SetData()
			dataErr := Common.RouterDataFilled()
			if dataErr != nil {
				msg := ""
				for _, err := range multierr.Errors(dataErr) {
					msg = fmt.Sprintf("%v\n", err.Error())
				}
				dialog.ShowInformation("Не хватает данных", msg, mainWindow)
				return
			}
			v, _ := Bindings.AwgConfFile.Get()
			if v == "" {

			}
			err := gokeenrestapi.Auth()
			if err != nil {
				dialog.ShowInformation("Ошибка авторизации", err.Error(), mainWindow)
				return
			}
			err = gokeenrestapi.Checks.CheckAWGInterfaceExistsFromConfFile(v)
			if err != nil {
				dialog.ShowInformation("Такое соединение уже существует", err.Error(), mainWindow)
				return
			}
			mainWindow.SetContent(ProgressBar("Создаём соединение..."))
			go func() {
				createdInterface, err := gokeenrestapi.AwgConf.AddInterface(v, "")
				if err != nil {
					dialog.ShowInformation("Ошибка создания соединения", err.Error(), mainWindow)
					return
				}
				fyne.Do(func() {
					mainWindow.SetContent(ProgressBar("Настраиваем соединение..."))
				})
				time.Sleep(time.Second * 1)
				err = gokeenrestapi.AwgConf.ConfigureOrUpdateInterface(v, createdInterface.Created)
				if err != nil {
					dialog.ShowInformation("Ошибка настройки соединения", err.Error(), mainWindow)
					return
				}
				err = gokeenrestapi.Interface.SetGlobalIpInInterface(createdInterface.Created, true)
				if err != nil {
					dialog.ShowInformation("Ошибка настройки IP соединения", err.Error(), mainWindow)
					return
				}
				err = gokeenrestapi.Interface.UpInterface(createdInterface.Created)
				if err != nil {
					dialog.ShowInformation("Ошибка включения соединения", err.Error(), mainWindow)
					return
				}
				fyne.Do(func() {
					mainWindow.SetContent(ProgressBar("Ждём пока соединение заработает..."))
				})
				err = gokeenrestapi.Interface.WaitUntilInterfaceIsUp(createdInterface.Created)
				if err != nil {
					dialog.ShowInformation("Соединение с сервером не установлено", err.Error(), mainWindow)
					return
				}
				quitButton := widget.NewButton("Выход", func() {
					fyne.CurrentApp().Quit()
				})
				fyne.Do(func() {
					p, _ := url.Parse(fmt.Sprintf("%v/otherConnections", viper.Get(config.ViperKeeneticUrl)))
					mainWindow.SetContent(container.NewVBox(
						widget.NewLabel(fmt.Sprintf("Соединение успешно создано и включено!\nID созданного соединения: %v\nТеперь можно приступить к настройке политик подключения или маршрутизации и наслаждаться VPN! Удачи! 🌐", createdInterface.Created)),
						widget.NewHyperlink("Открыть веб-интерфейс роутера", p),
						container.NewBorder(nil, nil, nil, quitButton)))
				})
			}()
		}
		c := container.NewVBox(widget.NewLabel(`==> 🛜 <==
Для добавления и настройки нового AWG соединения в keenetic роутере введите данные ниже
Если Ваш keenetic роутер находится в локальной сети (Вы подключены к его Wi-Fi сети),
то к нему можно обратиться по внутреннему IP адреса и http протоколу.
Если роутер доступен через интернет, к нему можно обратиться через KeenDNS имя и https протокол`),
			form)
		return c
	}
	f := Forms.RouterUrlLoginPassword()
	d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			return
		}
		if reader == nil {
			return
		}
		defer func() {
			_ = reader.Close()
		}()
		p := reader.URI().Path()
		if p != "" {
			_ = Bindings.AwgConfFile.Set(p)
			f = Forms.RouterUrlLoginPassword()
			f.AppendItem(&widget.FormItem{
				Text:     AwgFileConf,
				Widget:   widget.NewLabel(p),
				HintText: "Файл выбран",
			})
			mainWindow.SetContent(creater(f))
		}

	}, mainWindow)
	f.AppendItem(&widget.FormItem{
		Text: AwgFileConf,
		Widget: widget.NewButtonWithIcon("Выбрать файл", theme.FolderOpenIcon(), func() {
			d.Show()
		}),
		HintText: "",
	})
	return creater(f)
}
