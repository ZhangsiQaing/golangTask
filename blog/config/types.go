package config

type config struct {
	DB    DBConfig
	Redis RedisConfig
	Log   LogConfig
}

type DBConfig struct {
	DSN string
}

type RedisConfig struct {
	Addr string
}

type LogConfig struct {
	Level      string `yaml:"level"`       // 日志级别: debug, info, warn, error
	Filename   string `yaml:"filename"`    // 日志文件路径
	MaxSize    int    `yaml:"maxSize"`     // 单个文件最大大小(MB)
	MaxBackups int    `yaml:"maxBackups"`  // 最大备份文件数
	MaxAge     int    `yaml:"maxAge"`      // 最大保留天数
	Compress   bool   `yaml:"compress"`    // 是否压缩
	Console    bool   `yaml:"console"`     // 是否同时输出到控制台
}
