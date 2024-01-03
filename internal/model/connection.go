package model

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"nfdisk/internal/util"
)

type Connection struct {
	Id        int64  `gorm:"primaryKey;column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Endpoint  string `gorm:"column:endpoint" json:"endpoint"`
	AccessKey string `gorm:"column:access_key" json:"access_key"`
	SecretKey string `gorm:"column:secret_key" json:"secret_key"`
	Active    bool   `gorm:"-" json:"active"`
}

func initS3Client(gctx context.Context, endpoint, access, secret string) (*s3.Client, error) {
	var (
		err    error
		cfg    aws.Config
		client *s3.Client
	)

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: endpoint,
		}, nil
	})

	if cfg, err = config.LoadDefaultConfig(gctx,
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{AccessKeyID: access, SecretAccessKey: secret},
		}),
		config.WithRegion("auto"),
	); err != nil {
		return nil, err
	}

	client = s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
	})

	return client, nil
}

func AddConnection(name, endpoint, accessKey, secretKey string) error {
	connection := &Connection{
		Name:      name,
		Endpoint:  endpoint,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	return db.Create(connection).Error
}

func DelConnection(id int64) error {
	return db.Delete(&Connection{Id: id}).Error
}

func ConnectionList() ([]*Connection, error) {
	var (
		err  error
		list = make([]*Connection, 0)
	)

	err = db.Model(&Connection{}).
		Order("id ASC").
		//Offset(offset).
		//Limit(limit).
		Find(&list).Error

	return list, err
}

func DoConnect(ctx context.Context, id int64) error {
	var (
		err    error
		item   = &Connection{}
		client *s3.Client
	)

	if err = db.Session(&gorm.Session{}).
		WithContext(ctx).
		Model(&Connection{}).
		Where("id", id).
		Take(item).
		Error; err != nil {
		return err
	}

	logrus.Debugf("DoConnect: item.id=%d item.name=%s item.endpoint=%s item.access=%s item.secret=%s", item.Id, item.Name, item.Endpoint, item.AccessKey, item.SecretKey)

	if client, err = initS3Client(context.TODO(), item.Endpoint, item.AccessKey, item.SecretKey); err != nil {
		return err
	}

	Clients[item.Id] = client

	resp, err := client.ListBuckets(util.Timeout(), &s3.ListBucketsInput{})
	if err != nil {
		logrus.Errorf("DoConnect: list buckets err=%s", err.Error())
		return err
	}

	logrus.Infof("DoConnect: list buckets=%+v", resp.Buckets)

	return nil
}
