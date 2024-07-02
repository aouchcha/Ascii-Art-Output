package output

import (
	"fmt"
	"os"
)

func GenerateFile(file, result string) {
	file += ".txt"
	err := os.WriteFile(file, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error hhhhhhhhhh", err)
		return
	}
}
