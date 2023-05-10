package config


var cfg *config

type config struct {
	DB  DBConfig
}


type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DataBase string
}

