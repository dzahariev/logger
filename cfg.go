package logger

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	configFileName   = ".env.yaml"
	configFileFormat = "yaml"
)

const (
	keyLogFormat = "log.format"
	keyLogLevel  = "log.level"
)

// Config configuration fields
type Config struct {
	Format string `json:"log.format"`
	Level  string `json:"log.level"`
}

// Default configurtion
func (l *Config) Default() {
	l.Format = StringFormatter.Name
	l.Level = Error.Name
}

// InitConfig the values from Config
func (l *Config) InitConfig() error {
	l.Default()
	v, err := GetViper()
	if err != nil {
		return fmt.Errorf("failed to initialise viper: %s", err)
	}
	l.LoadValues(v)

	err = l.Validate()
	if err != nil {
		return fmt.Errorf("cannot load the configuration: %s", err)
	}
	return nil
}

// Validate validates the structure integrity
func (l *Config) Validate() error {
	if _, ok := AllFormatters()[l.Format]; !ok {
		return fmt.Errorf("format value should be one of %v", AllFormatters())
	}

	if _, ok := AllLevels()[l.Level]; !ok {
		return fmt.Errorf("level value should be one of %v", AllLevels())
	}

	return nil
}

// LoadValues from viper instance
func (l *Config) LoadValues(v *viper.Viper) {
	if v.IsSet(keyLogFormat) {
		l.Format = v.GetString(keyLogFormat)
	}
	if v.IsSet(keyLogLevel) {
		l.Level = v.GetString(keyLogLevel)
	}
}

// GetViper return the configured Viper instance
func GetViper() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetConfigType(configFileFormat)
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("cannot read cfg file: %s ", err)
	}
	return v, nil
}
