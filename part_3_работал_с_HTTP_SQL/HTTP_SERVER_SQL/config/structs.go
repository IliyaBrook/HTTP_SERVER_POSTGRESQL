package config

type DbConfig struct {
	DBUser     string `mapstructure:"DB_DBUSER"`
	DBPassword string `mapstructure:"DB_DBPASSWORD"`
}

type AppConfig struct {
	Mode           string `mapstructure:"APP_MODE"`
	TrustedProxies string `mapstructure:"APP_TRUSTED_PROXIES"`
}
