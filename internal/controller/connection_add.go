package controller

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/sirupsen/logrus"
	"nfdisk/internal/model"
)

func (a *App) AddConnection(name, endpoint, accessKey, secretKey string) *model.RespConnectionList {
	logrus.Debugf("AddConnection: name=%s endpoint=%s access_key=%s secret_key=%s", name, endpoint, accessKey, secretKey)

	var (
		err error
	)

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: endpoint,
		}, nil
	})

	if _, err = config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{AccessKeyID: accessKey, SecretAccessKey: secretKey},
		}),
		config.WithRegion("auto"),
	); err != nil {
		return &model.RespConnectionList{Status: 400, Msg: "配置错误", Data: nil, Err: err.Error()}
	}

	if err = model.AddConnection(name, endpoint, accessKey, secretKey); err != nil {
		return &model.RespConnectionList{
			Status: 500,
			Msg:    "新建连接错误",
			Data:   nil,
			Err:    err.Error(),
		}
	}

	list, err := model.ConnectionList()
	if err != nil {
		return &model.RespConnectionList{
			Status: 500,
			Msg:    "获取连接列表错误",
			Err:    err.Error(),
		}
	}

	return &model.RespConnectionList{
		Status: 200,
		Msg:    "新建成功",
		Data:   list,
	}
}
