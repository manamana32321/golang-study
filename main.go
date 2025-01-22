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

func showAddress(a int) {
	fmt.Println(&a)
}

func showAddress2(a *int) {
	fmt.Println(a)
}

func showValue(a *int) {
	fmt.Println(*a)
}

func pointer() {
	a := 2
	b := &a
	a = 10
	fmt.Println(a, *b)
	*b = 20
	showValue(b)
}

func main() {
	fmt.Println(is반오십(24))
	pointer()
}
