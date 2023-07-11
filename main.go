package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
	//"strings"
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
	str4 = "cheatcode543210string" //uniq string life user without special knowledge never enters
	errorString = BasicErrorMsg
	for {

		fmt.Println("Будет выведен результат, если введете \" два операнда и один оператор между ними (через пробелы) \"")
		//fmt.Print("Enter text: ")
		//text, _ := reader.ReadString('\n')
		fmt.Scanln(&str1, &str2, &str3, &str4)
		if str3 == "" {
			errorString = ("Учтите, Вы ввели слишком мало. " + BasicErrorMsg)
			break
		}

		if str4 != "cheatcode543210string" {
			errorString = ("Учтите, Вы ввели слишком много. " + BasicErrorMsg)
			break
		}

		if !checkOpS(str2) {
			errorString = ("Учтите, Вы не ввели арифметический знак из ряда + - * / . Или ввели его не так, как надо. " + BasicErrorMsg)
			break
		}

		if (checkOperandsRome(str1) && checkOperandsArab(str3)) || (checkOperandsRome(str3) && checkOperandsArab(str1)) {
			errorString = ("Учтите, У вас разносортица в операндах. Вводите только римские или только арабские цифры. " + BasicErrorMsg)
			break
		}

		if true {
			fmt.Println("is " + str1 + " Arabic number 1 to 10 ?")
			fmt.Println(checkOperandsArab(str1))
			fmt.Println("is " + str1 + " Rome number I to X ?")
			fmt.Println(checkOperandsRome(str1))

			fmt.Println("is " + str3 + " Arabic number 1 to 10 ?")
			fmt.Println(checkOperandsRome(str3))
			fmt.Println("is " + str3 + " Rome number I to X ?")
			fmt.Println(checkOperandsRome(str3))

		} else {
			//fmt.Println(errorstr)
			break
		} //if

		str1 = ""
		str2 = ""
		str3 = ""
		//str4 = "cheatcode543210string" //uniq string life user without special knowledge never enters
	} //for
	fmt.Println(errorString)
} //main
