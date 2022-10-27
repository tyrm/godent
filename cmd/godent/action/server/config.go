package server

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"reflect"
)

func logConfig(l *logrus.Entry) {
	keyVals := reflect.ValueOf(config.Keys)

	l.Trace("config")
	for i := 0; i < keyVals.NumField(); i++ {
		k := keyVals.Field(i).String()
		l.Tracef("  %s: %v", k, viper.Get(k))
	}
}
