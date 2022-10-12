package util

import "github.com/tyrm/godent/internal/re"

func IsValidMatrixServerName(name string) bool {
	host, _, err := ParseServerName(name)
	if err != nil {
		return false
	}

	return re.IsIPv4Address(host) || re.IsIPv6AddressLiteral(host) || re.IsHostname(host)
}
