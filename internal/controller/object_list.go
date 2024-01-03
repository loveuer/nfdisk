package controller

import (
	"nfdisk/internal/model"
	"nfdisk/internal/util"
)

func (a *App) ListObject(id int64, bucket string, parent string, start string) *model.RespObjectList {
	list, err := model.ListObject(
		util.Timeout(),
		id,
		bucket,
		parent,
		start,
	)

	if err != nil {
		return &model.RespObjectList{
			Status: 500,
			Err:    err.Error(),
			Msg:    "",
		}
	}

	return &model.RespObjectList{Status: 200, Data: list, Msg: "获取列表成功"}
}
