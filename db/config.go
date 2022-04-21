package db

type Configuration struct {
	USERNAME string
	PASSWORD string
	HOST     string
	PORT     string
	DB_NAME  string
}

func InitConfiguration(username string, password string, host string, port string, dbname string) Configuration {
	configuration := Configuration{}
	configuration.USERNAME = username
	configuration.PASSWORD = password
	configuration.HOST = host
	configuration.PORT = port
	configuration.DB_NAME = dbname

	return configuration
}
