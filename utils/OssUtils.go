package utils

import (
	"fmt"
	"mime/multipart"
	"strings"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// LoadConfig 读取并解析 YAML 配置文件
// func LoadConfig(configPath string) (*Config, error) {
//     data, err := os.ReadFile(configPath)
//     if err != nil {
//         return nil, fmt.Errorf("error reading config file: %v", err)
//     }

//     var config Config
//     err = yaml.Unmarshal(data, &config)
//     if err != nil {
//         return nil, fmt.Errorf("error parsing config file: %v", err)
//     }

//     return &config, nil
// }

// Upload file to aliyun OSS and return the URL
func UploadFileToOss(file *multipart.FileHeader) (string, error) {
	// Get the properties of aliyun OSS, Create Oss Client
	config, err := LoadConfig("./config/config.yaml")
	if err != nil {
		return "", err
	}
	client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyId, config.Oss.AccessKeySecret)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(config.Oss.BucketName)
	if err != nil {
		return "", err
	}
	
	// Upload file
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	// create uuid for file name
	uuid := uuid.New()
	// get file suffix name
	suffixName := file.Filename[strings.LastIndex(file.Filename, "."):]
	filename := uuid.String() + suffixName
	path := "sky-take-out-go/" + filename
	// upload this file 
	err = bucket.PutObject(path, src)
	if err != nil {
		return "", err
	}

	err = bucket.SetObjectACL(path, oss.ACLPublicRead)
    if err != nil {
        return "", err
    }

	url := fmt.Sprintf("https://%s.%s/%s", config.Oss.BucketName, config.Oss.Endpoint, path)

	return url, nil
}