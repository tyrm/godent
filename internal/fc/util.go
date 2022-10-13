package fc

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	gdhttp "github.com/tyrm/godent/internal/http"
)

func cachePeriodFromHeaders(r *http.Response) (int, error) {
	l := logger.WithField("func", "cachePeriodFromHeaders")

	cc := parseCacheControl(r.Header)

	if _, ok := cc["no-store"]; ok {
		l.Debug("found 'no-store' returning 0")

		return 0, nil
	}

	if maxAge, ok := cc["max-age"]; ok {
		maxAgeInt, err := strconv.ParseInt(maxAge, 10, 64)
		if err == nil {
			l.Debugf("found 'max-age' returning %d", maxAgeInt)

			return int(maxAgeInt), nil
		}
	}

	if expires := r.Header.Get(gdhttp.HeaderExpires); expires != "" {
		t, err := time.Parse(time.RFC1123, expires)
		if err == nil {
			cachePeriod := int(time.Until(t).Seconds())
			l.Debugf("found 'expires' returning %d", cachePeriod)

			return cachePeriod, nil
		}
	}
	l.Debug("no cache period found")

	return 0, ErrNotFound
}

func parseCacheControl(r http.Header) map[string]string {
	cacheControl := map[string]string{}

	headers := r.Values(gdhttp.HeaderCacheControl)
	for _, header := range headers {
		for _, h := range strings.Split(header, ",") {
			split := strings.Split(strings.TrimSpace(h), "=")

			value := ""
			if len(split) > 1 {
				value = split[1]
			}

			cacheControl[strings.ToLower(split[0])] = value
		}
	}

	return cacheControl
}
