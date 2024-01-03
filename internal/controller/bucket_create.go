package controller

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"nfdisk/internal/model"
	"nfdisk/internal/util"
)

func (a *App) CreateBucket(id int64, name string) *model.RespObjectList {
	client, ok := model.Clients[id]
	if !ok {
		return &model.RespObjectList{
			Status: 500,
			Err:    "client not found",
		}
	}

	if client == nil {
		return &model.RespObjectList{
			Status: 500,
			Err:    "client is nil",
		}
	}

	_, err := client.CreateBucket(util.Timeout(10), &s3.CreateBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		return &model.RespObjectList{
			Status: 500,
			Err:    err.Error(),
			Msg:    "创建 bucket 失败",
		}
	}

	list, err := model.ListBucket(util.Timeout(10), id)
	if err != nil {
		return &model.RespObjectList{
			Status: 500,
			Err:    err.Error(),
			Msg:    "获取 buckets 失败",
		}
	}

	return &model.RespObjectList{Status: 200, Data: list, Msg: "新建 bucket 成功"}
}
