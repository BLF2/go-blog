package constant

import "go-blog/dto"

var SUCCESS = dto.BaseResult{
	Code: "000000",
	Msg: "success",
}

var FAIL = dto.BaseResult{
	Code: "0000001",
	Msg: "fail",
}

var FORBIDDEN = dto.BaseResult{
	Code: "403000",
	Msg:  "need login",
}

var LOGIN_FAIL = dto.BaseResult{
	Code: "403001",
	Msg:  "登陆失败，用户名不存在或者用户名与密码不存在",
	Data: nil,
}
