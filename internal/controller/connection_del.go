package controller

import "nfdisk/internal/model"

func (a *App) DelConnection(id int64) *model.RespConnectionList {
	if err := model.DelConnection(id); err != nil {
		return &model.RespConnectionList{
			Status: 500,
			Msg:    "",
			Err:    err.Error(),
		}
	}

	list, err := model.ConnectionList()
	if err != nil {
		return &model.RespConnectionList{
			Status: 500,
			Msg:    "",
			Err:    err.Error(),
		}
	}

	return &model.RespConnectionList{
		Status: 200,
		Msg:    "获取列表成功",
		Data:   list,
	}
}
