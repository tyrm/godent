package util

import (
	"errors"
	"strconv"
	"strings"
)

func ParseServerName(name string) (string, int, error) {
	if strings.HasSuffix(name, "]") {
		// probably an IPv6 Address
		return name, 0, nil
	}

	splitName := strings.Split(name, ":")

	hostname := splitName[0]
	port := int64(0)

	if len(splitName) > 1 {
		n, err := strconv.ParseInt(splitName[1], 10, 32)
		if err != nil {
			return "", 0, err
		}

		port = n
	}

	if port != 0 {
		if splitName[1] != strconv.FormatInt(port, 10) || !(1 <= port && port < 65536) {
			return "", 0, errors.New("invalid port")
		}
	}

	return hostname, int(port), nil
}
