package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

// DBConfig - Database configuration
type DBConfig struct {
	Type             string
	Host             string
	ReadHost         string
	WriteHost        string
	Port             string
	Username         string
	DatabaseName     string
	Password         string
	DatabaseURL      string
	DatabaseReadURL  string
	DatabaseWriteURL string
}

// DatabaseConfig
var DatabaseConfig DBConfig

// GetDatabaseConfig
func GetDatabaseConfig() (err error) {
	//Load the database configuration in the database struct

	DB_CREDENTIALS := viper.GetString("DB_MYSQL_SECRETS")
	if len(DB_CREDENTIALS) > 0 {
		if err = json.Unmarshal([]byte(DB_CREDENTIALS), &DatabaseConfig); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(DatabaseConfig)

	} else {
		DatabaseConfig.Type = viper.GetString("DB_TYPE")
		DatabaseConfig.Username = viper.GetString("DB_USERNAME")
		DatabaseConfig.Password = viper.GetString("DB_PASSWORD")
		DatabaseConfig.DatabaseName = viper.GetString("DB_NAME")
		DatabaseConfig.Host = viper.GetString("DB_HOST")
		DatabaseConfig.WriteHost = viper.GetString("DB_HOST_WRITE")
		DatabaseConfig.ReadHost = viper.GetString("DB_HOST_READ")
		DatabaseConfig.Port = viper.GetString("PORT")
	}

	DatabaseConfig.DatabaseURL, err = generatedDatabaseURL(DatabaseConfig.Host)
	if err != nil {
		return err
	}

	DatabaseConfig.DatabaseReadURL, err = generatedDatabaseURL(DatabaseConfig.Host)
	if err != nil {
		return err
	}

	DatabaseConfig.DatabaseWriteURL, err = generatedDatabaseURL(DatabaseConfig.Host)
	return err
}

// GeneratedDatabaseURL is the function that generated db url
func generatedDatabaseURL(host string) (databaseURL string, err error) {

	if DatabaseConfig.Type == "mysql" {
		databaseURL = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			DatabaseConfig.Username,
			DatabaseConfig.Password,
			host,
			DatabaseConfig.Port,
			DatabaseConfig.DatabaseName,
		)
		return
	}
	err = fmt.Errorf("invalid database type: %s", DatabaseConfig.Type)
	return
}
