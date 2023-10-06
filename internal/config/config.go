package config

type Config struct {
	Port        string      `env:"PORT,default=8080"`
	AWSConfig   AWSConfig   `env:", prefix=AWS_"`
	EmailConfig EmailConfig `env:", prefix=EMAIL_"`
}
