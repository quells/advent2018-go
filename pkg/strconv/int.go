package strconv

import (
	"errors"
	"strconv"
)

func ParseSignedInt(s string) (i int, err error) {
	if s == "" {
		err = errors.New("empty string")
		return
	}

	sign := 1
	if s[0] == '+' {
		s = s[1:]
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	i, err = strconv.Atoi(s)
	if err != nil {
		return
	}
	i *= sign

	return
}

func MustParseSignedInt(s string) int {
	i, err := ParseSignedInt(s)
	if err != nil {
		panic(err)
	}
	return i
}
