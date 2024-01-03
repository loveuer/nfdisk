package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/sirupsen/logrus"
	"nfdisk/internal/model"
	"nfdisk/internal/util"
	"regexp"
	"strings"
)

func (a *App) UploadObject(id int64, bucket string, path string, key string, content string) *model.RespObjectList {
	var (
		err      error
		fullpath string
	)

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

	if path == "" {
		fullpath = key
	} else {
		fullpath = fmt.Sprintf("%s/%s", path, key)
	}

	logrus.Debugf("UploadObject: fullpath=%s", fullpath)

	if ok, err = regexp.MatchString(`^data:.+;base64,.+$`, content); !ok {
		if !ok || err != nil {
			logrus.Errorf("UploadObject: reg match ok=%v err=%v", ok, err)
			return &model.RespObjectList{Status: 500, Msg: "文件内容错误1"}
		}
	}

	strs := strings.Split(content, ";")
	if len(strs) != 2 {
		return &model.RespObjectList{Status: 500, Msg: "文件内容错误2"}
	}

	contentType := strings.TrimPrefix(strs[0], "data:")

	bs, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(strs[1], "base64,"))
	if err != nil {
		logrus.Errorf("UploadObject: decode base64 conetent err=%v", err)
		return &model.RespObjectList{Status: 500, Err: err.Error()}
	}

	_, err = client.PutObject(
		util.Timeout(60),
		&s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(fullpath),
			Body:        bytes.NewReader(bs),
			ContentType: aws.String(contentType),
			ACL:         types.ObjectCannedACLPublicRead,
		},
	)
	if err != nil {
		return &model.RespObjectList{Status: 500, Err: err.Error()}
	}

	return &model.RespObjectList{Status: 200, Msg: "上传文件成功"}
}
