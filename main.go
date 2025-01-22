package main

import (
	"fmt"

	"golang-study/banking"
)

func main() {
	account := banking.Account{Owner: "nico", Balance: 1000}
	fmt.Println(account)
}
