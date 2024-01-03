package model

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"strings"
)

type NFObject struct {
	Name         string `json:"name"`
	Key          string `json:"key"`
	LastModified int64  `json:"last_modified"`
	Size         int64  `json:"size"`
	Type         string `json:"type"` // bucket, folder, file
	ContentType  string `json:"content_type"`
}

func ListBucket(ctx context.Context, id int64) ([]*NFObject, error) {
	client, ok := Clients[id]
	if !ok {
		return nil, errors.New("client not found")
	}

	if client == nil {
		return nil, errors.New("client nil")
	}

	resp, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		logrus.Errorf("ListBucket: client=%d err=%v", id, err)
		return nil, err
	}

	logrus.Debugf("ListBucket: buckets=%+v", resp.Buckets)

	return lo.Map(resp.Buckets, func(t types.Bucket, i int) *NFObject {
		return &NFObject{
			Name:         *t.Name,
			LastModified: t.CreationDate.UnixMilli(),
			Type:         "bucket",
		}
	}), nil
}

func ListObject(ctx context.Context, id int64, bucket string, parent string, start string) ([]*NFObject, error) {
	client, ok := Clients[id]
	if !ok {
		return nil, errors.New("client not found")
	}

	if client == nil {
		return nil, errors.New("client nil")
	}

	logrus.Debugf("ListObject: id=%d bucket=%s prefix=%s start=%s", id, bucket, parent, start)

	req := &s3.ListObjectsV2Input{
		Bucket:     aws.String(bucket),
		MaxKeys:    aws.Int32(1000),
		Prefix:     aws.String(parent),
		StartAfter: aws.String(start),
		Delimiter:  aws.String("/"),
	}

	resp, err := client.ListObjectsV2(ctx, req)
	if err != nil {
		logrus.Errorf("ListObject: id=%d bucket=%s prefix=%s start=%s err=%v", id, bucket, parent, start, err)
		return nil, err
	}

	list := make([]*NFObject, 0)
	for idx := range resp.CommonPrefixes {
		name := strings.TrimPrefix(*resp.CommonPrefixes[idx].Prefix, parent)
		if name == "" {
			continue
		}

		list = append(list, &NFObject{
			Name: name,
			Type: "folder",
		})
	}

	list = append(list, lo.Map(resp.Contents, func(item types.Object, i int) *NFObject {
		isDir := strings.HasSuffix(*item.Key, "/")
		t := "file"
		if isDir {
			t = "folder"
		}
		return &NFObject{
			Key:          *item.Key,
			Name:         strings.TrimPrefix(*item.Key, parent),
			LastModified: item.LastModified.UnixMilli(),
			Size:         *item.Size,
			Type:         t,
		}
	})...)

	//lo.Map(list, func(item *NFObject, index int) bool {
	//	resp, err := client.HeadObject(ctx, &s3.HeadObjectInput{
	//		Bucket: aws.String(bucket),
	//		Key:    aws.String(item.key),
	//	})
	//	if err != nil {
	//		return false
	//	}
	//
	//	item.ContentType = *resp.ContentType
	//
	//	return true
	//})

	return list, nil
}
