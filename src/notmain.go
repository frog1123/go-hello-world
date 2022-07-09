package main

import (
	"fmt"

	"github.com/fatih/color"
)

func sayHello() {
	const text1 string = "hello from "
	const text2 string = "go"

	fmt.Print(text1)
	color.Cyan("go")
}

func logArray() {
	var arr = []string{"frog", "toad", "dave", "billy"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}