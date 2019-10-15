package config

// Config 全局配置示例
var Config config

// config 全部配置
type config struct {
	// 数据库配置
	DB db `mapstructure:"db"`

	// 应用配置功能
	App app `mapstructure:"app"`

	// 监听配置
	Listen listen `mapstructure:"listen"`

	//并发配置
	Goroutine goroutine_config `mapstructure:"goroutine"`

	LogConfig log_config `mapstructure:"log_config"`
}

type db struct {
	// 连接字符串
	Dsn string `mapstructure:"dsn"`

	// 密码
	Pwd string `mapstructure:"pwd"`

	// 最大空闲连接数
	MaxIdleConn int `mapstructure:"maxidle_conn"`

	// 最大连接数
	MaxOpenConn int `mapstructure:"maxopen_conn"`
}

// 监听功能
type listen struct {
	Domain  string `mapstructure:"domain"`
	Address string `mapstructure:"address"`
}

// 应用配置功能
type app struct {
	Id     int    `mapstructure:"id"`
	Secret string `mapstructure:"secret"`
}

type goroutine_config struct {
	ConcurrentNum int `mapstructure:"concurrent_num"`
	RetryCount    int `mapstructure:"retry_count"`
	SleepTime     int `mapstructure:"sleep_time"`
}

type log_config struct {
	Encoding    string `mapstructure:"encoding"`
	LogLevel    string `mapstructure:"log_level"`
	ServiceName string `mapstructure:"srvice_name"`
	StdoutPath  string `mapstructure:"stdout_path"`
	StderrPath  string `mapstructure:"stderr_path"`
}
