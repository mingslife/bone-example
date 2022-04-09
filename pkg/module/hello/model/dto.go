package model

type HelloReq struct {
	Name string
}

type HelloRsp struct {
	Msg string `json:"msg"`
}
