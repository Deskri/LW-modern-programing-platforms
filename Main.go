package main

import (
	"MODERN_PROGRAMMING_PLATFORMS/content"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	application := app.New() //инициализация приложения
	application.Settings().SetTheme(theme.DarkTheme()) //тема приложения - черная
	content.OpenMainWindow(application) //создание окна
}