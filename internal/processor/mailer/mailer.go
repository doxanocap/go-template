package mailer

import (
	"app/internal/cns"
	IProcessor "app/internal/manager/interfaces/processor"
	"go.uber.org/zap"
)

type Mailer struct {
	provider IProcessor.IMailerProvider
	log      *zap.Logger
}

func Init(provider IProcessor.IMailerProvider, log *zap.Logger) *Mailer {
	return &Mailer{
		provider: provider,
		log:      log,
	}
}

func (m *Mailer) Send(address string, to []string, message []byte) error {
	m.log.Named("Send").With(
		zap.String(cns.MailingAddress, address),
		zap.Strings(cns.MailSentTo, to)).Info("message")

	return m.provider.Send(address, to, message)
}
