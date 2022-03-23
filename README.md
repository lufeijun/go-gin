# 说明

gin 练手


# gin 请求参数获取

1、c.Request.PostForm 可以获取所有请求参数，但是在此之前，必须调用一次 c.PostForm("ooxx")，否则 c.Request.PostForm 值为 null。原因：PostForm 底层有类似初始化赋值的函数执行


# 协程部分

1、开启协程处理函数，在协程函数中必须有处理异常的部分，即 defer 下的 recover，否则整个服务会挂掉
