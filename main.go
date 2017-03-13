package main

import (
	"./files"
	"fmt"
)


func main() {
	files.ReadFile("./files/text1.txt")
	fmt.Print("\n \n \n")
	files.ReadFile("./files/text2.txt")
}