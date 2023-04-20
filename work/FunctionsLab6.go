package work

import (
	"strconv"
	"unicode"
)

func Lab6Task5(someText string) (answer string) {
	// переменая для подсчета длины массива
	var countOfArrayIndex int = 0
	// переменная для определения открыта ли скобка
	var isOpen bool = false
	// определение массивов строк
	for _, letter := range someText { // перебираем каждый символ в передаваемом тексте
		if letter == '"' && !isOpen{ // если попадается в тексте скобка и переменной открытии скобы не присвоено значение, то присваевамием значение "true"
			isOpen = true
		}else if letter == '"'&& isOpen{ // если попадается в тексте скобка и переменной открытии скобы присвоено значение, то присваевамием значение "false" и записываем +1 колличество массивов строк
			isOpen = false
			countOfArrayIndex++
		}
	}
	// в случае, если была ошибка ввода и скоба не закрыта выведет ошибку
	if isOpen{
		answer = "Ошибка: массив не закрыт"
		return
	}
	// если колличество массивов строк 0, просит ввести массив/массивы
	if countOfArrayIndex == 0 {
		answer = "Введите массив/массивы"
		return
	}
	// создается срез пустых строк
	arrayString := make([]string, countOfArrayIndex)
	// создается срез нулевых значений для подсчета заглавных букв
	arrayStringUpper := make([]int, countOfArrayIndex)
	// обнуляется счетчик колличества строк
	countOfArrayIndex = 0
	// внесение данных в массивы из строки переданной в функцию
	for _, letter := range someText {
		if letter == '"' && !isOpen{
			isOpen = true
		}else if letter == '"'&& isOpen {
			isOpen = false
			countOfArrayIndex++
		}else if isOpen { // если скобка открыта, переносятся руны из текста в массив строки
			arrayString[countOfArrayIndex] += string(letter)
		}
	}
	// определение колличества заглавных букв
	for i := 0; i < len(arrayString); i++ {
		// перебор строки по рунам
		for _, someRune := range arrayString[i] {
			if unicode.IsUpper(someRune) { // если буква в верхнем регистре
				arrayStringUpper[i]++  // то число заглавных букв увеличивается на 1
			}
		}
	}
	// формирование ответа
	for i := 0; i < len(arrayString); i++ {
		answer += "Массив " + strconv.Itoa(i) + ": длина строки - " + strconv.Itoa(len([]rune(arrayString[i]))) + "; колличество заглавных букв - " + strconv.Itoa(arrayStringUpper[i]) + "\n" 
	}
	// возвращаем ответ
	return answer
}