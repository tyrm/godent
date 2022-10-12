package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName    string
	ApplicationWebsite string
	SoftwareVersion    string

	// database
	DBAddress       string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBDatabase      string
	DBTLSMode       string
	DBTLSCACert     string
	DBEncryptionKey string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// server
	ExternalHostname string
	ServerHTTPBind   string

	// matrix
	RequireTermsAgreed string
	PrivacyURLs        string
	PrivacyVersion     string
	TermsURLs          string
	TermsVersion       string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ApplicationName:    "application-name",
	ApplicationWebsite: "application-website",
	SoftwareVersion:    "software-version", // Set at build

	// database
	DBAddress:       "db-address",
	DBPort:          "db-port",
	DBUser:          "db-user",
	DBPassword:      "db-password",
	DBDatabase:      "db-database",
	DBTLSMode:       "db-tls-mode",
	DBTLSCACert:     "db-tls-ca-cert",
	DBEncryptionKey: "db-crypto-key",

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// server
	ExternalHostname: "external-hostname",
	ServerHTTPBind:   "http-bind",

	// matrix
	RequireTermsAgreed: "require-terms-agreed",
	PrivacyURLs:        "privacy-urls",
	PrivacyVersion:     "privacy-version",
	TermsURLs:          "terms-urls",
	TermsVersion:       "terms-version",
}
