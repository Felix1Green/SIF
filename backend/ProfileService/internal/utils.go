package internal

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizeInput(sanitizer *bluemonday.Policy, arr ...*string) {
	fmt.Println("Sanitizing input")
	for index, val := range arr {
		if val == nil {
			continue
		}
		*arr[index] = sanitizer.Sanitize(*val)
	}
}
