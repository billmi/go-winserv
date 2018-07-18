package check

import "regexp"

func CheckIsMobile(mobileNum string) bool {
	var  regular = "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}