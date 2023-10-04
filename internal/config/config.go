package config

type Config struct {
	Port      string    `env:"PORT,default=8080"`
	AWSConfig AWSConfig `env:", prefix=AWS_"`
}
