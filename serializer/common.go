package serializer

type Response struct{
	Code int 			`json:"code"`
	Data interface{} 	`json:"data"`
	Msg string			`json:"msg"`
}

func SuccessData(data interface{}) Response {
	var res Response
	res.Code = 0
	res.Msg = "ok"
	res.Data = data
	return res
}

func ErrorData(msg error, code ...int) Response {
	var res Response
	if len(code) > 0 {
		res.Code = code[0]
	}else{
		res.Code = 1
	}
	res.Msg = msg.Error()
	res.Data = nil
	return res
}

func ErrorMsg(msg string, code ...int) Response {
	var res Response
	if len(code) > 0 {
		res.Code = code[0]
	} else {
		res.Code = 1
	}
	res.Msg = msg
	res.Data = nil
	return res
}
