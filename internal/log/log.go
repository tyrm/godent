package log

import (
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
)

// Init the logging engine.
func Init() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logLevel := viper.GetString(config.Keys.LogLevel)

	if logLevel != "" {
		level, err := logrus.ParseLevel(logLevel)
		if err != nil {
			return err
		}
		logrus.SetLevel(level)
	}

	return nil
}

// WithPackageField creates a new logrus entry with the package name added as a field.
func WithPackageField(m interface{}) *logrus.Entry {
	return logrus.WithField("package", strings.ReplaceAll(strings.TrimPrefix(reflect.TypeOf(m).PkgPath(), "github.com/tyrm/godent/"), "/", "."))
}
