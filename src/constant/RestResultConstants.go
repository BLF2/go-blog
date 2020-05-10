package constant

import "go-blog/src/dto"

var SUCCESS = dto.BaseResult{
	Code: "000000",
	Msg: "success",
}

var FAIL = dto.BaseResult{
	Code: "0000001",
	Msg: "fail",
}
