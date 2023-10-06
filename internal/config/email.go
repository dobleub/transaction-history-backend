package config

type EmailConfig struct {
	IAMUsername  string `env:"IAM_USERNAME"`
	SMTPUsername string `env:"SMTP_USERNAME"`
	SMTPPassword string `env:"SMTP_PASSWORD"`
	SMTPHost     string `env:"SMTP_HOST"`
	SMTPPort     int    `env:"SMTP_PORT"`
	STMPSecure   bool   `env:"SMTP_SECURE"`
}
