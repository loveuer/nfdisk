package controller

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
	"nfdisk/internal/model"
	"nfdisk/internal/util"
	"strings"
)

func (a *App) DeleteBucket(id int64, name string) *model.RespObjectList {
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

	_, err := client.DeleteBucket(util.Timeout(10), &s3.DeleteBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		if strings.Contains(err.Error(), "BucketNotEmpty") {
			return &model.RespObjectList{Status: 500, Err: err.Error(), Msg: "删除失败: 无法删除有内容的桶"}
		}

		logrus.Errorf("DeleteBucket: err=%v", err)
		return &model.RespObjectList{Status: 500, Err: err.Error(), Msg: "删除失败"}
	}

	list, err := model.ListBucket(util.Timeout(10), id)
	if err != nil {
		return &model.RespObjectList{Status: 500, Err: err.Error(), Msg: "刷新列表失败"}
	}

	return &model.RespObjectList{Status: 200, Data: list, Msg: "删除桶成功"}
}
