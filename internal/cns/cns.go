package cns

const (
	AppDevelopmentEnv = "development"
	AppProductionEnv  = "production"
)

const (
	NilString = ""
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
	UserTable         = "user"
	UserTableID       = "id"
	UserTableUUID     = "uuid"
	EmailColumn       = "email"
	UsernameColumn    = "username"
	PhoneNumberColumn = "phone_number"
	PasswordColumn    = "password"

	StorageTable   = "storage"
	StorageTableID = "id"
	KeyColumn      = "key"
	FormatColumn   = "format"

	UserParamsTable    = "user_params"
	UserParamsTableID  = "token_id"
	RefreshTokenColumn = "refresh_token"
	UpdatedAtColumn    = "updated_at"
	CreatedAtColumn    = "create_at"
)
