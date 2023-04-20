package content

import (
	"MODERN_PROGRAMMING_PLATFORMS/work"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)

func SetContentLab7(window fyne.Window, application fyne.App) {
	textInfoTask5 := "Задание 5. Разработать программное обеспечение для осуществления тестирования сотрудников предприятия. Программа должна задавать пользователю 5 вопросов на которые пользователь может \nответить тремя вариантами ответа (да, нет, воздержусь). После пройденного тестирования, результаты должны суммироваться в файл. Так же разработайте программу, демонстрирующую количество \n ответов по каждому из вариантов ответов на каждый вопрос."
	
	textTestingApplicationConductingTheTest := "Проведение тестирования сотрудника"
	textTestingResults := "Результаты тестирования"

	labelQuestion1 := widget.NewLabel("Вопрос 1. Вы довольны своей работой?")
	labelQuestion2 := widget.NewLabel("Вопрос 2. Вы высыпаетесь?")
	labelQuestion3 := widget.NewLabel("Вопрос 3. Вы довольны своей зарплатой?")
	labelQuestion4 := widget.NewLabel("Вопрос 4. Вам хватает времени на личный досуг/дела?")
	labelQuestion5 := widget.NewLabel("Вопрос 5. Вы счастливы?")

	labelYes := widget.NewLabel("Да")
	labelNo := widget.NewLabel("Нет")
	labelAbstain := widget.NewLabel("Воздержусь")

	labelQuestion1NumberOfYes := widget.NewLabel("0")
	labelQuestion1NumberOfNo := widget.NewLabel("0")
	labelQuestion1NumberOfAbstain := widget.NewLabel("0")
	labelQuestion2NumberOfYes := widget.NewLabel("0")
	labelQuestion2NumberOfNo := widget.NewLabel("0")
	labelQuestion2NumberOfAbstain := widget.NewLabel("0")
	labelQuestion3NumberOfYes := widget.NewLabel("0")
	labelQuestion3NumberOfNo := widget.NewLabel("0")
	labelQuestion3NumberOfAbstain := widget.NewLabel("0")
	labelQuestion4NumberOfYes := widget.NewLabel("0")
	labelQuestion4NumberOfNo := widget.NewLabel("0")
	labelQuestion4NumberOfAbstain := widget.NewLabel("0")
	labelQuestion5NumberOfYes := widget.NewLabel("0")
	labelQuestion5NumberOfNo := widget.NewLabel("0")
	labelQuestion5NumberOfAbstain := widget.NewLabel("0")

	radioGroupQuestion1 := widget.NewRadioGroup ([]string{"Да", "Нет", "Воздержусь"}, func (s string) {
		
	})
	radioGroupQuestion2 := widget.NewRadioGroup ([]string{"Да", "Нет", "Воздержусь"}, func (s string) {
		
	})
	radioGroupQuestion3 := widget.NewRadioGroup ([]string{"Да", "Нет", "Воздержусь"}, func (s string) {
		
	})
	radioGroupQuestion4 := widget.NewRadioGroup ([]string{"Да", "Нет", "Воздержусь"}, func (s string) {
		
	})
	radioGroupQuestion5 := widget.NewRadioGroup ([]string{"Да", "Нет", "Воздержусь"}, func (s string) {
		
	})

	buttonEndOfTest := widget.NewButton("Отправить", func() {
		work.EmployeeTestingProcessing(
			radioGroupQuestion1, 
			radioGroupQuestion2, 
			radioGroupQuestion3, 
			radioGroupQuestion4, 
			radioGroupQuestion5, 
			labelQuestion1, 
			labelQuestion2, 
			labelQuestion3, 
			labelQuestion4, 
			labelQuestion5, 
			application,
		)
	})

	buttonPatseTestResults := widget.NewButton("Расчитать", func() {
		work.ParseTestResult(
			labelQuestion1NumberOfYes,
			labelQuestion1NumberOfNo,
			labelQuestion1NumberOfAbstain,
			labelQuestion2NumberOfYes,
			labelQuestion2NumberOfNo,
			labelQuestion2NumberOfAbstain,
			labelQuestion3NumberOfYes,
			labelQuestion3NumberOfNo,
			labelQuestion3NumberOfAbstain,
			labelQuestion4NumberOfYes,
			labelQuestion4NumberOfNo,
			labelQuestion4NumberOfAbstain,
			labelQuestion5NumberOfYes,
			labelQuestion5NumberOfNo,
			labelQuestion5NumberOfAbstain,
			labelQuestion1,
			labelQuestion2,
			labelQuestion3,
			labelQuestion4,
			labelQuestion5,
			application,
		)
	})

	accordionItemTestingApplicationConductingTheTest := widget.NewAccordionItem(
		textTestingApplicationConductingTheTest,
		container.NewVBox(
			labelQuestion1, 
			radioGroupQuestion1,
			labelQuestion2, 
			radioGroupQuestion2,
			labelQuestion3, 
			radioGroupQuestion3,
			labelQuestion4, 
			radioGroupQuestion4,
			labelQuestion5, 
			radioGroupQuestion5,
			buttonEndOfTest,
			),
	)

	accordionItemTestingResults := widget.NewAccordionItem(
		textTestingResults,
		container.NewVBox(
			container.NewHBox(
				container.NewVBox(
					labelQuestion1,
					container.NewHBox(
						container.NewVBox(
							labelYes,
							labelQuestion1NumberOfYes,
						),
						container.NewVBox(
							labelNo,
							labelQuestion1NumberOfNo,
						),
						container.NewVBox(
							labelAbstain,
							labelQuestion1NumberOfAbstain,
						),
					),
				),
				container.NewVBox(
					labelQuestion2,
					container.NewHBox(
						container.NewVBox(
							labelYes,
							labelQuestion2NumberOfYes,
						),
						container.NewVBox(
							labelNo,
							labelQuestion2NumberOfNo,
						),
						container.NewVBox(
							labelAbstain,
							labelQuestion2NumberOfAbstain,
						),
					),
				),
				container.NewVBox(
					labelQuestion3,
					container.NewHBox(
						container.NewVBox(
							labelYes,
							labelQuestion3NumberOfYes,
						),
						container.NewVBox(
							labelNo,
							labelQuestion3NumberOfNo,
						),
						container.NewVBox(
							labelAbstain,
							labelQuestion3NumberOfAbstain,
						),
					),
				),
				container.NewVBox(
					labelQuestion4,
					container.NewHBox(
						container.NewVBox(
							labelYes,
							labelQuestion4NumberOfYes,
						),
						container.NewVBox(
							labelNo,
							labelQuestion4NumberOfNo,
						),
						container.NewVBox(
							labelAbstain,
							labelQuestion4NumberOfAbstain,
						),
					),
				),
				container.NewVBox(
					labelQuestion5,
					container.NewHBox(
						container.NewVBox(
							labelYes,
							labelQuestion5NumberOfYes,
						),
						container.NewVBox(
							labelNo,
							labelQuestion5NumberOfNo,
						),
						container.NewVBox(
							labelAbstain,
							labelQuestion5NumberOfAbstain,
						),
					),
				),
			),
			buttonPatseTestResults,
		),
	)


	accordionTestingApplications := widget.NewAccordion(
		accordionItemTestingApplicationConductingTheTest,
		accordionItemTestingResults,
	)

	accordionItemTask5 := widget.NewAccordionItem(
		textInfoTask5,
		accordionTestingApplications,
	)

	accordion := widget.NewAccordion(
		accordionItemTask5,
	)

	window.SetContent(container.NewVScroll(container.NewVBox(
		accordion)))
}