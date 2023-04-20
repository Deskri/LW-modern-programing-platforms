package content

import (
	"fyne.io/fyne/v2"
)

func OpenMainWindow(application fyne.App) {
	window := application.NewWindow("СПП") // создание окна и присваивание тайтла
	window.Resize(fyne.NewSize(800, 500)) // параметры окна размеры при создании
	window.CenterOnScreen() //размещение окна по центру экрана
	//window.SetFullSize(true) // открытие окна в полном разрешении

	// элемент навигации выходящий на Лабараторную работу №2
	menuItemLab2 := fyne.NewMenuItem("Лабораторная работа №2", func(){
		SetContentLab2(window)
	})
	// элемент навигации выходящий на Лабараторную работу №3
	menuItemLab3 := fyne.NewMenuItem("Лабораторная работа №3", func(){
		SetContentLab3(window, application)
	})
	// элемент навигации выходящий на Лабараторную работу №4
	menuItemLab4 := fyne.NewMenuItem("Лабораторная работа №4", func(){
		SetContentLab4(window)
	})
	// элемент навигации выходящий на Лабараторную работу №5
	menuItemLab5 := fyne.NewMenuItem("Лабораторная работа №5", func(){
		SetContentLab5(window)
	})
	// элемент навигации выходящий на Лабараторную работу №6
	menuItemLab6 := fyne.NewMenuItem("Лабораторная работа №6", func(){
		SetContentLab6(window)
	})
	// элемент навигации выходящий на Лабараторную работу №7
	menuItemLab7 := fyne.NewMenuItem("Лабораторная работа №7", func(){
		SetContentLab7(window, application)
	})
	// группа элементов навигации по лабораторным работам
	menuWork := fyne.NewMenu("Лабораторные работы", menuItemLab2,menuItemLab3, menuItemLab4, menuItemLab5, menuItemLab6, menuItemLab7)
	// навигационный бар
	mainMenu := fyne.NewMainMenu(menuWork)
	//размещение навигационного бара в окне
	window.SetMainMenu(mainMenu)

	//графический вывод окна
	window.ShowAndRun()
}