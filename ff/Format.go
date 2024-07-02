package output

import (
	"log"
	"os"
	"strings"
)

func Format(banner, text string) ([]string, []string) {
	var sep string
	InputFile := banner + ".txt"
	data, err := os.ReadFile(InputFile)
	if err != nil {
		log.Fatalln("Error with reading ", err)
	}
	sep = strings.ReplaceAll(string(data), "\r", "\n")
	slice := RemoveEmptyStrings(strings.Split(sep, "\n"))
	sliceArgs := strings.Split(text, "\\n")
	return slice, sliceArgs
}

func RemoveEmptyStrings(slice []string) []string {
	var temp []string
	for _, text := range slice {
		if text != "" {
			temp = append(temp, text)
		}
	}
	return temp
}
