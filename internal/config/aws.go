package config

type AWSConfig struct {
	Lambda          string `env:"LAMBDA_RUNTIME_API"`
	AccessKeyId     string `env:"ACCESS_KEY_ID"`
	SecretAccessKey string `env:"SECRET_ACCESS_KEY"`
	DefaultRegion   string `env:"DEFAULT_REGION"`
	Bucket          string `env:"BUCKET"`
}
