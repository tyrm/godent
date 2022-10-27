package pubkey

import "github.com/tyrm/godent/internal/log"

type empty struct{}

var logger = log.WithPackageField(empty{})
