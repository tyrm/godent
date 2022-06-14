package db

import "context"

// DB represents a database client.
type DB interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// Create stores the object
	Create(ctx context.Context, i interface{}) Error
	// DoMigration runs database migrations
	DoMigration(ctx context.Context) Error
	// LoadTestData adds test data to the database
	LoadTestData(ctx context.Context) Error
	// ReadByID returns a model by its ID
	ReadByID(ctx context.Context, id int64, i interface{}) Error
	// ResetCache clears any caches in the module
	ResetCache(ctx context.Context) Error
	// Update updates stored data
	Update(ctx context.Context, i interface{}) Error
}
