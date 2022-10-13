package fc

import "time"

const (
	dnsDefaultCachePeriod             = 1 * time.Minute
	homeServerInvalidCachePeriod      = 1 * time.Hour
	wellKnownDefaultCachePeriod       = 24 * time.Hour
	wellKnownDefaultCachePeriodJitter = 10 * time.Minute
	wellKnownDefaultMaxPeriod         = 48 * time.Hour
)
