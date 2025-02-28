package config

type DBConfig struct {
	Host      string
	User      string
	Password  string
	DBName    string
	Port      string
	Charset   string
	ParseTime string
	Locale    string
}

func DBConfigFromEnv() DBConfig {
	return DBConfig{
		Host:      Env("MYSQL_HOST"),
		User:      Env("MYSQL_USER"),
		Password:  Env("MYSQL_PASSWORD"),
		DBName:    Env("MYSQL_DB"),
		Port:      Env("MYSQL_PORT"),
		Charset:   Env("MYSQL_CHARSET"),
		ParseTime: Env("MYSQL_PARSE_TIME"),
		Locale:    Env("MYSQL_LOCAL"),
	}
}
