package main

import (
	"errors"
	"fmt"
)

func main() {
	div, err := division(7, 0)

	if err != nil {

		fmt.Println("error en la division: ", err.Error())
		return
	}
	fmt.Println("Division:", div)
}

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("Division por 0")
	}

	return a / b, nil
}
