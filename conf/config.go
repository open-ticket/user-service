package conf

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

type Config struct {
	Port int64
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	Database string
	SSLMode  string
}

const tagPrefix = "viper"

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("USER")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// set defaults.
	viper.SetDefault("Port", 3000)
	viper.SetDefault("DB.Host", "localhost")
	viper.SetDefault("DB.Port", 5432)
	viper.SetDefault("DB.User", "postgres")
	viper.SetDefault("DB.Password", "password")
	viper.SetDefault("DB.Database", "openticket_user")
	viper.SetDefault("DB.SSLMode", "disable")

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/.openticket/user_service")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return populateConfig(new(Config))
}

func populateConfig(config *Config) (*Config, error) {
	err := recursivelySet(reflect.ValueOf(config), "")
	if err != nil {
		return nil, err
	}

	return config, nil
}

func recursivelySet(val reflect.Value, prefix string) error {
	if val.Kind() != reflect.Ptr {
		return errors.New("WTF")
	}

	// dereference
	val = reflect.Indirect(val)
	if val.Kind() != reflect.Struct {
		return errors.New("FML")
	}

	// grab the type for this instance
	vType := reflect.TypeOf(val.Interface())

	// go through child fields
	for i := 0; i < val.NumField(); i++ {
		thisField := val.Field(i)
		thisType := vType.Field(i)
		tag := prefix + getTag(thisType)

		switch thisField.Kind() {
		case reflect.Struct:
			if err := recursivelySet(thisField.Addr(), tag+"."); err != nil {
				return err
			}
		case reflect.Int:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			// you can only set with an int64 -> int
			configVal := int64(viper.GetInt(tag))
			thisField.SetInt(configVal)
		case reflect.String:
			thisField.SetString(viper.GetString(tag))
		case reflect.Bool:
			thisField.SetBool(viper.GetBool(tag))
		default:
			return fmt.Errorf("unexpected type detected ~ aborting: %s", thisField.Kind())
		}
	}

	return nil
}

func getTag(field reflect.StructField) string {
	// check if maybe we have a special magic tag
	tag := field.Tag
	if tag != "" {
		for _, prefix := range []string{tagPrefix, "mapstructure", "json"} {
			if v := tag.Get(prefix); v != "" {
				return v
			}
		}
	}

	return field.Name
}
