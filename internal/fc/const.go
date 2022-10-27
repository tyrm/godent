package fc

const (
	minute = 60
	hour   = 3600

	dnsDefaultCachePeriod             = 5 * minute
	homeServerInvalidCachePeriod      = 1 * hour
	wellKnownDefaultCachePeriod       = 24 * hour
	wellKnownDefaultCachePeriodJitter = 10 * minute
	wellKnownDefaultMaxPeriod         = 48 * hour
)
