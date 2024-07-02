package output

import (
	"strings"
)

func Args(slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if strings.HasPrefix(slice[i], "--output=") {
			return true
		}
	}
	return false
}
