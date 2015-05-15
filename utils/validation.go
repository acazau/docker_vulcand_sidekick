package utils

import (
	"regexp"
	"strings"
)

func IsValidIP4(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")
	if res, _ := regexp.Compile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); res.MatchString(ipAddress) {
		return true
	}
	return false
}
