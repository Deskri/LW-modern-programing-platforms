package work

import (
	"strconv"
)

func Lab4Task5(someText string) (answer string) {
	// создается числовая переменная счетчика чисел в строке
	var counterOfNumbers int = 0
	// перебор всего входящего текста
	for i := 0; i < len(someText); i++ {
		if someText[i] == '0' ||
			someText[i] == '1' ||
			someText[i] == '2' ||
			someText[i] == '3' ||
			someText[i] == '4' ||
			someText[i] == '5' ||
			someText[i] == '6' ||
			someText[i] == '7' ||
			someText[i] == '8' ||
			someText[i] == '9' ||
			someText[i] == '+' ||
			someText[i] == '-' ||
			someText[i] == '*' {// условие, если равен элемент строки числу или знаком операций, то...
				counterOfNumbers++ // увеличиваем на 1 колличество чисел
			}
	}
	// конвертируем значение колличества чиел в строке и присваем в ответ
	answer = strconv.Itoa(counterOfNumbers)
	// возвращаем ответ
	return answer
}