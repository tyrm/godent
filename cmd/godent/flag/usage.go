package flag

import "github.com/tyrm/godent/internal/config"

var usage = config.KeyNames{
	ConfigPath: "Path to a file containing feditools configuration. Values set in this file will be overwritten by values set as env vars or arguments",
	LogLevel:   "Log level to run at: [trace, debug, info, warn, fatal]",

	// application
	ApplicationName: "Name of the application, used in various places internally",

	// database
	DBType:         "Database type: eg., postgres",
	DBAddress:      "Database ipv4 address, hostname, or filename",
	DBPort:         "Database port",
	DBUser:         "Database username",
	DBPassword:     "Database password",
	DBDatabase:     "Database name",
	DBTLSMode:      "Database tls mode",
	DBTLSCACert:    "Path to CA cert for db tls connection",
	DBLoadTestData: "Should test data be loaded into the database",
}
