package tools

import (
	"regexp"
)


var NumExp *regexp.Regexp = regexp.MustCompile("(?:^\\d*?)\\d$")

var MinerNameExp *regexp.Regexp = regexp.MustCompile("(\\d|\\w|\\.|-){1,64}")


var EmailExp *regexp.Regexp = regexp.MustCompile("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")


func IsNum(data string) bool {

	return NumExp.MatchString(data)
}

func IsMinerName(data string) bool {

	return MinerNameExp.MatchString(data)
}

func IsEmail(data string) bool {

	return EmailExp.MatchString(data)
}