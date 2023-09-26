package config

type Config struct {
	DataBaseUser     string `json:"data_base_user"`
	DataBasePassword string `json:"data_base_password"`
	DataBaseName     string `json:"data_base_name"`
}

var Cfg Config

func InitializeConfig() {
	Cfg.DataBaseUser = "postgres"
	Cfg.DataBaseName = "postgres"
	Cfg.DataBasePassword = "13134777"
}
