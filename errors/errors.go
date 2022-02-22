package errors

// 自定义错误数据，也没太有必要

type MyError struct {
	Code    int
	Message string
	Data    interface{}
}

var (
	LOGIN_UNKNOWN = NewError(202, "用户不存在")
	LOGIN_ERROR   = NewError(203, "账号或密码错误")
	VALID_ERROR   = NewError(300, "参数错误")
	ERROR         = NewError(400, "操作失败")
	UNAUTHORIZED  = NewError(401, "您还未登录")
	NOT_FOUND     = NewError(404, "资源不存在")
	INNER_ERROR   = NewError(500, "系统发生异常")
)

func (this *MyError) Error() string {
	return this.Message
}

func NewError(code int, message string) *MyError {
	return &MyError{
		Message: message,
		Code:    code,
	}
}

func GetError(e *MyError, data interface{}) *MyError {
	e.Data = data
	return e
}
