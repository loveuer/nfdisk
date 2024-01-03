package controller

import (
	"nfdisk/internal/model"
)

func (a *App) ListConnection(page, size int) *model.RespConnectionList {
	list, err := model.ConnectionList()
	if err != nil {
		return &model.RespConnectionList{
			Status: 500,
			Msg:    err.Error(),
			Data:   list,
		}
	}

	return &model.RespConnectionList{
		Status: 200,
		Msg:    "请求成功",
		Data:   list,
	}
}
