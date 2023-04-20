package work

import (
	"fmt"
	"MODERN_PROGRAMMING_PLATFORMS/content/appWindows"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"math"
)

func Lab3Task1(application fyne.App, entryTask1NumA, entryTask1NumB, entryTask1NumC *widget.Entry, labelAnswerTask1 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.Atoi(entryTask1NumA.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask1.SetText("Введите корректное значение в число А")
		entryTask1NumA.Show()
		entryTask1NumB.Hide()
		entryTask1NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.Atoi(entryTask1NumB.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask1.SetText("Введите корректное значение в число B")
		entryTask1NumA.Hide()
		entryTask1NumB.Show()
		entryTask1NumC.Hide()
		return
	}

	if a > b && entryTask1NumC.Text == ""  { // в случае, если число "a" больше "b" и число "c" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask1NumA.Hide()
		entryTask1NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask1NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask1.SetText("Введите число С")
	}else if a > b && entryTask1NumC.Text != "" { // в случае, если число "a" больше "b" и число "c" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask1NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask1.SetText("Введите корректное значение в число C")
			entryTask1NumA.Hide()
			entryTask1NumB.Hide()
			entryTask1NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "(a+b)/c" из числового значения в строку
		labelAnswerTask1.SetText(strconv.Itoa(b + c))
		// возвращает изноальную видимость полей ввода
		entryTask1NumA.Show()
		entryTask1NumB.Show()
		entryTask1NumC.Hide()
		// убирает текст из полей ввода
		entryTask1NumA.SetText("")
		entryTask1NumB.SetText("")
		entryTask1NumC.SetText("")
	}else if a == b{ // в случае, если число "a" равно "b" 
		appWindows.OpenMessageWindow(application, "Конец")
		// убирает текст из полей ввода
		entryTask1NumA.SetText("")
		entryTask1NumB.SetText("")
		entryTask1NumC.SetText("")
	}else if a < b && entryTask1NumC.Text == "" { // условие все оставшиеся варианты и число "с" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask1NumA.Hide()
		entryTask1NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask1NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask1.SetText("Введите число С")
	}else if a < b && entryTask1NumC.Text != "" { // условие все оставшиеся варианты и число "с" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask1NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask1.SetText("Введите корректное значение в число C")
			entryTask1NumA.Hide()
			entryTask1NumB.Hide()
			entryTask1NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "c + b" из числового значения в строку
		labelAnswerTask1.SetText(strconv.Itoa(a + b + c))
		// возвращает изноальную видимость полей ввода
		entryTask1NumA.Show()
		entryTask1NumB.Show()
		entryTask1NumC.Hide()
		// убирает текст из полей ввода
		entryTask1NumA.SetText("")
		entryTask1NumB.SetText("")
		entryTask1NumC.SetText("")
		// вывод окна с текстом "Ну все, допрыгались!"
		appWindows.OpenMessageWindow(application, "Ну все, допрыгались!")
	}
}

func Lab3Task2(application fyne.App, entryTask2NumA, entryTask2NumB, entryTask2NumC *widget.Entry, labelAnswerTask2 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.Atoi(entryTask2NumA.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask2.SetText("Введите корректное значение в число А")
		entryTask2NumA.Show()
		entryTask2NumB.Hide()
		entryTask2NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.Atoi(entryTask2NumB.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask2.SetText("Введите корректное значение в число B")
		entryTask2NumA.Hide()
		entryTask2NumB.Show()
		entryTask2NumC.Hide()
		return
	}

	if a > b && entryTask2NumC.Text == ""  { // в случае, если число "a" больше "b" и число "c" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask2NumA.Hide()
		entryTask2NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask2NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask2.SetText("Введите число С")
	}else if a > b && entryTask2NumC.Text != "" { // в случае, если число "a" больше "b" и число "c" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask2NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask2.SetText("Введите корректное значение в число C")
			entryTask2NumA.Hide()
			entryTask2NumB.Hide()
			entryTask2NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "(a+b)/c" из числового значения в строку
		labelAnswerTask2.SetText(strconv.Itoa(a + b + c))
		// возвращает изноальную видимость полей ввода
		entryTask2NumA.Show()
		entryTask2NumB.Show()
		entryTask2NumC.Hide()
		// убирает текст из полей ввода
		entryTask2NumA.SetText("")
		entryTask2NumB.SetText("")
		entryTask2NumC.SetText("")
	}else if a == b{ // в случае, если число "a" равно "b" 
		appWindows.OpenMessageWindow(application, "Ерунда какая то")
		// убирает текст из полей ввода
		entryTask2NumA.SetText("")
		entryTask2NumB.SetText("")
		entryTask2NumC.SetText("")
	}else if a < b { // условие все оставшиеся варианты и число "с" пусто
		// формирование ответа в виде строки, конвератция выражениея "c + b" из числового значения в строку
		labelAnswerTask2.SetText(strconv.Itoa(a - b))
		// убирает текст из полей ввода
		entryTask2NumA.SetText("")
		entryTask2NumB.SetText("")
		entryTask2NumC.SetText("")
	}
}

func Lab3Task3(application fyne.App, entryTask3NumA, entryTask3NumB, entryTask3NumC *widget.Entry, labelAnswerTask3 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.Atoi(entryTask3NumA.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask3.SetText("Введите корректное значение в число А")
		entryTask3NumA.Show()
		entryTask3NumB.Hide()
		entryTask3NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.Atoi(entryTask3NumB.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask3.SetText("Введите корректное значение в число B")
		entryTask3NumA.Hide()
		entryTask3NumB.Show()
		entryTask3NumC.Hide()
		return
	}

	if a > b && entryTask3NumC.Text == ""  { // в случае, если число "a" больше "b" и число "c" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask3NumA.Hide()
		entryTask3NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask3NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask3.SetText("Введите число С")
	}else if a > b && entryTask3NumC.Text != "" { // в случае, если число "a" больше "b" и число "c" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask3NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask3.SetText("Введите корректное значение в число C")
			entryTask3NumA.Hide()
			entryTask3NumB.Hide()
			entryTask3NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "a+b+c" из числового значения в строку
		labelAnswerTask3.SetText(strconv.Itoa(a + b + c))
		// возвращает изноальную видимость полей ввода
		entryTask3NumA.Show()
		entryTask3NumB.Show()
		entryTask3NumC.Hide()
		// убирает текст из полей ввода
		entryTask3NumA.SetText("")
		entryTask3NumB.SetText("")
		entryTask3NumC.SetText("")
	}else if a == b{ // в случае, если число "a" равно "b" 
		appWindows.OpenMessageWindow(application, "Равны")
		// убирает текст из полей ввода
		entryTask3NumA.SetText("")
		entryTask3NumB.SetText("")
		entryTask3NumC.SetText("")
		// отчистка ответа 
		labelAnswerTask3.SetText("")
	}else if a < b && entryTask3NumC.Text == "" { // условие все оставшиеся варианты и число "с" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask3NumA.Hide()
		entryTask3NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask3NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask3.SetText("Введите число С")
	}else if a < b && entryTask3NumC.Text != "" { // условие все оставшиеся варианты и число "с" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask3NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask3.SetText("Введите корректное значение в число C")
			entryTask3NumA.Hide()
			entryTask3NumB.Hide()
			entryTask3NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "a-b-c" из числового значения в строку
		labelAnswerTask3.SetText(strconv.Itoa(a - b - c))
		// возвращает изноальную видимость полей ввода
		entryTask3NumA.Show()
		entryTask3NumB.Show()
		entryTask3NumC.Hide()
		// убирает текст из полей ввода
		entryTask3NumA.SetText("")
		entryTask3NumB.SetText("")
		entryTask3NumC.SetText("")
		// вывод окна с текстом "Ну все, допрыгались!"
		appWindows.OpenMessageWindow(application, "Задание завершено")
	}
}

func Lab3Task4(application fyne.App, entryTask4NumA, entryTask4NumB, entryTask4NumC *widget.Entry, labelAnswerTask4 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.Atoi(entryTask4NumA.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask4.SetText("Введите корректное значение в число А")
		entryTask4NumA.Show()
		entryTask4NumB.Hide()
		entryTask4NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.Atoi(entryTask4NumB.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask4.SetText("Введите корректное значение в число B")
		entryTask4NumA.Hide()
		entryTask4NumB.Show()
		entryTask4NumC.Hide()
		return
	}

	if a > b && entryTask4NumC.Text == ""  { // в случае, если число "a" больше "b" и число "c" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask4NumA.Hide()
		entryTask4NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask4NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask4.SetText("Введите число С")
	}else if a > b && entryTask4NumC.Text != "" { // в случае, если число "a" больше "b" и число "c" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask4NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask4.SetText("Введите корректное значение в число C")
			entryTask4NumA.Hide()
			entryTask4NumB.Hide()
			entryTask4NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "a+b+c" из числового значения в строку
		labelAnswerTask4.SetText(strconv.Itoa(a + b + c))
		// возвращает изноальную видимость полей ввода
		entryTask4NumA.Show()
		entryTask4NumB.Show()
		entryTask4NumC.Hide()
		// убирает текст из полей ввода
		entryTask4NumA.SetText("")
		entryTask4NumB.SetText("")
		entryTask4NumC.SetText("")
	}else if a < b && (b - a) % 2 != 0 && entryTask4NumC.Text == "" {
		// скрываются поля ввода для число "a" и "b"
		entryTask4NumA.Hide()
		entryTask4NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask4NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask4.SetText("Введите число С")
	}else if a < b && (b - a) % 2 != 0 && entryTask4NumC.Text != "" {
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask4NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask4.SetText("Введите корректное значение в число C")
			entryTask4NumA.Hide()
			entryTask4NumB.Hide()
			entryTask4NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "a+b+c" из числового значения в строку
		labelAnswerTask4.SetText(strconv.Itoa(b + c))
		// возвращает изноальную видимость полей ввода
		entryTask4NumA.Show()
		entryTask4NumB.Show()
		entryTask4NumC.Hide()
		// убирает текст из полей ввода
		entryTask4NumA.SetText("")
		entryTask4NumB.SetText("")
		entryTask4NumC.SetText("")
	}else {
		appWindows.OpenMessageWindow(application, "Конец")
		// убирает текст из полей ввода
		entryTask4NumA.SetText("")
		entryTask4NumB.SetText("")
		entryTask4NumC.SetText("")
		labelAnswerTask4.SetText("")
	}
}

func Lab3Task5(application fyne.App, entryTask5NumA, entryTask5NumB, entryTask5NumC *widget.Entry, labelAnswerTask5 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.Atoi(entryTask5NumA.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask5.SetText("Введите корректное значение в число А")
		entryTask5NumA.Show()
		entryTask5NumB.Hide()
		entryTask5NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.Atoi(entryTask5NumB.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask5.SetText("Введите корректное значение в число B")
		entryTask5NumA.Hide()
		entryTask5NumB.Show()
		entryTask5NumC.Hide()
		return
	}

	if a > b && b > 5 { // в случае, если число "a" больше "b" и число "b" больше 5, то выводится окно с сообщением "осень"
		appWindows.OpenMessageWindow(application, "осень")
		// убирает текст из полей ввода
		entryTask5NumA.SetText("")
		entryTask5NumB.SetText("")
		entryTask5NumC.SetText("")
	}else if a == b && entryTask5NumC.Text == "" { // в случае, если число "a" равно "b" и число "c" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask5NumA.Hide()
		entryTask5NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask5NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask5.SetText("Введите число С")
	}else if a == b && entryTask5NumC.Text != "" { // в случае, если число "a" равно "b" и число "c" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask5NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask5.SetText("Введите корректное значение в число C")
			entryTask5NumA.Hide()
			entryTask5NumB.Hide()
			entryTask5NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "(a+b)/c" из числового значения в строку
		labelAnswerTask5.SetText(strconv.Itoa((a + b) / c))
		// возвращает изноальную видимость полей ввода
		entryTask5NumA.Show()
		entryTask5NumB.Show()
		entryTask5NumC.Hide()
		// убирает текст из полей ввода
		entryTask5NumA.SetText("")
		entryTask5NumB.SetText("")
		entryTask5NumC.SetText("")
	}else if entryTask5NumC.Text == "" { // условие все оставшиеся варианты и число "с" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask5NumA.Hide()
		entryTask5NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask5NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask5.SetText("Введите число С")
	}else if entryTask5NumC.Text != "" { // условие все оставшиеся варианты и число "с" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask5NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask5.SetText("Введите корректное значение в число C")
			entryTask5NumA.Hide()
			entryTask5NumB.Hide()
			entryTask5NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "c + b" из числового значения в строку
		labelAnswerTask5.SetText(strconv.Itoa(c + b))
		// возвращает изноальную видимость полей ввода
		entryTask5NumA.Show()
		entryTask5NumB.Show()
		entryTask5NumC.Hide()
		// убирает текст из полей ввода
		entryTask5NumA.SetText("")
		entryTask5NumB.SetText("")
		entryTask5NumC.SetText("")
	}
}

func Lab3Task6(application fyne.App, entryTask6NumA, entryTask6NumB, entryTask6NumC *widget.Entry, labelAnswerTask6 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.Atoi(entryTask6NumA.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask6.SetText("Введите корректное значение в число А")
		entryTask6NumA.Show()
		entryTask6NumB.Hide()
		entryTask6NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.Atoi(entryTask6NumB.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask6.SetText("Введите корректное значение в число B")
		entryTask6NumA.Hide()
		entryTask6NumB.Show()
		entryTask6NumC.Hide()
		return
	}

	if a < b && entryTask6NumC.Text == ""  { // в случае, если число "a" меньше "b" и число "c" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask6NumA.Hide()
		entryTask6NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask6NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask6.SetText("Введите число С")
	}else if a < b && entryTask6NumC.Text != "" { // в случае, если число "a" меншье "b" и число "c" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask6NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask6.SetText("Введите корректное значение в число C")
			entryTask6NumA.Hide()
			entryTask6NumB.Hide()
			entryTask6NumC.Show()
			return
		}
		// проверка, является ли с > 0, если нет, просит внести больше 0
		if c <= 0 {
			labelAnswerTask6.SetText("Введите значение в число C, которое больше 0")
			entryTask6NumA.Hide()
			entryTask6NumB.Hide()
			entryTask6NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "(a+b+c)/3" из числового значения в строку
		labelAnswerTask6.SetText(strconv.Itoa((a + b + c) / 3))
		// возвращает изноальную видимость полей ввода
		entryTask6NumA.Show()
		entryTask6NumB.Show()
		entryTask6NumC.Hide()
		// убирает текст из полей ввода
		entryTask6NumA.SetText("")
		entryTask6NumB.SetText("")
		entryTask6NumC.SetText("")
	}else { // во всех остальных случаях
		appWindows.OpenMessageWindow(application, "Конец работы")
		// убирает текст из полей ввода
		entryTask6NumA.SetText("")
		entryTask6NumB.SetText("")
		entryTask6NumC.SetText("")
		// отчистка ответа 
		labelAnswerTask6.SetText("")
	}
}

func Lab3Task7(application fyne.App, entryTask7NumA, entryTask7NumB, entryTask7NumC *widget.Entry, labelAnswerTask7 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.Atoi(entryTask7NumA.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask7.SetText("Введите корректное значение в число А")
		entryTask7NumA.Show()
		entryTask7NumB.Hide()
		entryTask7NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.Atoi(entryTask7NumB.Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask7.SetText("Введите корректное значение в число B")
		entryTask7NumA.Hide()
		entryTask7NumB.Show()
		entryTask7NumC.Hide()
		return
	}

	if a < b && entryTask7NumC.Text == ""  { // в случае, если число "a" меньше "b" и число "c" пусто
		// скрываются поля ввода для число "a" и "b"
		entryTask7NumA.Hide()
		entryTask7NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask7NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask7.SetText("Введите число С")
	}else if a < b && entryTask7NumC.Text != "" { // в случае, если число "a" меншье "b" и число "c" содержит какие-то символы
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.Atoi(entryTask7NumC.Text)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask7.SetText("Введите корректное значение в число C")
			entryTask7NumA.Hide()
			entryTask7NumB.Hide()
			entryTask7NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "(a+b+c)/3" из числового значения в строку
		labelAnswerTask7.SetText(strconv.Itoa(a + b + c * 2))
		// возвращает изноальную видимость полей ввода
		entryTask7NumA.Show()
		entryTask7NumB.Show()
		entryTask7NumC.Hide()
		// убирает текст из полей ввода
		entryTask7NumA.SetText("")
		entryTask7NumB.SetText("")
		entryTask7NumC.SetText("")
	}else { // во всех остальных случаях
		appWindows.OpenMessageWindow(application, "конец программы")
		// убирает текст из полей ввода
		entryTask7NumA.SetText("")
		entryTask7NumB.SetText("")
		entryTask7NumC.SetText("")
		// отчистка ответа 
		labelAnswerTask7.SetText("")
	}
}

func Lab3Task8(application fyne.App, entryTask8NumA, entryTask8NumB, entryTask8NumC *widget.Entry, labelAnswerTask8 *widget.Label) {
	// ковертация числа "a" из строки в числовое значение (int)
	a, err := strconv.ParseFloat(entryTask8NumA.Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask8.SetText("Введите корректное значение в число А")
		entryTask8NumA.Show()
		entryTask8NumB.Hide()
		entryTask8NumC.Hide()
		return
	}
	// ковертация числа "b" из строки в числовое значение (int)
	b, err := strconv.ParseFloat(entryTask8NumB.Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		labelAnswerTask8.SetText("Введите корректное значение в число B")
		entryTask8NumA.Hide()
		entryTask8NumB.Show()
		entryTask8NumC.Hide()
		return
	}

	if a == b || a == -b { // в случае, если число "a" равно по модулю "b"
		// формирование ответа в виде строки, конвератция выражениея "sin(2*(a*2+b*2))" из числового значения в строку
		labelAnswerTask8.SetText(fmt.Sprintf("%f",(math.Sin(2 * (a * 2 + b * 2)))))
		// убирает текст из полей ввода
		entryTask8NumA.SetText("")
		entryTask8NumB.SetText("")
		entryTask8NumC.SetText("")
	}else if a < b { 
		appWindows.OpenMessageWindow(application, "осень")
		// убирает текст из полей ввода
		entryTask8NumA.SetText("")
		entryTask8NumB.SetText("")
		entryTask8NumC.SetText("")
		// отчистка ответа 
		labelAnswerTask8.SetText("")
	}else if a < 0 && b < 0 && entryTask8NumC.Text == "" { 
		// скрываются поля ввода для число "a" и "b"
		entryTask8NumA.Hide()
		entryTask8NumB.Hide()
		// показывается поле ввода для числа "c"
		entryTask8NumC.Show()
		// в ответ выводится текст с просьбой ввести число "c" 
		labelAnswerTask8.SetText("Введите число С")
	}else if a < 0 && b < 0 && entryTask8NumC.Text != "" { 
		// ковертация числа "c" из строки в числовое значение (int)
		c, err := strconv.ParseFloat(entryTask8NumC.Text, 64)
		// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
		if err != nil {
			labelAnswerTask8.SetText("Введите корректное значение в число C")
			entryTask8NumA.Hide()
			entryTask8NumB.Hide()
			entryTask8NumC.Show()
			return
		}
		// формирование ответа в виде строки, конвератция выражениея "a+b-c" из числового значения в строку
		labelAnswerTask8.SetText(fmt.Sprintf("%f",(a + b - c)))
		// возвращает изноальную видимость полей ввода
		entryTask8NumA.Show()
		entryTask8NumB.Show()
		entryTask8NumC.Hide()
		// убирает текст из полей ввода
		entryTask8NumA.SetText("")
		entryTask8NumB.SetText("")
		entryTask8NumC.SetText("")
		appWindows.OpenMessageWindow(application, "конец программы")
	}
}