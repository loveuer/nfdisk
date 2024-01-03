package controller

import (
	"nfdisk/internal/model"
	"nfdisk/internal/util"
)

func (a *App) DoConnect(id int64) *model.RespObjectList {
	var buckets = make([]*model.NFObject, 0)

	err := model.DoConnect(util.Timeout(), id)
	if err != nil {
		return &model.RespObjectList{
			Status: 500,
			Msg:    util.MSG500,
			Err:    err.Error(),
		}
	}

	if buckets, err = model.ListBucket(util.Timeout(), id); err != nil {
		return &model.RespObjectList{
			Status: 500,
			Msg:    util.MSG500,
			Err:    err.Error(),
		}
	}

	return &model.RespObjectList{
		Status: 200,
		Msg:    util.MSG200,
		Data:   buckets,
	}
}
