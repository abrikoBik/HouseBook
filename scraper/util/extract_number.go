package util

import (
	"regexp"
	"strconv"
	"strings"
)

// ExtractNumber extracts all digits from a string using regular expressions and then concatenates them.
func ExtractNumber(str string) (result int, err error) {
	rx, err := regexp.Compile(`\d`)
	if err != nil {
		return 0, err
	}
	number := strings.Join(rx.FindAllString(str, -1), "")
	if number == "" {
		return 0, nil
	}
	result, err = strconv.Atoi(number)
	return
}
