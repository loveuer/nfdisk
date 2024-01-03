package model

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/labstack/gommon/log"
	"github.com/samber/lo"
	"os"

	"testing"
)

var (
	tc *s3.Client
)

func testInit() {

	var (
		err  error
		item = &Connection{
			Id:        1,
			Name:      "test",
			Endpoint:  "http://10.220.10.57:7480",
			AccessKey: "A1S6ODB20RS2ILHA1HL3",
			SecretKey: "36bUlZQ4iJ0u46gG4fwOcoHFqK8zKCZ1Hc8UZs51",
		}
	)

	log.Printf("DoConnect: item.id=%d item.name=%s item.endpoint=%s item.access=%s item.secret=%s", item.Id, item.Name, item.Endpoint, item.AccessKey, item.SecretKey)

	if tc, err = initS3Client(context.TODO(), item.Endpoint, item.AccessKey, item.SecretKey); err != nil {
		log.Fatal(0, err)
	}
}

func TestListBucket(t *testing.T) {

	resp, err := tc.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		t.Error(2, err)
		return
	}

	lo.Map(resp.Buckets, func(o types.Bucket, i int) error {
		t.Logf("list bucket: idx=%d bucket.name=%s", i, *o.Name)
		return nil
	})
}
func TestPutObject(t *testing.T) {
	testInit()
	f, err := os.Open("./avatar.jpg")
	if err != nil {
		t.Error(1, err)
		return
	}
	defer f.Close()

	resp, err := tc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("other"),
		Key:    aws.String("base.jpg"),
		//Body:   f,
	}, func(options *s3.Options) {
		options.UsePathStyle = true
	})
	if err != nil {
		t.Error(2, err)
	}

	_ = resp
}

func TestListObject(t *testing.T) {
	testInit()

	items, err := tc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:    aws.String("other"),
		Delimiter: aws.String("/"),
		MaxKeys:   aws.Int32(20),
		Prefix:    aws.String("dir1/dir2/"),
	})
	if err != nil {
		t.Error(3, err)
		return
	}

	for idx := range items.CommonPrefixes {
		t.Logf("list object: idx=%d one.dir =%s", idx, *items.CommonPrefixes[idx].Prefix)
	}

	lo.Map(items.Contents, func(o types.Object, idx int) error {
		t.Logf("list object: idx=%d one.name=%s", idx, *o.Key)
		return nil
	})
}
