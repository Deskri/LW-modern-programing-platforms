package work

import (
	"strconv"
	"math"
	"fmt"
)

// функциия инверсии строки (переворот), на вход поступает строка, на выходе измененая строка
func InvertedWords (someText string) (answer string) {
	answer = ""
	for _, letter := range someText {
		answer = string(letter) + answer
	}
	return
}
// функция подсчета суммы 3-ех чисел, на вход поступают текстовом 3 числа, на выходе сумма текстом
func SumOf3Num (num1Text, num2Text, num3Text string) (answer string) {
	// ковертация числа 1 из строки в числовое значение (int)
	num1, err := strconv.Atoi(num1Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение в первое число"
		return
	}
	// ковертация числа 2 из строки в числовое значение (int)
	num2, err := strconv.Atoi(num2Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение во второе число"
		return
	}
	// ковертация числа 3 из строки в числовое значение (int)
	num3, err := strconv.Atoi(num3Text)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение в третье число"
		return
	}
	// присвивние ответу суммы 3-ех чисел и ковертация в строку этой суммы
	answer = strconv.Itoa(num1 + num2 + num3)
	// возвращение ответа
	return
}
// функция расчета гипотенузы прямоугольного треугольника, на вход поступают два катета текстом, на выходе гипотенуза текстом 
func Hypotenuse (cathetus1Text, cathetus2Text string) (answer string) {
	// ковертация первого катета из строки в числовое значение с плавующей точкой (float64)
	cathetus1, err := strconv.ParseFloat(cathetus1Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение первого катета"
		return
	}
	// ковертация второго катета из строки в числовое значение с плавующей точкой (float64)
	cathetus2, err := strconv.ParseFloat(cathetus2Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение второго катета"
		return
	}
	// присвивание ответу расчета гипотенузы и конвертация расчета в строку
	answer = fmt.Sprintf("%f", math.Sqrt(math.Pow(cathetus1, 2) + math.Pow(cathetus2, 2)))
	// возвращение ответа
	return
}
// функция расчета площади прямоугольного треугольника, на вход поступают два катета текстом, на выходе площадь текстом 
func Area (cathetus1Text, cathetus2Text string) (answer string) {
	// ковертация первого катета из строки в числовое значение с плавующей точкой (float64)
	cathetus1, err := strconv.ParseFloat(cathetus1Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение первого катета"
		return
	}
	// ковертация второго катета из строки в числовое значение с плавующей точкой (float64)
	cathetus2, err := strconv.ParseFloat(cathetus2Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение второго катета"
		return
	}
	// присвивание ответу расчета площади и конвертация расчета в строку
	answer = fmt.Sprintf("%f", (cathetus1 * cathetus2) / 2)
	// возвращение ответа
	return
}
// функция расчета периметра прямоугольного треугольника, на вход поступают два катета текстом, на выходе периметр текстом
func Perimeter (cathetus1Text, cathetus2Text string) (answer string) {
	// ковертация первого катета из строки в числовое значение с плавующей точкой (float64)
	cathetus1, err := strconv.ParseFloat(cathetus1Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение первого катета"
		return
	}
	// ковертация второго катета из строки в числовое значение с плавующей точкой (float64)
	cathetus2, err := strconv.ParseFloat(cathetus2Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение второго катета"
		return
	}
	// вызов функции расчета гипотенузы
	hypotenuse, err := 	strconv.ParseFloat(Hypotenuse(cathetus1Text, cathetus2Text), 64)
	// в случае, если не удается привезти расчет гипотенузы к число, возвращает ошибку
	if err != nil {
		answer  = "Не получилось привезти гипотенузу к числу"
		return
	}
	// присвивание ответу расчета периметра и конвертация расчета в строку
	answer = fmt.Sprintf("%f", cathetus1 + cathetus2 + hypotenuse)
	// возвращение ответа
	return
}
// функциия расчета среднего значения 3-ех чисел, на вход поступают 3 числа текстом, на выходе среднее значение текстом
func ArithmeticMean (num1Text, num2Text, num3Text string) (answer string) {
	// ковертация числа 1 из строки в числовое значение с плавующей точкой (float64)
	num1, err := strconv.ParseFloat(num1Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение в первое число"
		return
	}
	// ковертация числа 2 из строки в числовое значение с плавующей точкой (float64)
	num2, err := strconv.ParseFloat(num2Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение во второе число"
		return
	}
	// ковертация числа 3 из строки в числовое значение с плавующей точкой (float64)
	num3, err := strconv.ParseFloat(num3Text, 64)
	// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
	if err != nil {
		answer  = "Введите корректное значение в третье число"
		return
	}
	// присваивание ответу расчета среднего значения и конвертация среднего значения в строку
	answer = fmt.Sprintf("%f", (num1 + num2 + num3) / 3)
	// возвращение ответа
	return
}
// функция расчета выражения, на вход поступают 3 числа текстом, на выходе результат выражения текстом
func CalculatingAnExpression (aText, bText, cText string) (answer string) {
	a, err := strconv.ParseFloat(aText , 64)

	if err != nil {
		answer  = "Введите корректное значение в число а"
		return
	}

	b, err := strconv.ParseFloat(bText, 64)

	if err != nil {
		answer  = "Введите корректное значение в число b"
		return
	}

	c, err := strconv.ParseFloat(cText, 64)

	if err != nil {
		answer  = "Введите корректное значение в число c"
		return
	}
	//а^2-(b-c)^2+ln(b^2)
	answer = fmt.Sprintf("%f", math.Pow(a, 2) - math.Pow((b - c), 2) + math.Log(math.Pow(b, 2)))
	return
}
// функция расчета среднего значения 5-ти чисел, на вход поступают 5 числа текстом, на выходе среднее значение текстом
func ArithmeticMeanOfTheFiveNumbers (num1Text, num2Text, num3Text, num4Text, num5Text string) (answer string) {
	num1, err := strconv.ParseFloat(num1Text, 64)

	if err != nil {
		answer  = "Введите корректное значение в первое число"
		return
	}

	num2, err := strconv.ParseFloat(num2Text, 64)

	if err != nil {
		answer  = "Введите корректное значение во второе число"
		return
	}

	num3, err := strconv.ParseFloat(num3Text, 64)

	if err != nil {
		answer  = "Введите корректное значение в третье число"
		return
	}

	num4, err := strconv.ParseFloat(num4Text, 64)

	if err != nil {
		answer  = "Введите корректное значение в четвертое число"
		return
	}
	num5, err := strconv.ParseFloat(num5Text, 64)

	if err != nil {
		answer  = "Введите корректное значение в пятое число"
		return
	}
	answer = fmt.Sprintf("%f", (num1 + num2 + num3 + num4 + num5) / 5)
	return
}
// функция расчета расстояния между двумя координатами, на вход подаются точка абсциз и точка ординат первой и второй координаты текстом,
// на выходе расстояние текстом
func CoordinateDistance (x1Text, y1Text, x2Text, y2Text string) (answer string) {
	x1, err := strconv.ParseFloat(x1Text, 64)

	if err != nil {
		answer  = "Введите корректное значение x1"
		return
	}

	y1, err := strconv.ParseFloat(y1Text, 64)

	if err != nil {
		answer  = "Введите корректное значение y1"
		return
	}

	x2, err := strconv.ParseFloat(x2Text, 64)

	if err != nil {
		answer  = "Введите корректное значение x2"
		return
	}

	y2, err := strconv.ParseFloat(y2Text, 64)

	if err != nil {
		answer  = "Введите корректное значение y2"
		return
	}

	var cathetus1 float64
	if (x1 > x2) {
		cathetus1 = x1 - x2
	}else if (x2 > x1) {
		cathetus1 = x2 - x1
	}else {
		cathetus1 = 0
	}

	var cathetus2 float64
	if (y1 > y2) {
		cathetus2 = y1 - y2
	} else if (y2 > y1) {
		cathetus2 = y2 - y1
	} else {
		cathetus2 = 0
	}

	answer = Hypotenuse(fmt.Sprintf("%f", cathetus1), fmt.Sprintf("%f",cathetus2))
	return
}