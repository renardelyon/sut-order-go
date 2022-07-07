package config

type RedisConfig struct {
	Addr     string `mapstructure:"REDIS_ADDRESS"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DBNUMBER"`
}

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBUrl       string `mapstructure:"DB_URL"`
	ProductHost string `mapstructure:"PRODUCT_HOST"`
	StorageHost string `mapstructure:"STORAGE_HOST"`
	RedisConfig RedisConfig
}
