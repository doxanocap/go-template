package model

import "io"

type Storage struct {
	ID     int64  `json:"id" db:"id"`
	Key    string `json:"key" db:"key"`
	Format string `json:"format" db:"format"`
}

type StoragePictureSave struct {
	BucketName string
}

type HandlePictureReq struct {
	File     io.Reader
	FileName string
	Key      string
	Size     int64
}
