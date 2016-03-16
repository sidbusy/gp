package gp

import (
	"regexp"
)

var (
	mobileRE = regexp.MustCompile(`^1(3[0-9]|4[57]|5[012356789]|7[01]|8[0-9])\d{8}$`)
)

func IsMobile(num string) bool {
	return mobileRE.MatchString(num)
}
