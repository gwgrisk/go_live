package main

import (
	"bufio"
	"fmt"
	"os"

	"siact.com/learngo/lang/functional/fib"
)

func tryDefer() {
	defer println(1)
	defer println(2)
	println(3)
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	tryDefer()
	writeFile("lang/defer/defer1.txt")
}
