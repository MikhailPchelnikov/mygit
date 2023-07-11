package main

import (
	//"bufio"
	"fmt"
	"strings"
	//"os"
	"strconv"
)

func checkOpS(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	OpsSet := make(map[string]void)

	OpsSet["+"] = voidEl
	OpsSet["-"] = voidEl
	OpsSet["/"] = voidEl
	OpsSet["*"] = voidEl

	if _, ok := OpsSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}

func checkOperandsArab(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	ArabSet := make(map[string]void)

	//ArabSet["0"] = voidEl
	ArabSet["1"] = voidEl
	ArabSet["2"] = voidEl
	ArabSet["3"] = voidEl
	ArabSet["4"] = voidEl
	ArabSet["5"] = voidEl
	ArabSet["6"] = voidEl
	ArabSet["7"] = voidEl
	ArabSet["8"] = voidEl
	ArabSet["9"] = voidEl
	ArabSet["10"] = voidEl
	if _, ok := ArabSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}

func checkOperandsRome(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	RomeSet := make(map[string]void)

	RomeSet["I"] = voidEl
	RomeSet["II"] = voidEl
	RomeSet["III"] = voidEl
	RomeSet["IV"] = voidEl
	RomeSet["V"] = voidEl
	RomeSet["VI"] = voidEl
	RomeSet["VII"] = voidEl
	RomeSet["VIII"] = voidEl
	RomeSet["IX"] = voidEl
	RomeSet["X"] = voidEl

	if _, ok := RomeSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}

func intToRome100(i int) string {

	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman string
	for _, conversion := range conversions {
		for i >= conversion.value {
			roman = roman + conversion.digit
			i -= conversion.value
		}
	}

	return roman
}

const BasicErrorMsg = "Ваша строка не соотвествует формату: \" два операнда и один оператор между ними (через пробелы) \""

func main() {

	RomeMap := make(map[string]int)

	RomeMap["I"] = 1
	RomeMap["II"] = 2
	RomeMap["III"] = 3
	RomeMap["IV"] = 4
	RomeMap["V"] = 5
	RomeMap["VI"] = 6
	RomeMap["VII"] = 7
	RomeMap["VIII"] = 8
	RomeMap["IX"] = 9
	RomeMap["X"] = 10

	//reader := bufio.NewReader(os.Stdin)s
	var str1, str2, str3, str4, errorString string
	var op1, op2, res int

	str4 = "cheatcode543210string" //uniq string life user without special knowledge never enters
	errorString = BasicErrorMsg

	for {

		fmt.Println("Будет выведен результат, если введете \" два операнда и один оператор между ними (через пробелы) \"")
		fmt.Scanln(&str1, &str2, &str3, &str4)

		if str3 == "" {
			errorString = ("Учтите, Вы ввели слишком мало. или забыли про пробелы " + BasicErrorMsg)
			break
		}

		if str4 != "cheatcode543210string" {
			errorString = ("Учтите, Вы ввели слишком много. " + BasicErrorMsg)
			break
		}

		if !checkOpS(str2) {
			errorString = ("Учтите, Вы НЕ ввели арифметический знак из ряда + - * / . между операндами. И может еще что-то не так ввели " + BasicErrorMsg)
			break
		}

		str1 = strings.ToUpper(str1)
		str3 = strings.ToUpper(str3)

		if !(checkOperandsArab(str1) || checkOperandsRome(str1)) {
			errorString = ("Ваш как минимум первый операнд не вписывается в рабочий диапазон, вводите от 1 до 10 арабскими или римским (от I до X ). С нулем, отрицательными, 11 и более значениями приложение работать НЕ будет! " + BasicErrorMsg)
			break
		}

		if !(checkOperandsArab(str3) || checkOperandsRome(str3)) {
			errorString = ("Ваш второй операнд не вписывается в рабочий диапазон, вводите от 1 до 10 арабскими или римским (от I до X ). С нулем, отрицательными, 11 и более значениями приложение работать НЕ будет! " + BasicErrorMsg)
			break
		}

		if (checkOperandsRome(str1) && checkOperandsArab(str3)) || (checkOperandsRome(str3) && checkOperandsArab(str1)) {
			errorString = ("Учтите, У Вас разносортица в операндах. Вводите только римские или только арабские цифры. ") //+ BasicErrorMsg)
			break
		}
		/*rome routines*/
		if checkOperandsRome(str1) {
			op1 = RomeMap[str1]
			op2 = RomeMap[str3]
			switch str2 {
			case "+":
				res = op1 + op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			case "-":
				res = op1 - op2
			}
			if res < 0 {
				errorString = ("Извините, когда работаете с римскими цифрами, подгадвайте операции так, чтобы результат был больше 0, чтобы получать отрицательные значения задавайте вычисление арабскими цифрами") //+ BasicErrorMsg)
				break
			}

			fmt.Println(intToRome100(res)) //rome result

		} else { //arab tests already done
			op1, _ := strconv.Atoi(str1)
			op2, _ := strconv.Atoi(str3)
			switch str2 {
			case "+":
				res = op1 + op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			case "-":
				res = op1 - op2
			}
			fmt.Println(res) //arab result
		} //arabs

		str1 = ""
		str2 = ""
		str3 = ""
	} //for
	fmt.Println(errorString) //error result prints, then exits programm
} //main
