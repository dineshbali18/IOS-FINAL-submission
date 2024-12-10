package config

import "github.com/spf13/viper"

// InitializeConfig initializes the config
func InitializeConfig() {
	// Set configuration file which will be used to get/set config values
	viper.SetConfigFile(`../config.yml`)
	// Ask viper to overwrite any configuration values with their corresponding environment counterparts
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
