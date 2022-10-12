package db

import "context"

// DB represents a database client.
type DB interface {
	AcceptedTermsURL
	Account

	// Close closes the db connections
	Close(ctx context.Context) Error
	// DoMigration runs database migrations
	DoMigration(ctx context.Context) Error
}
