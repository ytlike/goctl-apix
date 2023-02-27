package enums

// 应用中心枚举
const (
	//应用能力平台地址类型
	APP_ABILITY_URL_TYPE_ALIYUN            int32 = 1
	APP_ABILITY_URL_TYPE_DIGITAL_CHONGQING int32 = 2
	APP_ABILITY_URL_TYPE_OTHER             int32 = 3

	//应用状态
	APP_STATUS_ON    int32 = 1
	APP_STATUS_CLOSE int32 = 2
)

func init() {
	message[APP_ABILITY_URL_TYPE_ALIYUN] = "paas能力平台地址类型：1-阿里云"
	message[APP_ABILITY_URL_TYPE_DIGITAL_CHONGQING] = "paas能力平台地址类型：2-数字重庆"
	message[APP_ABILITY_URL_TYPE_OTHER] = "paas能力平台地址类型： 3-其他"
	message[APP_STATUS_ON] = "状态：0-停用；"
	message[APP_STATUS_CLOSE] = "状态：1-启用；"
}
