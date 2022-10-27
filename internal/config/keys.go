package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName    string
	ApplicationWebsite string
	PrivacyURLs        string
	PrivacyVersion     string
	RequireTermsAgreed string
	SoftwareVersion    string
	SigningKey         string
	TermsMasterVersion string
	TermsURLs          string
	TermsVersion       string

	// cache
	CacheStore string

	// database
	DBAddress   string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBDatabase  string
	DBTLSMode   string
	DBTLSCACert string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// server
	ExternalHostname string
	ServerHTTPBind   string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ApplicationName:    "application-name",
	ApplicationWebsite: "application-website",
	PrivacyURLs:        "privacy-urls",
	PrivacyVersion:     "privacy-version",
	RequireTermsAgreed: "require-terms-agreed",
	SigningKey:         "signing-key",
	SoftwareVersion:    "software-version", // Set at build
	TermsMasterVersion: "terms-master-version",
	TermsURLs:          "terms-urls",
	TermsVersion:       "terms-version",

	// cache
	CacheStore: "cache-store",

	// database
	DBAddress:   "db-address",
	DBPort:      "db-port",
	DBUser:      "db-user",
	DBPassword:  "db-password",
	DBDatabase:  "db-database",
	DBTLSMode:   "db-tls-mode",
	DBTLSCACert: "db-tls-ca-cert",

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// server
	ExternalHostname: "external-hostname",
	ServerHTTPBind:   "http-bind",
}
