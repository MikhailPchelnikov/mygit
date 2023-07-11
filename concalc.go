package main

import (
	"fmt"
	"strconv"
	"strings"
)

//при допущении что пробелов не будет эта функция оказалась не нужна
//вместо нее действует перебор из 4х попыток поделить ввод пополам по арифметическому
//знаким знакам в качестве разделителя
/*
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
*/
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

const BasicErrorMsg = "Ваша строка не соотвествует формату: \" два операнда и один оператор между ними (с пробелами или без) \""

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

	var Str1, Str2, Str3, Str4, ErrorString, InputString string
	var Op1, Op2, res int
	var OperandError bool

	//проверку на наличие более четырех слов ввода можно оставить, хуже не будет
	Str4 = "cheatcode543210String" //uniq String life user without special knowledge never enters
	ErrorString = BasicErrorMsg

	for {

		fmt.Println("Будет выведен результат, если введете \" два операнда и один оператор между ними (через пробелы) \"")
		fmt.Scanln(&Str1, &Str2, &Str3, &Str4)

		//в алгоритме допускающим ввод без пробелов эта обработка и такое сообщение об ошибке ввода -- не к месту
		/*
			if Str3 == "" {
				errorString = ("Учтите, Вы ввели слишком мало. или забыли про пробелы " + BasicErrorMsg)
				break
			}
		*/

		if Str4 != "cheatcode543210String" {
			ErrorString = ("Учтите, Вы ввели слишком много. " + BasicErrorMsg)
			break
		}

		//в алгоритме обязывающем ставить пробелы была именно такая проверка на наличее арифметического знака
		/*
			if !checkOpS(Str2) {
				errorString = ("Учтите, Вы НЕ ввели арифметический знак из ряда + - * / . между операндами. И может еще что-то не так ввели " + BasicErrorMsg)
				break
			}
		*/
		InputString += Str1 //если пользователь таки использовал пробелы сначала собираем все в одну строку
		InputString += Str2 //но разобьем ее снова на три два операнда и знак в следующем блоке
		InputString += Str3 //который совмещен с проверкой на наличие арифметического знака
		//что позволит использовать уже отработанный алгоритм отлично работающий с тремя строковыми значениями

		for { //поиск арифметического знака и разбиение оставшегося текстого материала по этому знаку ТОЛЬКО на две части: до и после знака
			Str2 = "+"
			Str1, Str3, OperandError = strings.Cut(InputString, Str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один плюс выходим из вложенного for
			Str2 = "-"
			Str1, Str3, OperandError = strings.Cut(InputString, Str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один минус
			Str2 = "*"
			Str1, Str3, OperandError = strings.Cut(InputString, Str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один знак  *

			Str2 = "/"
			Str1, Str3, OperandError = strings.Cut(InputString, Str2)
			if OperandError {
				break
			} //если во введенном пользователе массиве текста есть хотя бы один знак /
			if !OperandError {
				break // если не нашлось ни одного валидного арифметического знака выходим для начала из этого for;;
			}
		}

		if !OperandError {
			ErrorString = ("Учтите, Вы НЕ ввели арифметический знак из ряда + - * / . между операндами. И может еще что-то не так ввели " + BasicErrorMsg)
			break // это выход уже из глобального for
		}

		Str1 = strings.ToUpper(Str1) //чтобы не было проблем с маленькими римскими
		Str3 = strings.ToUpper(Str3)

		if !(checkOperandsArab(Str1) || checkOperandsRome(Str1)) {
			ErrorString = ("Ваш как минимум первый операнд не вписывается в рабочий диапазон, вводите от 1 до 10 арабскими или римским (от I до X ). С нулем, отрицательными, 11 и более значениями приложение работать НЕ будет! " + BasicErrorMsg)
			break
		}

		if !(checkOperandsArab(Str3) || checkOperandsRome(Str3)) {
			ErrorString = ("Ваш второй операнд не вписывается в рабочий диапазон, вводите от 1 до 10 арабскими или римским (от I до X ). С нулем, отрицательными, 11 и более значениями приложение работать НЕ будет! " + BasicErrorMsg)
			break
		}

		if (checkOperandsRome(Str1) && checkOperandsArab(Str3)) || (checkOperandsRome(Str3) && checkOperandsArab(Str1)) {
			ErrorString = ("Учтите, У Вас разносортица в операндах. Вводите только римские или только арабские цифры. ") //+ BasicErrorMsg)
			break
		}
		/*rome routines*/
		if checkOperandsRome(Str1) {
			Op1 = RomeMap[Str1]
			Op2 = RomeMap[Str3]
			switch Str2 {
			case "+":
				res = Op1 + Op2
			case "*":
				res = Op1 * Op2
			case "/":
				res = Op1 / Op2
			case "-":
				res = Op1 - Op2
			}
			if res < 0 {
				ErrorString = ("Извините, когда работаете с римскими цифрами, подгадывайте операции так, чтобы результат был больше 0, чтобы получать отрицательные значения задавайте вычисление арабскими цифрами") //+ BasicErrorMsg)
				break
			}

			fmt.Println(intToRome100(res)) //rome result

		} else { //arab tests already done
			Op1, _ := strconv.Atoi(Str1)
			Op2, _ := strconv.Atoi(Str3)
			switch Str2 {
			case "+":
				res = Op1 + Op2
			case "*":
				res = Op1 * Op2
			case "/":
				res = Op1 / Op2
			case "-":
				res = Op1 - Op2
			}
			fmt.Println(res) //arab result
		} //arab

		Str1 = ""
		Str2 = ""
		Str3 = ""
	} //for
	fmt.Println(ErrorString) //error result prints, then exits program
} //main
