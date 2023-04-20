package content

import (
	"MODERN_PROGRAMMING_PLATFORMS/work"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetContentLab4(window fyne.Window) {
	//textInfoTask1 := "Задание 1. Переставить два слова. С клавиатуры вводится строка, состоящая ровно из двух слов, разделенных пробелом. Переставьте эти слова местами. Результат запишите в строку и выведите получившуюся строку. При решении этой задачи не стоит пользоваться циклами и инструкцией if."
	//textInfoTask2 := "Задание 2. Удаление фрагмента. С клавиатуры вводится строка, в которой буква h встречается минимум два раза. Удалите из этой строки первое и последнее вхождение буквы h, а также все символы, находящиеся между ними."
	//textInfoTask3 := "Задание 3. Замена подстроки. Считать строку и один символ (ключ). Замените в исходной строке все символы равные значению ключа на слово NULL."
	//textInfoTask4 := "Задание 4. Удалить каждый третий символ. Считать с клавиатуры строку и число. Удалите из нее все символы, чьи индексы делятся на введенное число."
	textInfoTask5 := "Задание 5. Составить программу подсчета общего количества цифр и знаков +, -, * в строке введенной с клавиатуры."
	//textInfoTask6 := "Задание 6. Считать строку с клавиатуры. Удалить все числа из введенной строки и вывести оставшийся текст на экран."
	//textInfoTask7 := "Задание 7. Считать строку русских символов с клавиатуры. Заменить русские символы аналогичными латинскими."
	//textInfoTask8 := "Задание 8. Заполнить строку случайными символами русского и латинского алфавита. Посчитать количество символов каждого из алфавитов."

	//ФОРМА ЗАДАНИЕ 5
	// поле ввода
	entryTask5 := widget.NewEntry()
	// лейбл вывода ответа
	labelAnswerTask5 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formTask5 := widget.NewForm(
		widget.NewFormItem("Введите текст:", entryTask5),
		widget.NewFormItem("Ответ:", labelAnswerTask5),
	)
	// присваивание кнопки названия "Выполнить"
	formTask5.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formTask5.OnSubmit = func() {
		labelAnswerTask5.SetText(work.Lab4Task5(entryTask5.Text))
	}

	//accordionItemTask1 := widget.NewAccordionItem(
	//	textInfoTask1,
	//	formTask1,
	//)
	//accordionItemTask2 := widget.NewAccordionItem(
	//	textInfoTask2,
	//	formTask2,
	//)
	//accordionItemTask3 := widget.NewAccordionItem(
	//	textInfoTask3,
	//	formTask3,
	//)
	//accordionItemTask4 := widget.NewAccordionItem(
	//	textInfoTask4,
	//	formTask4,
	//)
	accordionItemTask5 := widget.NewAccordionItem(
		textInfoTask5,
		formTask5,
	)
	//accordionItemTask6 := widget.NewAccordionItem(
	//	textInfoTask6,
	//	formTask6,
	//)
	//accordionItemTask7 := widget.NewAccordionItem(
	//	textInfoTask7,
	//	formTask7,
	//)
	//accordionItemTask8 := widget.NewAccordionItem(
	//	textInfoTask8,
	//	formTask8,
	//)
	accordion := widget.NewAccordion(
	//	accordionItemTask1,
	//	accordionItemTask2,
	//	accordionItemTask3,
	//	accordionItemTask4,
		accordionItemTask5,
	//	accordionItemTask6,
	//	accordionItemTask7,
	//	accordionItemTask8,
	)

	window.SetContent(container.NewVScroll(container.NewVBox(
		accordion)))
}