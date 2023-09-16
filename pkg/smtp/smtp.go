package smtp

import (
	"github.com/spf13/viper"
	"net/smtp"
)

type SMTP struct {
	Auth    *smtp.Auth
	Address string
	From    string
}

func InitConnection() *SMTP {
	from := viper.GetString("SMTP_EMAIL")
	password := viper.GetString("SMTP_PASSWORD")

	host := viper.GetString("SMTP_HOST")
	port := viper.GetString("SMTP_PORT")
	address := host + ":" + port

	auth := smtp.PlainAuth("", from, password, host)
	return &SMTP{
		Auth:    &auth,
		Address: address,
		From:    from,
	}
}

func (s *SMTP) Send(address string, to []string, message []byte) error {
	return smtp.SendMail(address, *s.Auth, s.From, to, message)
}
