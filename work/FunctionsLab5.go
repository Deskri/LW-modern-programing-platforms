package work

import (
	"strconv"
)

func Lab5Task5 (someText string) (answer string) {
	// создается массив из 5 чисел
	var array [5]int
	// создание числовой переменой счетчика интекса массива
	var counterOfArrayIndex int = 0
	// создание текстовой переменой числа
	var numberToSetInsideArray string = ""
	// создание числовой переменной колличества четных чисел
	var countOfEvenNumbers int = 0
	// присваивание ответу пустого значения
	answer = ""
	// перебор текста для формирование массивов чисел
	for i := 0; i < len(someText); i++ {
		if someText[i] != '0' &&
			someText[i] != '1' &&
			someText[i] != '2' &&
			someText[i] != '3' &&
			someText[i] != '4' &&
			someText[i] != '5' &&
			someText[i] != '6' &&
			someText[i] != '7' &&
			someText[i] != '8' &&
			someText[i] != '9' &&
			someText[i] != ' ' { // проверка каждого элемента строки на остуствие недопустимых символов
				answer = "Введите корректное значение"
				return	
		}else if someText[i] == '0' ||
			someText[i] == '1' ||
			someText[i] == '2' ||
			someText[i] == '3' ||
			someText[i] == '4' ||
			someText[i] == '5' ||
			someText[i] == '6' ||
			someText[i] == '7' ||
			someText[i] == '8' ||
			someText[i] == '9' { // если элемент строки равен числу, то текстовой переменной числа присвивается значение
			numberToSetInsideArray += string(someText[i])
		}
		if someText[i] == ' ' || i == len(someText) - 1 { // если элемент строки, это пробел или последний элемент строки, то...
			if (counterOfArrayIndex > 4) { // происходит проверка на колличество введеных чисел в массив
				answer = "В массиве больше 5 чисел"
				return	
			}else { // если прошли проверку, то ...
				var err error
				array[counterOfArrayIndex], err = strconv.Atoi(numberToSetInsideArray) // конвертируем строку в число и вносим в массив
				// в случае, если в строке присутствуют иные символы, помимо цифр, формируется текст ошибки
				if err != nil { // если не смогли конвертировать, то выдаем ошибку на какой позиции ошибка
					answer  = "В массиве на позиции " + strconv.Itoa(counterOfArrayIndex) + " не допустимое значение"
					return
				}
				// после внесения нового элемента в массив, увеличиваем значения элементов в масиве
				counterOfArrayIndex++
				// текстовой переменой числа присваеваем пустое значени
				numberToSetInsideArray = ""
			}
		}
	}
	// формируем ответ, где выводим только четные числа массива
	for i := 0; i < len(array); i++ {
		// условие, если элемента массива четный и массив не = 0, то ...
		if array[i] % 2 == 0 && array[i] != 0 {
			//если это первое четное число, но конвертируем и дополняем в ответ
			if countOfEvenNumbers == 0 {
				answer += strconv.Itoa(array[i])
			}else {// в остальных случаях, те же самые опперации, но перед числом вставляем проблел
				answer += " " + strconv.Itoa(array[i])
			}
			countOfEvenNumbers++
		}
	}
	// возвращаем ответ
	return answer
}