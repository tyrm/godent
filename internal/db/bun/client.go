package bun

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/db"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bunotel"
	"os"
	"runtime"
)

const (
	dbTypePostgres = "postgres"
	dbTypeSqlite   = "sqlite"

	dbTLSModeDisable = "disable"
	dbTLSModeEnable  = "enable"
	dbTLSModeRequire = "require"
	dbTLSModeUnset   = ""
)

// New creates a new bun database client.
func New(ctx context.Context) (*Client, error) {
	l := logger.WithField("func", "pgConn")

	opts, err := pgOptions()
	if err != nil {
		return nil, fmt.Errorf("could not doCreate bundb postgres options: %s", err)
	}

	sqldb := stdlib.OpenDB(*opts)

	setConnectionValues(sqldb)

	conn := bun.NewDB(sqldb, pgdialect.New())
	conn.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(viper.GetString(config.Keys.DBDatabase))))

	// ping to check the bun is there and listening
	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("postgres ping: %s", err)
	}
	l.Info("connected to POSTGRES database")

	return &Client{
		db: conn,
	}, nil
}

// Client is a DB interface compatible client for Bun.
type Client struct {
	db *bun.DB
}

var _ db.DB = (*Client)(nil)

func pgOptions() (*pgx.ConnConfig, error) {
	keys := config.Keys

	// these are all optional, the bun adapter figures out defaults
	port := viper.GetInt(keys.DBPort)
	address := viper.GetString(keys.DBAddress)
	username := viper.GetString(keys.DBUser)
	password := viper.GetString(keys.DBPassword)

	// validate database
	database := viper.GetString(keys.DBDatabase)
	if database == "" {
		return nil, errors.New("no database set")
	}

	var tlsConfig *tls.Config
	tlsMode := viper.GetString(keys.DBTLSMode)
	switch tlsMode {
	case dbTLSModeDisable, dbTLSModeUnset:
		break // nothing to do
	case dbTLSModeEnable:
		/* #nosec G402 */
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	case dbTLSModeRequire:
		tlsConfig = &tls.Config{
			InsecureSkipVerify: false,
			ServerName:         viper.GetString(keys.DBAddress),
			MinVersion:         tls.VersionTLS12,
		}
	}

	caCertPath := viper.GetString(keys.DBTLSCACert)
	if tlsConfig != nil && caCertPath != "" {
		// load the system cert pool first -- we'll append the given CA cert to this
		certPool, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("error fetching system CA cert pool: %s", err)
		}

		// open the file itself and make sure there's something in it
		/* #nosec G304 */
		caCertBytes, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, fmt.Errorf("error opening CA certificate at %s: %s", caCertPath, err)
		}
		if len(caCertBytes) == 0 {
			return nil, fmt.Errorf("ca cert at %s was empty", caCertPath)
		}

		// make sure we have a PEM block
		caPem, _ := pem.Decode(caCertBytes)
		if caPem == nil {
			return nil, fmt.Errorf("could not parse cert at %s into PEM", caCertPath)
		}

		// parse the PEM block into the certificate
		caCert, err := x509.ParseCertificate(caPem.Bytes)
		if err != nil {
			return nil, fmt.Errorf("could not parse cert at %s into x509 certificate: %s", caCertPath, err)
		}

		// we're happy, add it to the existing pool and then use this pool in our tls config
		certPool.AddCert(caCert)
		tlsConfig.RootCAs = certPool
	}

	cfg, _ := pgx.ParseConfig("")
	if address != "" {
		cfg.Host = address
	}
	if port > 0 {
		cfg.Port = uint16(port)
	}

	if username != "" {
		cfg.User = username
	}
	if password != "" {
		cfg.Password = password
	}
	if tlsConfig != nil {
		cfg.TLSConfig = tlsConfig
	}
	cfg.Database = database
	cfg.PreferSimpleProtocol = true
	cfg.RuntimeParams["application_name"] = viper.GetString(keys.ApplicationName)

	return cfg, nil
}

// processPostgresError processes an error, replacing any postgres specific errors with our own error type
func processError(err error) db.Error {
	l := logger.WithField("func", "processError")

	switch {
	case err == nil:
		return nil
	case err == sql.ErrNoRows:
		return db.ErrNoEntries
	default:
		// Attempt to cast as postgres
		pgErr, ok := err.(*pgconn.PgError)
		if !ok {
			return err
		}

		l.Debugf("postgres error %s: %s", pgErr.Code, pgErr.Error())

		// Handle supplied error code:
		// (https://www.postgresql.org/docs/10/errcodes-appendix.html)
		switch pgErr.Code {
		case "23505" /* unique_violation */ :
			return db.NewErrAlreadyExists(pgErr.Message)
		default:
			return err
		}
	}
}

// https://bun.uptrace.dev/postgres/running-bun-in-production.html#database-sql
func setConnectionValues(sqldb *sql.DB) {
	maxOpenConns := 4 * runtime.GOMAXPROCS(0)
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)
}
