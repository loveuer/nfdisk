package controller

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"nfdisk/internal/model"
	"nfdisk/internal/util"
	"os"
	"path"
	"time"
)

func (a *App) GetObject(id int64, bucket string, key string) *model.RespMsg {
	client, ok := model.Clients[id]
	if !ok {
		return &model.RespMsg{
			Status: 500,
			Err:    "client not found",
		}
	}

	if client == nil {
		return &model.RespMsg{Status: 500, Err: "client is nil"}
	}

	resp, err := client.GetObject(
		util.Timeout(60),
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
	)
	if err != nil {
		return &model.RespMsg{Status: 500, Err: err.Error()}
	}

	filename := path.Base(key)

	location, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory:     "",
		DefaultFilename:      filename,
		Title:                fmt.Sprintf("下载文件: %s", filename),
		Filters:              nil,
		ShowHiddenFiles:      true,
		CanCreateDirectories: true,
	})
	logrus.Debugf("GetObject: download bucket=%s key=%s to location=%s", bucket, key, location)
	if err != nil {
		logrus.Errorf("GetObject: runtime dialog err=%v", err)
		return &model.RespMsg{Status: 500, Err: err.Error()}
	}

	fileio, err := os.OpenFile(location, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		logrus.Errorf("GetObject: open location=%s err=%v", location, err)
		return &model.RespMsg{Status: 500, Err: err.Error()}
	}

	if _, err = io.Copy(fileio, resp.Body); err != nil {
		logrus.Errorf("GetObject: io copy err=%v", err)
		return &model.RespMsg{Status: 500, Err: err.Error()}
	}

	return &model.RespMsg{Status: 200, Msg: "成功"}
}

type Presigner struct {
	PresignClient *s3.PresignClient
}

// GetObject
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/example_s3_Scenario_PresignedUrl_section.html
func (presigner Presigner) GetObject(
	bucketName string, objectKey string, lifetimeSecs int64) (*v4.PresignedHTTPRequest, error) {
	request, err := presigner.PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		logrus.Errorf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return request, err
}

func (a *App) ShareObject(id int64, bucket string, key string) *model.RespShare {
	client, ok := model.Clients[id]
	if !ok {
		return &model.RespShare{
			Status: 500,
			Err:    "client not found",
		}
	}

	if client == nil {
		return &model.RespShare{Status: 500, Err: "client is nil"}
	}

	signer := Presigner{
		PresignClient: s3.NewPresignClient(client),
	}
	req, err := signer.GetObject(bucket, key, 10*60)
	if err != nil {
		return &model.RespShare{Status: 500, Err: err.Error()}
	}

	return &model.RespShare{Status: 200, Data: req.URL}
}
