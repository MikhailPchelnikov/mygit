package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
	//"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	var str1, str2, str3 string
	for {
		fmt.Print("Enter text: ")
		//text, _ := reader.ReadString('\n')
		fmt.Scanln(&str1, &str2, &str3)
		if str1 != "" && str2 != "" && str3 != "" {

			fmt.Println(str1, str2, str3)

		} else {
			fmt.Println("вводите задания в формате: \" два операнда и один оператор между ними (через пробел)\"")
			break
		} //if

		str1 = ""
		str2 = ""
		str3 = ""
	} //for
} //main
