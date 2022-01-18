package response

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message" default:"success"`
	Data    interface{} `json:"data"`
}

func GetResponse() (res Response) {
	res = Response{}

	res.Message = "success"
	res.Status = 0
	res.Data = struct{}{}
	return
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

func (this *Response) SetData(data interface{}) {
	this.Data = data
}

func (this *Response) SetMessage(msg string) {
	this.Status = 1 // 一旦设置 message 。表示有错误信息了，同时设置一下 status
	this.Message = msg
}
