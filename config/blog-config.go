package config

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port string
}

type DatabaseConfiguration struct {
	DBName     string
	Table      string
	DBUser     string
	DBPassword string
	Addr       string
	Network    string
}
