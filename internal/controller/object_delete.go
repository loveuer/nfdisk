package controller

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
	"nfdisk/internal/model"
)

func (a *App) DeleteObject(id int64, bucket string, key string) *model.RespMsg {
	var (
		err error
	)

	client, ok := model.Clients[id]
	if !ok {
		return &model.RespMsg{Status: 500, Err: "client not found"}
	}

	if client == nil {
		return &model.RespMsg{Status: 500, Err: "client is nil"}
	}

	logrus.Debugf("DeleteObject: id=%d bucket=%s key=%s", id, bucket, key)

	if _, err = client.DeleteObject(a.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}); err != nil {
		return &model.RespMsg{Status: 500, Err: err.Error()}
	}

	return &model.RespMsg{Status: 200, Msg: "删除成功"}
}
