package config

type APP struct {
	NodeID   int64  `json:"node_id" mapstructure:"node_id"`
	BindPort string `json:"bind_port" mapstructure:"bind_port"`
	RSAPath  string `json:"rsa_path" mapstructure:"rsa_path"`
}
