package errorsEnums

const (
	USER_EXIST             int32 = 400001
	EMAIL_VALIDATE_ERROR   int32 = 400002
	USER_NOT_EXIST         int32 = 400003
	PASSWORD_ERROR         int32 = 400004
	USER_NOT_AVAILABLE     int32 = 400005
	EMAIL_CODE_INCORRECT   int32 = 400006
	EMAIL_CODE_SEND_ERROR  int32 = 400007
	EMAIL_CODE_EXIST_ERROR int32 = 400008
	LOGIN_ERROR            int32 = 400009
)

func init() {
	message[USER_EXIST] = "用户已存在"
	message[EMAIL_VALIDATE_ERROR] = "邮箱格式错误"
	message[USER_NOT_EXIST] = "用户不存在"
	message[PASSWORD_ERROR] = "密码错误"
	message[USER_NOT_AVAILABLE] = "用户已被禁用"
	message[EMAIL_CODE_INCORRECT] = "邮箱验证码错误"
	message[EMAIL_CODE_SEND_ERROR] = "邮箱验证码发送失败"
	message[EMAIL_CODE_EXIST_ERROR] = "邮箱验证码已发送，请60秒后重试"
	message[LOGIN_ERROR] = "登录失败"
}
