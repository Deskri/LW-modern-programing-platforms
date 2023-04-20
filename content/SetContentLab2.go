package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"MODERN_PROGRAMMING_PLATFORMS/work"
)

func SetContentLab2(window fyne.Window) {
	//описание задания для заголовков выпадающих списков
	textInfoWork1Task1 := "Задание 1. Получить строку от пользователя, выделить из введенной строки слова и вывести их в обратном порядке."
	textInfoWork1Task2 := "Задание 2. Считать c клавиатуры три числа. Далее программа выполнит суммирование всех чисел и выведит результат на экран."
	textInfoWork1Task3 := "Задание 3. Нахождение гипотенузы и площади прямоугольного треугольника по двум катетам."
	textInfoWork1Task4 := "Задание 4. Вычисление среднего арифметического трех чисел."
	textInfoWork1Task5 := "Задание 5. Вычисления периметра и площади прямоугольного треугольного по двум катетам"
	textInfoWork1Task6 := "Задание 6. Программа запрашивает у пользователя a,b,c вычислить значение выражения а^2-(b-c)^2+ln(b^2)"
	textInfoWork1Task7 := "Задание 7. Программа запрашивает пять чисел. Производится вычисление средне-арифметического и вывод его на экран."
	textInfoWork1Task8 := "Задание 8. Реализовать программу вычисления расстояния между двумя точками по поданным на вход парам координат начала и конца."
	
	//ФОРМА ЗАДАНИЕ 1
	// поле ввода
	entryInvertedWordTask1 := widget.NewEntry()
	// лейбл вывода ответа
	labelAnswerInvertedWordTask1 := widget.NewLabel("Здесь будет Ваш ответ...")
	//создание формы
	formInvertedWordTask1 := widget.NewForm(
		widget.NewFormItem("Введите слово:", entryInvertedWordTask1),
		widget.NewFormItem("Ответ:", labelAnswerInvertedWordTask1),
	)
	// присваивание кнопки названия "Выполнить"
	formInvertedWordTask1.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formInvertedWordTask1.OnSubmit = func() {
		labelAnswerInvertedWordTask1.SetText(
			work.InvertedWords(
				entryInvertedWordTask1.Text))
	}

	//ФОРМА ЗАДАНИЕ 2
	// поле ввода
	entrySumOf3NumTask2Num1 := widget.NewEntry()  // первое число
	entrySumOf3NumTask2Num2 := widget.NewEntry()  // второе чилсо
	entrySumOf3NumTask2Num3 := widget.NewEntry()  // третье число
	// лейбл вывода ответа
	labelAnswerSumOf3NumTask2 := widget.NewLabel("Здесь будет сумма...")
	//создание формы
	formSumOf3NumTask2 := widget.NewForm(
		widget.NewFormItem("Введите первое число:", entrySumOf3NumTask2Num1),
		widget.NewFormItem("Введите второе число:", entrySumOf3NumTask2Num2),
		widget.NewFormItem("Введите третье число:", entrySumOf3NumTask2Num3),
		widget.NewFormItem("Сумма чисел:", labelAnswerSumOf3NumTask2),
	)
	// присваивание кнопки названия "Выполнить"
	formSumOf3NumTask2.SubmitText = "Выполнить"

	//функция выполняемая при нажатии кнопки "Выполнить"
	formSumOf3NumTask2.OnSubmit = func() {
		labelAnswerSumOf3NumTask2.SetText(
			work.SumOf3Num(
				entrySumOf3NumTask2Num1.Text, 
				entrySumOf3NumTask2Num2.Text, 
				entrySumOf3NumTask2Num3.Text))
	}

	//ФОРМА ЗАДАНИЕ 3
	// поле ввода
	entryHypotenuseAndAreaTask3Cathetus1 := widget.NewEntry()  // первый катет
	entryHypotenuseAndAreaTask3Cathetus2 := widget.NewEntry()  // второй катет
	// лейбл вывода ответа
	labelAnswerHypotenuseTask3 := widget.NewLabel("Здесь будет гипотенуза ...") // гипотенуза
	labelAnswerAreaTask3 := widget.NewLabel("Здесь будеn площадь ...") // площадь

	//создание формы
	formHypotenuseAndAreaTask3 := widget.NewForm(
		widget.NewFormItem("Введите первой катет:", entryHypotenuseAndAreaTask3Cathetus1),
		widget.NewFormItem("Введите второй катет:", entryHypotenuseAndAreaTask3Cathetus2),
		widget.NewFormItem("Гипотенуза:", labelAnswerHypotenuseTask3),
		widget.NewFormItem("Площадь:", labelAnswerAreaTask3),
	)
	// присваивание кнопки названия "Выполнить"
	formHypotenuseAndAreaTask3.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formHypotenuseAndAreaTask3.OnSubmit = func() {
		labelAnswerHypotenuseTask3.SetText(
			work.Hypotenuse(
				entryHypotenuseAndAreaTask3Cathetus1.Text, 
				entryHypotenuseAndAreaTask3Cathetus2.Text))
		labelAnswerAreaTask3.SetText(
			work.Area(
				entryHypotenuseAndAreaTask3Cathetus1.Text, 
				entryHypotenuseAndAreaTask3Cathetus2.Text))
	}

	//ФОРМА ЗАДАНИЕ 4
	// поле ввода
	entryArithmeticMeanTask4Num1 := widget.NewEntry()  // первое число
	entryArithmeticMeanTask4Num2 := widget.NewEntry()  // второе число
	entryArithmeticMeanTask4Num3 := widget.NewEntry()  // третье число
	// лейбл вывода ответа
	labelAnswerArithmeticMeanTask4 := widget.NewLabel("Здесь будет ответ ...")

	//создание формы
	formArithmeticMeanTask4 := widget.NewForm(
		widget.NewFormItem("Введите первой число:", entryArithmeticMeanTask4Num1),
		widget.NewFormItem("Введите второй число:", entryArithmeticMeanTask4Num2),
		widget.NewFormItem("Введите третье число:", entryArithmeticMeanTask4Num3),	
		widget.NewFormItem("Среднее арифметическое:", labelAnswerArithmeticMeanTask4),
	)
	// присваивание кнопки названия "Выполнить"
	formArithmeticMeanTask4.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formArithmeticMeanTask4.OnSubmit = func() {
		labelAnswerArithmeticMeanTask4.SetText(
			work.ArithmeticMean(
				entryArithmeticMeanTask4Num1.Text,
				entryArithmeticMeanTask4Num2.Text, 
				entryArithmeticMeanTask4Num3.Text))
	}

	//ФОРМА ЗАДАНИЕ 5
	// поле ввода
	entryPerimeterAndAreaTask5cathetus1 := widget.NewEntry()  // первый катет
	entryPerimeterAndAreaTask5cathetus2 := widget.NewEntry()  // второй катет
	// лейбл вывода ответа
	labelAnswerPerimeterTask5 := widget.NewLabel("Здесь будет периметр ...") // периметр
	labelAnswerAreaTask5 := widget.NewLabel("Здесь будеn площадь ...") // площадь

	//создание формы
	formPerimeterAndAreaTask5 := widget.NewForm(
		widget.NewFormItem("Введите первой катет:", entryPerimeterAndAreaTask5cathetus1),
		widget.NewFormItem("Введите второй катет:", entryPerimeterAndAreaTask5cathetus2),
		widget.NewFormItem("Периметр:", labelAnswerPerimeterTask5),
		widget.NewFormItem("Площадь:", labelAnswerAreaTask5),
	)
	// присваивание кнопки названия "Выполнить"
	formPerimeterAndAreaTask5.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formPerimeterAndAreaTask5.OnSubmit = func() {
		labelAnswerPerimeterTask5.SetText(
			work.Perimeter(
				entryPerimeterAndAreaTask5cathetus1.Text, 
				entryPerimeterAndAreaTask5cathetus2.Text))
		labelAnswerAreaTask5.SetText(
			work.Area(
				entryPerimeterAndAreaTask5cathetus1.Text, 
				entryPerimeterAndAreaTask5cathetus2.Text))
	}

	//ФОРМА ЗАДАНИЕ 6
	// поле ввода
	entryCalculatingAnExpressionTask6NumA := widget.NewEntry()  // число a
	entryCalculatingAnExpressionTask6NumB := widget.NewEntry()  // число b
	entryCalculatingAnExpressionTask6NumC := widget.NewEntry()  // число с
	// лейбл вывода ответа
	labelAnswerCalculatingAnExpressionTask6 := widget.NewLabel("Здесь будет ответ ...")

	//создание формы
	formCalculatingAnExpressionTask6 := widget.NewForm(
		widget.NewFormItem("Введите число a:", entryCalculatingAnExpressionTask6NumA),
		widget.NewFormItem("Введите число b:", entryCalculatingAnExpressionTask6NumB),
		widget.NewFormItem("Введите число c:", entryCalculatingAnExpressionTask6NumC),
		widget.NewFormItem("Ответ:", labelAnswerCalculatingAnExpressionTask6),
	)
	// присваивание кнопки названия "Выполнить"
	formCalculatingAnExpressionTask6.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formCalculatingAnExpressionTask6.OnSubmit = func() {
		labelAnswerCalculatingAnExpressionTask6.SetText(
			work.CalculatingAnExpression(
				entryCalculatingAnExpressionTask6NumA.Text, 
				entryCalculatingAnExpressionTask6NumB.Text, 
				entryCalculatingAnExpressionTask6NumC.Text))
	}

	//ФОРМА ЗАДАНИЕ 7
	// поле ввода
	entryArithmeticMeanOfTheFiveNumbersTask7Num1 := widget.NewEntry()  // первое число
	entryArithmeticMeanOfTheFiveNumbersTask7Num2 := widget.NewEntry()  // второе число
	entryArithmeticMeanOfTheFiveNumbersTask7Num3 := widget.NewEntry()  // третье число
	entryArithmeticMeanOfTheFiveNumbersTask7Num4 := widget.NewEntry()  // четвертое число
	entryArithmeticMeanOfTheFiveNumbersTask7Num5 := widget.NewEntry()  // пятое число
	// лейбл вывода ответа
	labelAnswerArithmeticMeanOfTheFiveNumbersTask7 := widget.NewLabel("Здесь будет ответ ...")

	//создание формы
	formArithmeticMeanOfTheFiveNumbersTask7 := widget.NewForm(
		widget.NewFormItem("Введите первой число:", entryArithmeticMeanOfTheFiveNumbersTask7Num1),
		widget.NewFormItem("Введите второй число:", entryArithmeticMeanOfTheFiveNumbersTask7Num2),
		widget.NewFormItem("Введите третье число:", entryArithmeticMeanOfTheFiveNumbersTask7Num3),
		widget.NewFormItem("Введите четвертое число:", entryArithmeticMeanOfTheFiveNumbersTask7Num4),
		widget.NewFormItem("Введите пятое число:", entryArithmeticMeanOfTheFiveNumbersTask7Num5), 
		widget.NewFormItem("Среднее арифметическое:", labelAnswerArithmeticMeanOfTheFiveNumbersTask7),
	)
	// присваивание кнопки названия "Выполнить"
	formArithmeticMeanOfTheFiveNumbersTask7.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formArithmeticMeanOfTheFiveNumbersTask7.OnSubmit = func() {
		labelAnswerArithmeticMeanOfTheFiveNumbersTask7.SetText(
			work.ArithmeticMeanOfTheFiveNumbers(
				entryArithmeticMeanOfTheFiveNumbersTask7Num1.Text,
				entryArithmeticMeanOfTheFiveNumbersTask7Num2.Text,
				entryArithmeticMeanOfTheFiveNumbersTask7Num3.Text,
				entryArithmeticMeanOfTheFiveNumbersTask7Num4.Text,
				entryArithmeticMeanOfTheFiveNumbersTask7Num5.Text))
	}

	//ФОРМА ЗАДАНИЕ 8
	// поле ввода
	entryCoordinateDistanceTask8Coordinate1X := widget.NewEntry()  // X первой координаты
	entryCoordinateDistanceTask8Coordinate1Y := widget.NewEntry()  // Y первой координаты
	entryCoordinateDistanceTask8Coordinate2X := widget.NewEntry()  // X второй координаты
	entryCoordinateDistanceTask8Coordinate2Y := widget.NewEntry()  // Y второй координаты
	// лейбл вывода ответа
	labelAnswerCoordinateDistanceTask8 := widget.NewLabel("Здесь будет ответ ...")

	//создание формы
	formCoordinateDistanceTask8 := widget.NewForm(
		widget.NewFormItem("Введите X для первой координаты:", entryCoordinateDistanceTask8Coordinate1X),
		widget.NewFormItem("Введите Y для первой координаты:", entryCoordinateDistanceTask8Coordinate1Y),
		widget.NewFormItem("Введите X для второй координаты:", entryCoordinateDistanceTask8Coordinate2X),
		widget.NewFormItem("Введите Y для второй координаты:", entryCoordinateDistanceTask8Coordinate2Y),
		widget.NewFormItem("Расстояние:", labelAnswerCoordinateDistanceTask8),
	)
	// присваивание кнопки названия "Выполнить"
	formCoordinateDistanceTask8.SubmitText = "Выполнить"
	//функция выполняемая при нажатии кнопки "Выполнить"
	formCoordinateDistanceTask8.OnSubmit = func() {
		labelAnswerCoordinateDistanceTask8.SetText(
			work.CoordinateDistance(
				entryCoordinateDistanceTask8Coordinate1X.Text,
				entryCoordinateDistanceTask8Coordinate1Y.Text,
				entryCoordinateDistanceTask8Coordinate2X.Text,
				entryCoordinateDistanceTask8Coordinate2Y.Text))
	}
	// создание элемента выпадающего списка с текстом задания 1 и исполняющей формой задания
	accordionItemTask1 := widget.NewAccordionItem(
		textInfoWork1Task1, //текст задания
		formInvertedWordTask1, //форма
	)
	// создание элемента выпадающего списка с текстом задания 2 и исполняющей формой задания
	accordionItemTask2 := widget.NewAccordionItem(
		textInfoWork1Task2, //текст задания
		formSumOf3NumTask2, //форма
	)
	// создание элемента выпадающего списка с текстом задания 3 и исполняющей формой задания
	accordionItemTask3 := widget.NewAccordionItem(
		textInfoWork1Task3, //текст задания
		formHypotenuseAndAreaTask3, //форма
	)
	// создание элемента выпадающего списка с текстом задания 4 и исполняющей формой задания
	accordionItemTask4 := widget.NewAccordionItem(
		textInfoWork1Task4, //текст задания
		formArithmeticMeanTask4, //форма
	)
	// создание элемента выпадающего списка с текстом задания 5 и исполняющей формой задания
	accordionItemTask5 := widget.NewAccordionItem(
		textInfoWork1Task5, //текст задания
		formPerimeterAndAreaTask5, //форма
	)
	// создание элемента выпадающего списка с текстом задания 6 и исполняющей формой задания
	accordionItemTask6 := widget.NewAccordionItem(
		textInfoWork1Task6, //текст задания
		formCalculatingAnExpressionTask6, //форма
	)
	// создание элемента выпадающего списка с текстом задания 7 и исполняющей формой задания
	accordionItemTask7 := widget.NewAccordionItem(
		textInfoWork1Task7, //текст задания
		formArithmeticMeanOfTheFiveNumbersTask7, //форма
	)
	// создание элемента выпадающего списка с текстом задания 8 и исполняющей формой задания
	accordionItemTask8 := widget.NewAccordionItem(
		textInfoWork1Task8, //текст задания
		formCoordinateDistanceTask8, //форма
	)
	// создание вапающего списка
	accordion := widget.NewAccordion(
		accordionItemTask1, // элемент выпадающего списка задания 1
		accordionItemTask2, // элемент выпадающего списка задания 2
		accordionItemTask3, // элемент выпадающего списка задания 3
		accordionItemTask4, // элемент выпадающего списка задания 4
		accordionItemTask5, // элемент выпадающего списка задания 5
		accordionItemTask6, // элемент выпадающего списка задания 6
		accordionItemTask7, // элемент выпадающего списка задания 7
		accordionItemTask8, // элемент выпадающего списка задания 8
	)
	// размещение элементов в основное окно с использованием контейнера с прокруткой
	window.SetContent(container.NewVScroll(container.NewVBox(
		accordion)))
}