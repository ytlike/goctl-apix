package errorsEnums

// 任务中心错误码
const (
	TASK_PUBLISHER_NOT_MATCH     int32 = 200001
	TASK_NOT_EXIST_ERROR         int32 = 200002
	DATA_NOT_EXIST               int32 = 200003
	CANNOT_MODIFY                int32 = 200004
	APP_NOT_MATCH                int32 = 200005
	TASK_STATUS_UNSUPPORTED      int32 = 200006
	SIGNATORY_STATUS_UNSUPPORTED int32 = 200007
	SIGNER_ALREADY_EXISTS        int32 = 200008
	TASK_ALREADY_EXPIRED         int32 = 200009
	TASK_ALREADY_CANCEL          int32 = 200010
	SIGNATORY_NOT_MATCH          int32 = 200011
	PUBLISHER_NOT_MATCH          int32 = 200012
	TASK_ALREADY_REJECTED        int32 = 200013
	SIGNER_NOT_MATCH             int32 = 200014
	SIGN_FILE_NOT_MATCH          int32 = 200015
	TASK_BIZ_ALREADY             int32 = 200016
)

func init() {
	message[TASK_PUBLISHER_NOT_MATCH] = "任务和发布人不匹配"
	message[TASK_NOT_EXIST_ERROR] = "任务不存在"
	message[DATA_NOT_EXIST] = "无法匹配任务"
	message[CANNOT_MODIFY] = "无法修改"
	message[APP_NOT_MATCH] = "应用不匹配"
	message[TASK_STATUS_UNSUPPORTED] = "任务状态不支持此操作"
	message[SIGNATORY_STATUS_UNSUPPORTED] = "签署方状态不支持此操作"
	message[SIGNER_ALREADY_EXISTS] = "签署人已存在"
	message[TASK_ALREADY_EXPIRED] = "任务已过期"
	message[TASK_ALREADY_CANCEL] = "任务已撤销"
	message[SIGNATORY_NOT_MATCH] = "签署方不匹配"
	message[PUBLISHER_NOT_MATCH] = "发布人不匹配"
	message[TASK_ALREADY_REJECTED] = "任务已拒签"
	message[SIGNER_NOT_MATCH] = "签署人不匹配"
	message[SIGN_FILE_NOT_MATCH] = "签署文件不匹配"
	message[TASK_BIZ_ALREADY] = "task_biz已存在"
}
