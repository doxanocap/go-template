package processor

type IMailerProcessor interface {
	Send(address string, to []string, message []byte) error
}

type IMailerProvider interface {
	Send(address string, to []string, message []byte) error
}
