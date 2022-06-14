package models

import "github.com/tyrm/godent/internal/log"

type empty struct{}

var logger = log.WithPackageField(empty{})
