package re

import "regexp"

var IPv4Address = regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)

func IsIPv4Address(s string) bool {
	return IPv4Address.MatchString(s)
}
