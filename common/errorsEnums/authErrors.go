package errorsEnums

const (
	TENANT_NOT_EXIST                 int32 = 500001
	TENANT_NOT_AVAILABLE             int32 = 500002
	TENANT_INFO_ERROR                int32 = 500003
	UNSUPPORTED_GRANT_TYPE           int32 = 500004
	MISSING_USERNAME_ERROR           int32 = 500005
	MISSING_PASSWORD_ERROR           int32 = 500006
	USER_INFO_ERROR                  int32 = 500007
	TENANT_INFO_MISS                 int32 = 500008
	TOKEN_ISEXPIRED                  int32 = 500009
	PERMISSION_DENIED                int32 = 500010
	AUTHORIZATION_TYPE_NOT_SUPPORTED int32 = 500011
	REMOTE_CHECK_ERROR               int32 = 500012
	TOKEN_GRANT_TYPE_NOT_SUPPORTED   int32 = 500013
	RELOAD_POLICY_ERROR              int32 = 500014
	INVALID_SIGN_ERROR               int32 = 500015
	VERIFY_SIGN_EXPIRE_ERROR         int32 = 500016
	VERIFY_SECRET_NOT_FOUND_ERROR    int32 = 500017
	UNSUPPORTED_ALGORITHM            int32 = 500018
	DEFECT_AUTHORIZATION_ERROR       int32 = 500019
	INVALID_TOKEN_ERROR              int32 = 500020
)

func init() {
	message[TENANT_NOT_AVAILABLE] = "租户不可用"
	message[TENANT_INFO_ERROR] = "租户信息错误"
	message[UNSUPPORTED_GRANT_TYPE] = "不支持的grantType"
	message[MISSING_USERNAME_ERROR] = "缺少用户名参数"
	message[MISSING_PASSWORD_ERROR] = "缺少密码参数"
	message[USER_NOT_EXIST] = "缺少密码参数"
	message[USER_INFO_ERROR] = "用户信息错误"
	message[TENANT_INFO_MISS] = "租户信息缺失"
	message[TOKEN_ISEXPIRED] = "token已过期"
	message[PERMISSION_DENIED] = "权限不足"
	message[AUTHORIZATION_TYPE_NOT_SUPPORTED] = "认证类型不支持"
	message[REMOTE_CHECK_ERROR] = "远程鉴权错误"
	message[TOKEN_GRANT_TYPE_NOT_SUPPORTED] = "token的grantType不支持"
	message[RELOAD_POLICY_ERROR] = "刷新权限错误"
	message[INVALID_SIGN_ERROR] = "无效的签名"
	message[VERIFY_SIGN_EXPIRE_ERROR] = "签名过期"
	message[VERIFY_SECRET_NOT_FOUND_ERROR] = "密钥不存在"
	message[UNSUPPORTED_ALGORITHM] = "不支持的签名算法"
	message[DEFECT_AUTHORIZATION_ERROR] = "缺失认证信息"
	message[INVALID_TOKEN_ERROR] = "无效的token"
}
