package config

type App struct {
	HttpPort string `json:"http_port" mapstructure:"http_port"`
	RSAPath  string `json:"rsa_path" mapstructure:"rsa_path"`
}
