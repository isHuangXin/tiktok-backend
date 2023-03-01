package oss

import (
	initialization "github.com/isHuangxin/tiktok-backend/init"
	"io"
)

var bucket = initialization.GetBucket()

func UploadFromFile(ossPath, localFilePath string) error {
	return bucket.PutObjectFromFile(ossPath, localFilePath)
}

func UploadFromReader(ossPath string, srcReader io.Reader) error {
	return bucket.PutObject(ossPath, srcReader)
}
