package main

import (
	"fmt"
)

func main() {
	marks := 86
	if marks >= 35 && marks < 55 {
		fmt.Println("grade e")
	} else if marks >= 55 && marks < 70 {
		fmt.Println("grade d")
	} else if marks >= 70 && marks < 80 {
		fmt.Println("grade c")
	} else if marks >= 80 && marks < 90 {
		fmt.Println("grade b")
	} else if marks >= 90 && marks <= 100 {
		fmt.Print("grade a")
	} else {
		fmt.Println("fail")
	}

}
