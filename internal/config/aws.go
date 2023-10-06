package config

type AWSConfig struct {
	Lambda          string `env:"LAMBDA_RUNTIME_API"`
	AccessKeyId     string `env:"ACCESSKEYID"`
	SecretAccessKey string `env:"SECRETACCESSKEY"`
	DefaultRegion   string `env:"DEFAULTREGION"`
	Bucket          string `env:"BUCKET"`
}
