package appWindows

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
)

func OpenMessageWindow(application fyne.App, textMessage string) {
	window := application.NewWindow("Message") // создание окна и присваивание тайтла
	window.CenterOnScreen() //размещение окна по центру экрана

	labelTextMessage := widget.NewLabel(textMessage)
	buttonCloseMessage := widget.NewButton("OK", func() {
		window.Close()
	})
	window.SetContent(container.NewVBox(labelTextMessage, buttonCloseMessage))
	//графический вывод окна
	window.Show()
}