package output

import (
	"strings"
)

func Checkflags(slice []string) bool {
	if len(slice) == 1 {
		if strings.HasPrefix(slice[0], "--output") {
			return false
		}
	}
	if len(slice) > 3 || len(slice) == 0 {
		return false
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "standard.txt" || slice[i] == "shadow.txt" || slice[i] == "thinkertoy.txt" {
			return false
		}
		if strings.HasPrefix(slice[i], "--output=") && i != 0 {
			return false
		}
	}

	temp := strings.Split(slice[0], "=")
	if temp[1] == "" || len(temp[1]) <= 4 || !strings.HasSuffix(temp[1], ".txt") {
		return false
	}

	return true
}
