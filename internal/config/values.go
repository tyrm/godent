package config

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ApplicationName    string
	ApplicationWebsite string
	PrivacyURLs        map[string]interface{}
	PrivacyVersion     string
	RequireTermsAgreed bool
	SigningKey         string
	SoftwareVersion    string
	TermsMasterVersion string
	TermsURLs          map[string]interface{}
	TermsVersion       string

	// database
	DBAddress   string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBDatabase  string
	DBTLSMode   string
	DBTLSCACert string

	// redis
	RedisAddress  string
	RedisDB       int
	RedisPassword string

	// server
	ExternalHostname string
	ServerHTTPBind   string

	// matrix
}

// Defaults contains the default values.
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ApplicationName:    "godent",
	ApplicationWebsite: "https://github.com/tyrm/godent",

	// database
	DBAddress:   "localhost",
	DBPort:      5432,
	DBUser:      "godent",
	DBPassword:  "godent",
	DBDatabase:  "godent",
	DBTLSMode:   "disable",
	DBTLSCACert: "",

	// redis
	RedisAddress:  "localhost:6379",
	RedisDB:       0,
	RedisPassword: "",

	// server
	ServerHTTPBind: ":5000",

	// matrix
	RequireTermsAgreed: true,
}
