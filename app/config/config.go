package config

import "github.com/spf13/viper"

type Config struct {
	AppPort    string `mapstructure:"APP_PORT"`
	DBUrl      string `mapstructure:"DB_URL"`
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBDatabase string `mapstructure:"DB_DATABASE"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	JWTKey     string `mapstructure:"JWT_KEY"`
}

func Load() (Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return config, err
	}

	return config, nil
}
