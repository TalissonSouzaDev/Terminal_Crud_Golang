package config

import "github.com/spf13/viper"

type config struct {
	Dbdriver   string `mapstructure:"DB_DRIVER"`
	Dbhost     string `mapstructure:"DB_HOST"`
	Dbport     string `mapstructure:"DB_PORT"`
	Dbname     string `mapstructure:"DB_NAME"`
	Dbuser     string `mapstructure:"DB_USER"`
	Dbpassword string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig() (*config, error) {
	viper.SetConfigFile("cmd/.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg, nil
}
