package consts

func IsNilString(str string) bool {
	return str == NilString
}

func IsValidFormat(format string) bool {
	return format == PictureFormatIMG ||
		format == PictureFormatPNG ||
		format == PictureFormatJPEG ||
		format == PictureFormatJPG
}
