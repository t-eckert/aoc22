package main

import (
	"fmt"
	"os"
)

func main() {
	raw, err := os.ReadFile("./13/test.input")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(raw))
}
