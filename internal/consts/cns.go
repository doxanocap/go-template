package consts

import "time"

const (
	AppDevelopmentEnv = "development"
	AppProductionEnv  = "production"
)

const (
	NilString = ""
	Space     = ""
	NewLine   = "\n"
)

var (
	ByteNewLine = []byte{'\n'}
	ByteSpace   = []byte{' '}
)

const (
	PictureFormatPNG  = "png"
	PictureFormatJPEG = "jpeg"
	PictureFormatIMG  = "img"
	PictureFormatJPG  = "jpg"
)

const (
	TokenMaxAge = 30 * 24 * 60 * 60 * 1000
	TokenPath   = "/"
)

// Postgres
const (
	StorageTable   = "storage"
	StorageTableID = "id"
	KeyColumn      = "key"
	FormatColumn   = "format"

)

// Mailer
const (
	MailingAddress = "mailing_address"
	MailSentTo     = "mail_sent_to"
)

// Storage
const (
	StorageFileName = "storage_filename"
	StorageFileSize = "storage_filesize"
)

// Websocket
const (
	WebsocketWriteWait      = 10 * time.Second
	WebsocketPongWait       = 60 * time.Second
	WebsocketPingPeriod     = (WebsocketPongWait * 9) / 10
	WebsocketMaxMessageSize = 2056
)

var ()
