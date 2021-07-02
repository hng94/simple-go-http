package config

type Config struct {
	Filename string
	Address  string
	Route    string
}

var GlobalConfig Config

func init() {
	GlobalConfig = Config{
		Filename: "timestamps.log",
		Address:  ":8080",
		Route:    "/",
	}
}
