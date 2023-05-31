package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("Nothing to recover.")
			return
		}

		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()

	// panic(errors.New("this is an error"))

	// panic(123)

	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
