package main

import (
	"fmt"
	"os"
)

func main() {
	var chislo uint64
	fmt.Println("Введите целое положительное число")
	fmt.Fscan(os.Stdin, &chislo)
	for i := 1; i <= int(chislo); i++ {
		fmt.Printf(fmt.Sprintf("%d %s", i, " "))
	}
	//var a string = "foo-bar"
	if chislo%3 == 0 {
		fmt.Printf("foo-bar")
	} else if chislo%3 > 0 {
		fmt.Printf("foo")
	}
}
