package content

import (
	"MODERN_PROGRAMMING_PLATFORMS/work"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetContentLab3(window fyne.Window, application fyne.App) {
	textInfoTask1 := "Задание 1. Пользователь вводит два числа a и b, если а > b, то Пользователь вводит третье число с и находит сумму b+c. Если а=b, то программа печатает слово «Конец». Если а<b, то Пользователь вводит \nтретье число с, находит сумму a+b+c, выводит на экран сумму a+b+c и слова «Ну все, допрыгались!»."
	textInfoTask2 := "Задание 2. Пользователь вводит два числа a и b. Если а > b, то Пользователь вводит третье число с, находит сумму a+b+c и выводит ее на экран. Если a<b, то находит разность a-b, если разность четная, \nто вводит третье число c и выводит его на экран. Во всех остальных случаях программа печатает слово «Ерунда какая то»."
	textInfoTask3 := "Задание 3. Пользователь вводит два числа a и b, если а > b, то Пользователь вводит третье число с и считает сумму всех трех чисел. Если а=b, то программа печатает слово «Равны». Если, а<b, то \nПользователь вводит	третье число с, находит разность a-b-c, выводит на экран разность и слова «Задание завершено»."
	textInfoTask4 := "Задание 4. Пользователь вводит два числа a и b, если а > b, то Пользователь вводит третье число с. Находит сумму a+b+c и выводит ее на экран. Если a<b, то находит их разность b-a. Если разность \nнечетная, то ввести третье число и вывести на экран b+c. Во всех остальных случаях напечатать слово «Конец»."
	textInfoTask5 := "Задание 5. Программа просит пользователя ввести a и b, если a>b>5, то программа выводит на экран слово ”осень”, если a=b, то выводит значение выражения (a+b)/c, во всех остальных случаях ввести c \nи вывести значение выражения c+b ."
	textInfoTask6 := "Задание 6. Программа просит пользователя ввести a и b, если a<b, то просит ввести с>0 и находит значение выражения (a+b+c)/3. Во всех остальных случаях выводит на экран «Конец работы»"
	textInfoTask7 := "Задание 7. Пользователь вводит2 числа a и b. Если a<b, то Пользователь вводит третье число c и зеленым цветом выводит значение выражения a+b+c2. Во всех остальных случаях программа выводит на \nэкран красным цветом слова «конец программы»."
	textInfoTask8 := "Задание 8. Пользователь вводит2 числа a и b. Если a=|b|, то программа выводит на экран значение выражения sin2(a2+b2), если a<b, то – выводит на экран слово ’осень’. Если a<0 и b<0, то вводит \nтретье число с и выводит на экран слова «конец программы» и значение выражения a+b-c."

	//ФОРМА ЗАДАНИЕ 1
	// поле ввода
	entryTask1NumA := widget.NewEntry()
	entryTask1NumB := widget.NewEntry()
	entryTask1NumC := widget.NewEntry()

	entryTask1NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask1 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask1 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask1NumA),
		widget.NewFormItem("Введите число B:", entryTask1NumB),
		widget.NewFormItem("Введите число C:", entryTask1NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask1),
	)
	// присваивание кнопки названия "Выполнить"
	formTask1.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask1.OnSubmit = func() {
		work.Lab3Task1(application, entryTask1NumA, entryTask1NumB, entryTask1NumC, labelAnswerTask1)
	}

	//ФОРМА ЗАДАНИЕ 2
	// поле ввода
	entryTask2NumA := widget.NewEntry()
	entryTask2NumB := widget.NewEntry()
	entryTask2NumC := widget.NewEntry()

	entryTask2NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask2 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask2 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask2NumA),
		widget.NewFormItem("Введите число B:", entryTask2NumB),
		widget.NewFormItem("Введите число C:", entryTask2NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask2),
	)
	// присваивание кнопки названия "Выполнить"
	formTask2.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask2.OnSubmit = func() {
		work.Lab3Task2(application, entryTask2NumA, entryTask2NumB, entryTask2NumC, labelAnswerTask2)
	}

	//ФОРМА ЗАДАНИЕ 3
	// поле ввода
	entryTask3NumA := widget.NewEntry()
	entryTask3NumB := widget.NewEntry()
	entryTask3NumC := widget.NewEntry()

	entryTask3NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask3 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask3 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask3NumA),
		widget.NewFormItem("Введите число B:", entryTask3NumB),
		widget.NewFormItem("Введите число C:", entryTask3NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask3),
	)
	// присваивание кнопки названия "Выполнить"
	formTask3.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask3.OnSubmit = func() {
		work.Lab3Task3(application, entryTask3NumA, entryTask3NumB, entryTask3NumC, labelAnswerTask3)
	}

	//ФОРМА ЗАДАНИЕ 4
	// поле ввода
	entryTask4NumA := widget.NewEntry()
	entryTask4NumB := widget.NewEntry()
	entryTask4NumC := widget.NewEntry()

	entryTask4NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask4 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask4 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask4NumA),
		widget.NewFormItem("Введите число B:", entryTask4NumB),
		widget.NewFormItem("Введите число C:", entryTask4NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask4),
	)
	// присваивание кнопки названия "Выполнить"
	formTask4.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask4.OnSubmit = func() {
		work.Lab3Task4(application, entryTask4NumA, entryTask4NumB, entryTask4NumC, labelAnswerTask4)
	}

	//ФОРМА ЗАДАНИЕ 5
	// поле ввода
	entryTask5NumA := widget.NewEntry()
	entryTask5NumB := widget.NewEntry()
	entryTask5NumC := widget.NewEntry()

	entryTask5NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask5 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask5 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask5NumA),
		widget.NewFormItem("Введите число B:", entryTask5NumB),
		widget.NewFormItem("Введите число C:", entryTask5NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask5),
	)
	// присваивание кнопки названия "Выполнить"
	formTask5.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask5.OnSubmit = func() {
		work.Lab3Task5(application, entryTask5NumA, entryTask5NumB, entryTask5NumC, labelAnswerTask5)
	}

	//ФОРМА ЗАДАНИЕ 6
	// поле ввода
	entryTask6NumA := widget.NewEntry()
	entryTask6NumB := widget.NewEntry()
	entryTask6NumC := widget.NewEntry()

	entryTask6NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask6 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask6 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask6NumA),
		widget.NewFormItem("Введите число B:", entryTask6NumB),
		widget.NewFormItem("Введите число C:", entryTask6NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask6),
	)
	// присваивание кнопки названия "Выполнить"
	formTask6.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask6.OnSubmit = func() {
		work.Lab3Task6(application, entryTask6NumA, entryTask6NumB, entryTask6NumC, labelAnswerTask6)
	}

	//ФОРМА ЗАДАНИЕ 7
	// поле ввода
	entryTask7NumA := widget.NewEntry()
	entryTask7NumB := widget.NewEntry()
	entryTask7NumC := widget.NewEntry()

	entryTask7NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask7 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask7 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask7NumA),
		widget.NewFormItem("Введите число B:", entryTask7NumB),
		widget.NewFormItem("Введите число C:", entryTask7NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask7),
	)
	// присваивание кнопки названия "Выполнить"
	formTask7.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask7.OnSubmit = func() {
		work.Lab3Task7(application, entryTask7NumA, entryTask7NumB, entryTask7NumC, labelAnswerTask7)
	}

	//ФОРМА ЗАДАНИЕ 8
	// поле ввода
	entryTask8NumA := widget.NewEntry()
	entryTask8NumB := widget.NewEntry()
	entryTask8NumC := widget.NewEntry()

	entryTask8NumC.Hide()
	// лейбл вывода ответа
	labelAnswerTask8 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask8 := widget.NewForm(
		widget.NewFormItem("Введите число A:", entryTask8NumA),
		widget.NewFormItem("Введите число B:", entryTask8NumB),
		widget.NewFormItem("Введите число C:", entryTask8NumC),
		widget.NewFormItem("Ответ:", labelAnswerTask8),
	)
	// присваивание кнопки названия "Выполнить"
	formTask8.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask8.OnSubmit = func() {
		work.Lab3Task8(application, entryTask8NumA, entryTask8NumB, entryTask8NumC, labelAnswerTask8)
	}

	accordionItemTask1 := widget.NewAccordionItem(
		textInfoTask1,
		formTask1,
	)
	accordionItemTask2 := widget.NewAccordionItem(
		textInfoTask2,
		formTask2,
	)
	accordionItemTask3 := widget.NewAccordionItem(
		textInfoTask3,
		formTask3,
	)
	accordionItemTask4 := widget.NewAccordionItem(
		textInfoTask4,
		formTask4,
	)
	accordionItemTask5 := widget.NewAccordionItem(
		textInfoTask5,
		formTask5,
	)
	accordionItemTask6 := widget.NewAccordionItem(
		textInfoTask6,
		formTask6,
	)
	accordionItemTask7 := widget.NewAccordionItem(
		textInfoTask7,
		formTask7,
	)
	accordionItemTask8 := widget.NewAccordionItem(
		textInfoTask8,
		formTask8,
	)
	accordion := widget.NewAccordion(
		accordionItemTask1,
		accordionItemTask2,
		accordionItemTask3,
		accordionItemTask4,
		accordionItemTask5,
		accordionItemTask6,
		accordionItemTask7,
		accordionItemTask8,
	)

	window.SetContent(container.NewVScroll(container.NewVBox(
		accordion)))
}