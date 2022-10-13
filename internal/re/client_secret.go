package re

import "regexp"

var ClientSecret = regexp.MustCompile(`^[0-9a-zA-Z.=_\-]+$`)

func IsClientSecret(s string) bool {
	return ClientSecret.MatchString(s)
}
