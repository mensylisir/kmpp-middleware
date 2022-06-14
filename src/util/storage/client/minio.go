package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"path/filepath"
)

type minioClient struct {
	Vars   map[string]interface{}
	Client *minio.Client
}

func NewMinioClient(vars map[string]interface{}) (*minioClient, error) {

	var accessKey string
	var secretKey string
	var endpoint string

	if _, ok := vars["access_key"]; ok {
		accessKey = vars["access_Key"].(string)
	} else {
		return nil, errors.New(constant.PARAM_EMPTY)
	}
	if _, ok := vars["secret_Key"]; ok {
		secretKey = vars["secret_Key"].(string)
	} else {
		return nil, errors.New(constant.PARAM_EMPTY)
	}
	if _, ok := vars["endpoint"]; ok {
		endpoint = vars["endpoint"].(string)
	} else {
		return nil, errors.New(constant.PARAM_EMPTY)
	}

	target, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		logger.Log.Fatalln(err)
		return nil, err
	}

	return &minioClient{
		Client: target,
		Vars:   vars,
	}, nil
}

func (mc *minioClient) CreateBuckets() error {
	var bucketName string
	var location string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
	} else {
		return errors.New(constant.PARAM_EMPTY)
	}
	if _, ok := mc.Vars["location"]; ok {
		location = mc.Vars["location"].(string)
	} else {
		return errors.New(constant.PARAM_EMPTY)
	}
	err := mc.Client.MakeBucket(context.TODO(), bucketName, minio.MakeBucketOptions{
		Region: location,
	})
	if err != nil {
		exists, errBucketExists := mc.Client.BucketExists(context.TODO(), bucketName)
		if errBucketExists == nil && exists {
			errMsg := fmt.Sprintf("We already own %s\n", bucketName)
			logger.Log.Error(errMsg)
			return errors.New(errMsg)
		} else {
			logger.Log.Error(err.Error())
			return err
		}
	} else {
		msg := fmt.Sprintf("Successfully created %s\n", bucketName)
		logger.Log.Info(msg)
		return nil
	}
}

func (mc *minioClient) ListBuckets() ([]interface{}, error) {
	var result []interface{}
	buckets, err := mc.Client.ListBuckets(context.TODO())
	if err != nil {
		return nil, err
	}
	for _, bucket := range buckets {
		result = append(result, bucket.Name)
	}
	return result, nil
}

func (mc *minioClient) ExistBucket() (bool, error) {
	var bucketName string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
		ok, err := mc.Client.BucketExists(context.TODO(), bucketName)
		if err != nil {
			logger.Log.Warn("check bucket exist error ")
			return false, nil
		}
		return ok, nil
	} else {
		return false, errors.New(constant.PARAM_EMPTY)
	}
}

func (mc *minioClient) RemoveBucket() error {
	var bucketName string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
		err := mc.Client.RemoveBucket(context.TODO(), bucketName)
		if err != nil {
			logger.Log.Error("remove bucket error ")
			return err
		}
	} else {
		return errors.New(constant.PARAM_EMPTY)
	}
	return nil
}

func (mc *minioClient) ListObjects(prefix string) (<-chan minio.ObjectInfo, error) {
	var bucketName string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
	} else {
		return nil, errors.New(constant.PARAM_EMPTY)
	}
	opts := minio.ListObjectsOptions{
		UseV1:     true,
		Prefix:    prefix,
		Recursive: true,
	}
	objectinfo := mc.Client.ListObjects(context.TODO(), bucketName, opts)
	return objectinfo, nil
}

func (mc *minioClient) Exist(path string) (bool, error) {
	var bucketName string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
		_, fileName := filepath.Split(path)
		_, err := mc.Client.StatObject(context.TODO(), bucketName, fileName, minio.StatObjectOptions{})
		if err != nil {
			logger.Log.Error("check bucket exist error ")
			return false, nil
		}
		return true, nil
	} else {
		return false, errors.New(constant.PARAM_EMPTY)
	}
}

func (mc *minioClient) Delete(path string) (bool, error) {
	var bucketName string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
		_, fileName := filepath.Split(path)
		opts := minio.RemoveObjectOptions{
			GovernanceBypass: true,
		}
		err := mc.Client.RemoveObject(context.TODO(), bucketName, fileName, opts)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New(constant.PARAM_EMPTY)
}

func (mc *minioClient) Upload(src, target string) (bool, error) {
	var bucketName string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
		_, fileName := filepath.Split(target)
		_, err := mc.Client.FPutObject(context.TODO(), bucketName, fileName, src, minio.PutObjectOptions{ContentType: ""})
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New(constant.PARAM_EMPTY)
}

func (mc *minioClient) Download(src, target string) (bool, error) {
	var bucketName string
	if _, ok := mc.Vars["bucket_name"]; ok {
		bucketName = mc.Vars["bucket_name"].(string)
		_, fileName := filepath.Split(target)
		err := mc.Client.FGetObject(context.TODO(), bucketName, fileName, target, minio.GetObjectOptions{})
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New(constant.PARAM_EMPTY)
}
