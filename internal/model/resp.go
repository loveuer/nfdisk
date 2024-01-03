package model

type RespConnectionList struct {
	Status int32         `json:"status"`
	Msg    string        `json:"msg"`
	Err    string        `json:"err"`
	Data   []*Connection `json:"data"`
}

type RespObjectList struct {
	Status int32       `json:"status"`
	Msg    string      `json:"msg"`
	Err    string      `json:"err"`
	Data   []*NFObject `json:"data"`
}

type RespObject struct {
	Status int32     `json:"status"`
	Msg    string    `json:"msg"`
	Err    string    `json:"err"`
	Data   *NFObject `json:"data"`
}

type RespMsg struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"err"`
}

type RespShare struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"err"`
	Data   string `json:"data"`
}
