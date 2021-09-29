package structs

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func ToClientData(status int, message string, data interface{}) (res Response) {
	res.Status = status
	res.Data = data
	res.Message = message
	return
}

func (this *Response) ToClientData() {

	if this.Message == "" {
		this.Status = 0 // int 类型本身零值就是 0，在这里不需要特别赋值的
		this.Message = "success"
	} else {
		this.Status = 1 // 有错误消息了
	}
}
