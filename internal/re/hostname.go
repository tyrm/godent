package re

import "regexp"

var Hostname = regexp.MustCompile(`(?i)^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)(?:\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*$`)

func IsHostname(s string) bool {
	return Hostname.MatchString(s)
}
