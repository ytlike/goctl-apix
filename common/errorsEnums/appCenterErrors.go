package errorsEnums

// 应用中心错误码
const (
	APP_AUTHORIZATION_ERROR          int32 = 300001
	APP_DISABLE_ERROR                int32 = 300002
	APP_NOT_EXIST_ERROR              int32 = 300003
	APP_CODE_EXIST_ERROR             int32 = 300004
	APP_OPEN_ABILITY_ERROR           int32 = 300005
	APP_OPENED_ERROR                 int32 = 300006
	DATA_TYPE_CONVERT_ERROR          int32 = 300007
	APPLICATION_NAME_EXIST_ERROR     int32 = 300008
	APPLICATION_CALLBACK_EXIST_ERROR int32 = 300009
)

func init() {
	message[APP_AUTHORIZATION_ERROR] = "应用信息错误"
	message[APP_DISABLE_ERROR] = "应用已禁用"
	message[APP_NOT_EXIST_ERROR] = "应用不存在"
	message[APP_CODE_EXIST_ERROR] = "应用编码已存在"
	message[APP_OPEN_ABILITY_ERROR] = "开通应用签署能力异常"
	message[APP_OPENED_ERROR] = "应用已经开通"
	message[DATA_TYPE_CONVERT_ERROR] = "数据类型转换异常"
	message[APPLICATION_NAME_EXIST_ERROR] = "应用名称已存在"
	message[APPLICATION_CALLBACK_EXIST_ERROR] = "应用回调信息已存在"

}
