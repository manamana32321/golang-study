package main

import "fmt"

func is반오십(age int) (result bool) {
	defer fmt.Println("is반오십 함수가 종료되었습니다.")
	switch koreanAge := age + 1; koreanAge {
	case 25:
		result = true
	}
	return result
}

func main() {
	fmt.Println(is반오십(24))
}
