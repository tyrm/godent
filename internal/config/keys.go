package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName    string
	ApplicationWebsite string
	SoftwareVersion    string
	TokenSalt          string

	// database
	DBType          string
	DBAddress       string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBDatabase      string
	DBTLSMode       string
	DBTLSCACert     string
	DBLoadTestData  string
	DBEncryptionKey string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// server
	ServerHTTPBind string
	ServerRoles    string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ApplicationName:    "application-name",
	ApplicationWebsite: "application-website",
	SoftwareVersion:    "software-version", // Set at build
	TokenSalt:          "token-salt",

	// database
	DBType:          "db-type",
	DBAddress:       "db-address",
	DBPort:          "db-port",
	DBUser:          "db-user",
	DBPassword:      "db-password",
	DBDatabase:      "db-database",
	DBTLSMode:       "db-tls-mode",
	DBTLSCACert:     "db-tls-ca-cert",
	DBLoadTestData:  "test-data", // CLI only
	DBEncryptionKey: "db-crypto-key",

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// server
	ServerHTTPBind: "http-bind",
	ServerRoles:    "server-role",
}
