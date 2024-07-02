package main

import (
	"fmt"
	"os"
	function "output/ff"
)

func main() {
	Args := os.Args[1:]
	if function.Args(Args) {
		if function.Checkflags(Args) {
			AsciiOutput()
		} else {
			fmt.Println("Error de format main line 15")
			return
		}
	} else {
		for i := 0; i < len(Args); i++ {
			if Args[i] == "standard.txt" || Args[i] == "shadow.txt" || Args[i] == "thinkertoy.txt" {
				fmt.Printf("Error %s instead of banner\n", Args[i])
				return
			}
		}
		AsciiFS()
	}
}

func AsciiOutput() {
	file, text, banner := function.Data(os.Args[1:])
	slice, slicedArgs := function.Format(banner, text)
	result := function.AsciiArt(slice, slicedArgs)
	if function.IsItAllNL(result) {
		result = result[1:]
	}
	function.GenerateFile(file, result)
}

func AsciiFS() {
	_, text, banner := function.Data(os.Args[1:])
	slice, slicedArgs := function.Format(banner, text)
	result := function.AsciiArt(slice, slicedArgs)
	if function.IsItAllNL(result) {
		result = result[1:]
	}
	fmt.Print(result)
}
