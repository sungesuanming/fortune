package errno

var (
	// Common errors
	SendOK              = &Errno{Code: 200, Message: "发送成功"}
	LoginOK             = &Errno{Code: 200, Message: "发送成功"}
	RegisterOK          = &Errno{Code: 200, Message: "注册成功"}
	OK                  = &Errno{Code: 200, Message: "成功"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrParamIsNull  = &Errno{Code: 10003, Message: "The param should not be null."}
	ErrParamInvalid = &Errno{Code: 10004, Message: "The param type is invalid."}
	ErrNoAuth       = &Errno{Code: 10005, Message: "暂无该功能操作权限，请检查登录状态和角色身份"}
	// player errors

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrRedis      = &Errno{Code: 20004, Message: "Redis server error."}

	ErrUserNotFound             = &Errno{Code: 20102, Message: "用户不存在."}
	ErrTokenInvalid             = &Errno{Code: 20103, Message: "用户未登录."}
	ErrPasswordIncorrect        = &Errno{Code: 20104, Message: "密码错误."}
	ErrSendMessage              = &Errno{Code: 20105, Message: "验证码发送失败."}
	ErrVerification             = &Errno{Code: 20106, Message: "验证码错误."}
	ErrVerificationNil          = &Errno{Code: 20107, Message: "未获取验证码."}
	ErrUserIsExist              = &Errno{Code: 20108, Message: "用户已存在."}
	ErrFileIsExist              = &Errno{Code: 20109, Message: "文件已存在."}
	ApplicationIsNotExist       = &Errno{Code: 20110, Message: "用户未报名."}
	WeightIsExist               = &Errno{Code: 20111, Message: "权重值不可重复."}
	CategoryIsUpload            = &Errno{Code: 20112, Message: "同一类别只能提交一个案例."}
	CategoryCantModify          = &Errno{Code: 20113, Message: "已进入评审阶段，无法上传或修改参赛资料."}
	MatchUnbegin                = &Errno{Code: 20114, Message: "大赛未开始，无法上传或修改参赛资料."}
	ErrVerificationTooFrequency = &Errno{Code: 20115, Message: "验证码请求过于频繁."}

	ErrSeeLocation = &Errno{Code: 20301, Message: "是否消费10积分或充值VIP查看精准定位."}
	ErrIsDLocation = &Errno{Code: 20302, Message: "已经是可看状态无需再次消费."}
)
