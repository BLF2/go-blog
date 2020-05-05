package dto

type BaseResult struct {
	Code string `json:"code"`
	Msg string  `json:"msg"`
	Data interface{} `json:"data"`
}
