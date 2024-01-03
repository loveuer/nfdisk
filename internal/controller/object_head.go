package controller

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"nfdisk/internal/model"
	"path"
)

func (a *App) HeadObject(id int64, bucket string, key string) *model.RespObject {
	client, ok := model.Clients[id]
	if !ok {
		return &model.RespObject{Status: 500, Err: "client not found"}
	}

	if client == nil {
		return &model.RespObject{Status: 500, Err: "client is nil"}
	}

	resp, err := client.HeadObject(a.ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return &model.RespObject{Status: 500, Err: err.Error()}
	}

	data := &model.NFObject{
		Name:         path.Base(key),
		Key:          key,
		LastModified: resp.LastModified.UnixMilli(),
		Size:         *resp.ContentLength,
		Type:         "file",
		ContentType:  *resp.ContentType,
	}

	return &model.RespObject{Status: 200, Data: data}
}
