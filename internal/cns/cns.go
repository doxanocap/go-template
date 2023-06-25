package cns

const (
	NilString = ""
)

func IsNilString(str string) bool {
	return str == NilString
}

const (
	PictureFormatPNG  = "png"
	PictureFormatJPEG = "jpeg"
	PictureFormatIMG  = "img"
	PictureFormatJPG  = "jpg"
)

func IsValidFormat(format string) bool {
	return format == PictureFormatIMG ||
		format == PictureFormatPNG ||
		format == PictureFormatJPEG ||
		format == PictureFormatJPG
}
