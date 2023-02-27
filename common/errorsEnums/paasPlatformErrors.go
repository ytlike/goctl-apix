package errorsEnums

// 任务中心错误码
const (
	CALL_PAAS_ERROR              int32 = 900001
	GET_PAAS_TOKEN_ERROR         int32 = 900002
	GET_PAAS_FILE_METADATA_ERROR int32 = 900003
	GET_PAAS_FILE_URL_ERROR      int32 = 900004
	ADD_CLIENT_ERROR             int32 = 900005
)

func init() {
	message[CALL_PAAS_ERROR] = "调用paas平台接口失败"
	message[GET_PAAS_TOKEN_ERROR] = "获取paas平台token失败"
	message[GET_PAAS_FILE_METADATA_ERROR] = "获取paas平台文件元数据错误"
	message[GET_PAAS_FILE_URL_ERROR] = "获取paas平台文件访问地址错误"
	message[ADD_CLIENT_ERROR] = "新增Client错误"

}
