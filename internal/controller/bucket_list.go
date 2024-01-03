package controller

import (
	"nfdisk/internal/model"
	"nfdisk/internal/util"
)

func (a *App) ListBucket(id int64) *model.RespObjectList {
	list, err := model.ListBucket(util.Timeout(10), id)
	if err != nil {
		return &model.RespObjectList{Status: 500, Msg: "获取 bucket 列表失败", Err: err.Error()}
	}

	return &model.RespObjectList{Status: 200, Data: list}
}
