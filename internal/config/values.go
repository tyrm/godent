package config

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ApplicationName    string
	ApplicationWebsite string
	SoftwareVersion    string
	TokenSalt          string

	// database
	DBType          string
	DBAddress       string
	DBPort          int
	DBUser          string
	DBPassword      string
	DBDatabase      string
	DBTLSMode       string
	DBTLSCACert     string
	DBLoadTestData  bool
	DBEncryptionKey string

	// redis
	RedisAddress  string
	RedisDB       int
	RedisPassword string

	// server
	ServerHTTPBind string
}

// Defaults contains the default values.
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ApplicationName:    "godent",
	ApplicationWebsite: "https://github.com/tyrm/godent",

	// database
	DBType:         "postgres",
	DBAddress:      "localhost",
	DBPort:         5432,
	DBUser:         "godent",
	DBPassword:     "godent",
	DBDatabase:     "godent",
	DBTLSMode:      "disable",
	DBTLSCACert:    "",
	DBLoadTestData: false,

	// redis
	RedisAddress:  "localhost:6379",
	RedisDB:       0,
	RedisPassword: "",

	// server
	ServerHTTPBind: ":5000",
}
