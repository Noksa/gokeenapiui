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
			connectionName, _ := Bindings.AwgName.Get()
			err := gokeenrestapi.Auth()
			if err != nil {
				dialog.ShowInformation("Ошибка авторизации", err.Error(), mainWindow)
				return
			}
			err = gokeenrestapi.Checks.CheckAWGInterfaceExistsFromConfFile(v)
			if err != nil {
				dialog.ShowInformation("Ошибка чтения conf файла", err.Error(), mainWindow)
				return
			}
			mainWindow.SetContent(ProgressBar("Создаём соединение..."))
			go func() {
				createdInterface, err := gokeenrestapi.AwgConf.AddInterface(v, connectionName)
				if err != nil {
					dialog.ShowInformation("Ошибка создания соединения", err.Error(), mainWindow)
					errorProcess(err)
					return
				}
				fyne.Do(func() {
					mainWindow.SetContent(ProgressBar("Настраиваем соединение..."))
				})
				time.Sleep(time.Second * 1)
				err = gokeenrestapi.AwgConf.ConfigureOrUpdateInterface(v, createdInterface.Created)
				if err != nil {
					dialog.ShowInformation("Ошибка настройки соединения", err.Error(), mainWindow)
					errorProcess(err)
					return
				}
				err = gokeenrestapi.Interface.SetGlobalIpInInterface(createdInterface.Created, true)
				if err != nil {
					dialog.ShowInformation("Ошибка настройки IP соединения", err.Error(), mainWindow)
					errorProcess(err)
					return
				}
				err = gokeenrestapi.Interface.UpInterface(createdInterface.Created)
				if err != nil {
					dialog.ShowInformation("Ошибка включения соединения", err.Error(), mainWindow)
					errorProcess(err)
					return
				}
				fyne.Do(func() {
					mainWindow.SetContent(ProgressBar("Ждём пока соединение заработает..."))
				})
				err = gokeenrestapi.Interface.WaitUntilInterfaceIsUp(createdInterface.Created)
				if err != nil {
					dialog.ShowInformation("Соединение с сервером не установлено", err.Error(), mainWindow)
					errorProcess(err)
					return
				}
				quitButton := widget.NewButton("Выход", func() {
					fyne.CurrentApp().Quit()
				})
				fyne.Do(func() {
					p, _ := url.Parse(fmt.Sprintf("%v/otherConnections", viper.Get(config.ViperKeeneticUrl)))
					mainWindow.SetContent(container.NewVBox(
						widget.NewLabel(fmt.Sprintf("Соединение успешно создано и включено!\nID созданного соединения: %v\nТеперь можно приступить к настройке политик подключения или маршрутизации и наслаждаться VPN!\nУдачи! 🌐", createdInterface.Created)),
						widget.NewHyperlink("Открыть веб-интерфейс роутера", p),
						container.NewBorder(nil, nil, quitButton, nil, quitButton)))
				})
			}()
		}
		c := container.NewVBox(widget.NewLabel(`==> 🛜
Для добавления и настройки нового AWG соединения в keenetic роутере введите данные ниже
Если Ваш keenetic роутер находится в локальной сети (Вы подключены к его Wi-Fi сети),
то к нему можно обратиться по внутреннему IP адресу и HTTP протоколу.
Если роутер доступен через интернет, к нему можно обратиться через KeenDNS имя и HTTPS протокол
Примеры:
Внутренний адрес: http://192.168.1.1
Внешний адрес: https://super-keenetic.keenetic.pro`),
			form)
		return c
	}
	f := Forms.RouterUrlLoginPassword()
	var d *dialog.FileDialog
	repickFileButton := widget.NewButtonWithIcon("Выбрать другой файл", theme.UploadIcon(), func() {
		d.Show()
	})
	d = dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
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
			addAwgNameField(f)
			f.AppendItem(&widget.FormItem{
				Text:     AwgFileConf,
				Widget:   widget.NewLabel(p),
				HintText: "Файл выбран",
			})
			f.Append("", repickFileButton)
			mainWindow.SetContent(creater(f))
		}

	}, mainWindow)
	d.SetTitleText("Выбор AWG конфиг файла")
	d.SetDismissText("Отмена")
	d.SetConfirmText("Выбрать")
	d.Resize(fyne.NewSize(800, 600))
	addAwgNameField(f)
	f.AppendItem(&widget.FormItem{
		Text: AwgFileConf,
		Widget: widget.NewButtonWithIcon("Выбрать файл", theme.FolderOpenIcon(), func() {
			d.Show()
		}),
		HintText: "",
	})
	return creater(f)
}

func addAwgNameField(f *widget.Form) {
	f.AppendItem(&widget.FormItem{
		Text:     "Имя соед.",
		Widget:   widget.NewEntryWithData(Bindings.AwgName),
		HintText: "Имя создаваемого соединения. Опционально, по умолчанию равно имени файла",
	})
}

func errorProcess(err error) {
	if err == nil {
		return
	}
	fyne.Do(func() {
		mainWindow.SetContent(container.NewVBox(widget.NewLabel(fmt.Sprintf("Произошла ошибка! Попробуйте снова :(\n\nОшибка: %v", err.Error()))))
	})
}
