package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
	//"strings"
)



func checkOpS(testStrKey string) bool {


	type void struct {} //  sets helpers
	var voidEl void		//	sets helpers
	OpSset := make(map([string]void)

	OpSset["+"] = voidEl
	OpSset["-"] = voidEl
	OpSset["/"] = voidEl
	OpSset["*"] = voidEl

	if _,ok := OpSset[testStrKey]; ok {
		return true
	} else {
		return false
	}
}


func main() {
	//reader := bufio.NewReader(os.Stdin)





	}

	var str1, str2, str3 string
	for {
		fmt.Println("Будет выведен результат, если введете \" два операнда и один оператор между ними (через пробелы) \"")
		fmt.Print("Enter text: ")
		//text, _ := reader.ReadString('\n')
		fmt.Scanln(&str1, &str2, &str3)
		if str1 != "" && str2 != "" && str3 != ""  {

			fmt.Println(str1, str2, str3)

		} else {
			fmt.Println("Ваша строка не соотвествует формату: \" два операнда и один оператор между ними (через пробелы) \"")
			break
		} //if

		str1 = ""
		str2 = ""
		str3 = ""
	} //for
} //main
