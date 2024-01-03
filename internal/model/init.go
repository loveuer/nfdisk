package model

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	Clients = map[int64]*s3.Client{}
)

func Init() {
	var (
		err error
	)

	if db, err = gorm.Open(sqlite.Open(".nf-disk.sqlite"), &gorm.Config{}); err != nil {
		logrus.Fatalf("model.Init: gorm open err=%v", err)
	}

	if err = db.AutoMigrate(&Connection{}); err != nil {
		logrus.Fatalf("model.Init: auto migrate err=%v", err)
	}
}
