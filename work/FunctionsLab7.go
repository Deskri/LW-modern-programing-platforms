package work

import (
	"io/ioutil"
	"strconv"
	"fyne.io/fyne/v2"
	"MODERN_PROGRAMMING_PLATFORMS/content/appWindows"
	"fyne.io/fyne/v2/widget"
)

func EmployeeTestingProcessing(
	radioGroupQuestion1, 
	radioGroupQuestion2, 
	radioGroupQuestion3, 
	radioGroupQuestion4, 
	radioGroupQuestion5 *widget.RadioGroup,
	labelQuestion1,
	labelQuestion2,
	labelQuestion3,
	labelQuestion4,
	labelQuestion5 *widget.Label,
	application fyne.App,
) {

	var textToFileResultOfTest string = ""

	arrayRadioGroupQuestion := []*widget.RadioGroup {
		radioGroupQuestion1, 
		radioGroupQuestion2, 
		radioGroupQuestion3, 
		radioGroupQuestion4, 
		radioGroupQuestion5,
	}

	arrayLabelQuestion := []*widget.Label {
		labelQuestion1,
		labelQuestion2,
		labelQuestion3,
		labelQuestion4,
		labelQuestion5,
	}

	for i := 0; i < len(arrayRadioGroupQuestion); i++ {
		if arrayRadioGroupQuestion[i].Selected == "" {
			appWindows.OpenMessageWindow(application, "Заполните вопрос №" + strconv.Itoa(i + 1))
			return
		}else {
			textToFileResultOfTest += arrayLabelQuestion[i].Text + " - " + arrayRadioGroupQuestion[i].Selected + "\n"
		}
	}

	dataFromFile, err := ioutil.ReadFile("temp/testResults/resultsOfTest.txt")
	
	if err != nil {
		appWindows.OpenMessageWindow(application, "Ошибка при чтении файла")
	}
	textToFileResultOfTest = string(dataFromFile) + textToFileResultOfTest
	ioutil.WriteFile("temp/testResults/resultsOfTest.txt",[]byte(textToFileResultOfTest), 0777)

	appWindows.OpenMessageWindow(application, "Тестирование пройдено!")
}

func ParseTestResult(
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
	labelQuestion5 *widget.Label,
	application fyne.App,
) {
	arrayLabelQuestionNumber := []*widget.Label {
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
	}

	numberAnswerQuestion1Yes := 0
	numberAnswerQuestion1No := 0
	numberAnswerQuestion1Abstain := 0
	numberAnswerQuestion2Yes := 0
	numberAnswerQuestion2No := 0
	numberAnswerQuestion2Abstain := 0
	numberAnswerQuestion3Yes := 0
	numberAnswerQuestion3No := 0
	numberAnswerQuestion3Abstain := 0
	numberAnswerQuestion4Yes := 0
	numberAnswerQuestion4No := 0
	numberAnswerQuestion4Abstain := 0
	numberAnswerQuestion5Yes := 0
	numberAnswerQuestion5No := 0
	numberAnswerQuestion5Abstain := 0

	arrayNumberOfMakeChoseTheAnswer := []int {
		numberAnswerQuestion1Yes,
		numberAnswerQuestion1No,
		numberAnswerQuestion1Abstain,
		numberAnswerQuestion2Yes,
		numberAnswerQuestion2No,
		numberAnswerQuestion2Abstain,
		numberAnswerQuestion3Yes,
		numberAnswerQuestion3No,
		numberAnswerQuestion3Abstain,
		numberAnswerQuestion4Yes,
		numberAnswerQuestion4No,
		numberAnswerQuestion4Abstain,
		numberAnswerQuestion5Yes,
		numberAnswerQuestion5No,
		numberAnswerQuestion5Abstain,
	}

	dataFromFile, err := ioutil.ReadFile("temp/testResults/resultsOfTest.txt")
	
	if err != nil {
		appWindows.OpenMessageWindow(application, "Ошибка при чтении файла")
	}

	textFromFile := string(dataFromFile)

	var currentWord string

	var isQuestion1 bool = false
	var isQuestion2 bool = false
	var isQuestion3 bool = false
	var isQuestion4 bool = false
	var isQuestion5 bool = false

	for _, letter := range textFromFile {
		currentWord += string(letter)
		if currentWord == labelQuestion1.Text {
			currentWord = ""
			isQuestion1 = true
		} else if currentWord == labelQuestion2.Text {
			currentWord = ""
			isQuestion2 = true
		} else if currentWord == labelQuestion3.Text {
			currentWord = ""
			isQuestion3 = true
		} else if currentWord == labelQuestion4.Text {
			currentWord = ""
			isQuestion4 = true
		} else if currentWord == labelQuestion5.Text {
			currentWord = ""
			isQuestion5 = true
		} else if currentWord == " - " {
			currentWord = ""
		} else if currentWord == "\n" {
			currentWord = ""
		} else if currentWord == "Да" && isQuestion1 {
			isQuestion1 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[0]++
		} else if currentWord == "Нет" && isQuestion1 {
			isQuestion1 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[1]++
		} else if currentWord == "Воздержусь" && isQuestion1 {
			isQuestion1 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[2]++
		} else if currentWord == "Да" && isQuestion2 {
			isQuestion2 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[3]++
		} else if currentWord == "Нет" && isQuestion2 {
			isQuestion2 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[4]++
		} else if currentWord == "Воздержусь" && isQuestion2 {
			isQuestion2 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[5]++
		} else if currentWord == "Да" && isQuestion3 {
			isQuestion3 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[6]++
		} else if currentWord == "Нет" && isQuestion3 {
			isQuestion3 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[7]++
		} else if currentWord == "Воздержусь" && isQuestion3 {
			isQuestion3 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[8]++
		} else if currentWord == "Да" && isQuestion4 {
			isQuestion4 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[9]++
		} else if currentWord == "Нет" && isQuestion4 {
			isQuestion4 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[10]++
		} else if currentWord == "Воздержусь" && isQuestion4 {
			isQuestion4 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[11]++
		} else if currentWord == "Да" && isQuestion5 {
			isQuestion5 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[12]++
		} else if currentWord == "Нет" && isQuestion5 {
			isQuestion5 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[13]++
		} else if currentWord == "Воздержусь" && isQuestion5 {
			isQuestion5 = false
			currentWord = ""
			arrayNumberOfMakeChoseTheAnswer[14]++
		}
	}
	for i := 0; i < len(arrayLabelQuestionNumber); i++ {
		arrayLabelQuestionNumber[i].SetText(strconv.Itoa(arrayNumberOfMakeChoseTheAnswer[i]))
	}
}