package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root123456@tcp(zsq-mysql:3306)/blog_zsq?charset=utf8mb4&parseTime=True&loc=Local",
	},
	Redis: RedisConfig{
		Addr: "redis-master:6379",
	},
}
