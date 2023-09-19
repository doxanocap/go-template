package model

type Config struct {
	LogJSON   bool   `env:"LOG_JSON"`
	RabbitDSN string `env:"RABBIT_DSN"`

	Psql
	App
	Token
	Redis
	AWS
	SMTP
}

type App struct {
	Version     string `env:"APP_VERSION"`
	Environment string `env:"APP_ENVIRONMENT"`

	ServerIP   string `env:"SERVER_IP"`
	ServerPort string `env:"SERVER_PORT"`
}

type Token struct {
	RefreshSecret string `env:"REFRESH_TOKEN_SECRET"`
	AccessSecret  string `env:"ACCESS_TOKEN_SECRET"`
}

type Psql struct {
	PsqlHost     string `env:"PSQL_HOST"`
	PsqlPort     string `env:"PSQL_PORT"`
	PsqlUser     string `env:"PSQL_USER"`
	PsqlPassword string `env:"PSQL_PASSWORD"`
	PsqlDatabase string `env:"PSQL_DATABASE"`
	PsqlSSL      string `env:"PSQL_SSL"`
}

type Redis struct {
	Prefix   string `env:"REDIS_PREFIX"`
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
	Database int    `env:"REDIS_DATABASE"`
}

type AWS struct {
	S3
}

type S3 struct {
	AccessKey   string `env:"S3_ACCESS_KEY"`
	SecretKey   string `env:"S3_SECRET_KEY"`
	EndpointUrl string `env:"S3_ENDPOINT_URL"`
	BucketName  string `env:"S3_BUCKET_NAME"`
	UseSSL      bool   `env:"S3_USE_SSL"`
}

type SMTP struct {
	Email    string `env:"SMTP_EMAIL"`
	Password string `env:"SMTP_PASSWORD"`
	Host     string `env:"SMTP_HOST"`
	Port     string `env:"SMTP_PORT"`
}
