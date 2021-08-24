package config

type DBConfig struct {
	Type           string `mapstructure:"type"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Host           string `mapstructure:"host" default:"127.0.0.1"`
	Port           string `mapstructure:"port" default:"3306"`
	Database       string `mapstructure:"database"`
	ConnectTimeout int    `mapstructure:"connect_timeout" default:"10"`
	Debug          bool   `mapstructure:"debug" default:"true"`
	MaxLifetime    int    `mapstructure:"max_lifetime" default:"7200"`  //设置连接可以重用的最长时间(单位：秒)
	MaxOpenConns   int    `mapstructure:"max_open_conns" default:"150"` //设置数据库的最大打开连接数
	MaxIdleConns   int    `mapstructure:"max_idle_conns" default:"50"`  //设置空闲连接池中的最大连接数
}
