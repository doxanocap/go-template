package mailer

import (
	IProcessor "app/internal/manager/interfaces/processor"
	"fmt"
	"github.com/doxanocap/pkg/errs"

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
	log := m.log.Named("[Send]")

	err := m.provider.Send(address, to, message)
	if err != nil {
		log.Error(fmt.Sprintf("mailer provider: %s", err))
		err = errs.Wrap("processor.mailer.Send", err)
	}
	return err
}
