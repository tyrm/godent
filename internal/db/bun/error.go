package bun

import (
	"database/sql"

	"github.com/jackc/pgconn"
	"github.com/tyrm/godent/internal/db"
)

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
