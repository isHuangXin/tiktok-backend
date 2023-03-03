package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	initialization "github.com/isHuangxin/tiktok-backend/init"
	"io"
	"sync"
)

var (
	bucket     *oss.Bucket = nil
	bucketOnce sync.Once
)

func initBucket() {
	bucketOnce.Do(func() {
		bucket = initialization.GetBucket()
	})
}

func UploadFromFile(ossPath, localFilePath string) error {
	initBucket()
	return bucket.PutObjectFromFile(ossPath, localFilePath)
}

func UploadFromReader(ossPath string, srcReader io.Reader) error {
	initBucket()
	return bucket.PutObject(ossPath, srcReader)
}
