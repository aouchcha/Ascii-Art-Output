package output

import (
	"log"
)

func AsciiArt(slice, slicedArgs []string) string {
	var result string

	for _, word := range slicedArgs {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						log.Fatalln("You entered an inprintable caracter !!!")
					}
					start := int(char-32)*8 + i
					result += slice[start]
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}
