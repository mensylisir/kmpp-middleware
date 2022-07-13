package storage

import (
	"errors"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/util/storage/client"
)

type CloudStorageClient interface {
	CreateBuckets() error
	ListBuckets() ([]interface{}, error)
	Exist(path string) (bool, error)
	Delete(path string) (bool, error)
	Upload(src, target string) (bool, error)
	Download(src, target string) (bool, error)
}

func NewCloudStorageClient(vars map[string]interface{}) (CloudStorageClient, error) {
	if vars["type"] == constant.Minio {
		return client.NewMinioClient(vars)
	}
	return nil, errors.New(constant.PARAM_EMPTY)
}
