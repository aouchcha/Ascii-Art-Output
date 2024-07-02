package output

import (
	"log"
	"os"
	"strings"
)

func Data(slice []string) (string, string, string) {
	var file string
	var text string
	var banner string = "standard"
	count := 0
	banners := []string{"standard", "shadow", "thinkertoy"}
	for i, ban := range banners {
		if slice[len(slice)-1] == ban {
			banner = banners[i]
			count++
		}
	}
	if len(slice) == 1 {
		text = slice[0]
		banner = "standard"
		return file, text, banner
	}
	if len(slice) > 1 {
		if strings.HasPrefix(slice[0], "--output=") {
			temp := strings.Split(slice[0], "=")
			temp2 := strings.Split(temp[1], ".txt")
			if len(temp2) != 2 || temp2[0] == "" {
				log.Fatalln("Error: there if a problem in ", temp[1])
			}
			file = temp2[0]
			if len(slice[1:]) == 2 && count == 0 {
				log.Fatalln("Error de format data.go , FS problem banner not exist")
				os.Exit(0)
			} else if len(slice[1:]) == 1 {
				text = slice[1]
				banner = "standard"
			} else {
				text = slice[1]
			}
		} else {
			if len(slice) > 2 {
				log.Fatalln("Error de format len = 3 and there's no flag data.go line 45")
				os.Exit(0)
			} else if len(slice) == 2 && count == 0 {
				log.Fatalln("Error de format data.go , FS problem banner not exist")
				os.Exit(0)
			}
			text = slice[0]
			file = ""
			return file, text, banner
		}
	}
	return file, text, banner
}
