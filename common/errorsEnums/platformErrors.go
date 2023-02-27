package errorsEnums

import "fmt"

// 全局错误码
const (
	SUCCESS                 int32 = 200
	FAILURE                 int32 = 400
	INTERNAL_SERVER_ERROR   int32 = 500
	RPC_CALL_ERROR          int32 = 100000
	PARAM_MISS              int32 = 100001
	PARAM_TYPE_ERROR        int32 = 100002
	PARAM_VALID_ERROR       int32 = 100003
	REDIS_LOCK_ERROR        int32 = 100005
	DB_ERROR                int32 = 100006
	REDIS_ERROR             int32 = 100007
	API_TOKEN_RENEWAL_ERROR int32 = 100008
	ROCKETMQ_SEND_ERROR     int32 = 100012
)

var message map[int32]string = make(map[int32]string)

func init() {
	message[SUCCESS] = "操作成功"
	message[FAILURE] = "业务异常"
	message[INTERNAL_SERVER_ERROR] = "服务器异常"
	message[RPC_CALL_ERROR] = "RPC调用失败"
	message[PARAM_MISS] = "缺少必要的请求参数"
	message[PARAM_TYPE_ERROR] = "请求参数类型错误"
	message[PARAM_VALID_ERROR] = "参数校验失败"
	message[REDIS_LOCK_ERROR] = "获取分布式锁失败"
	message[DB_ERROR] = "数据库操作失败"
	message[REDIS_ERROR] = "redis操作失败"
	message[API_TOKEN_RENEWAL_ERROR] = "api端的token续期失败"
	message[ROCKETMQ_SEND_ERROR] = "MQ发送消息失败"
}

type PlatformError struct {
	errCode int32
	errMsg  string
	data    interface{}
	cause   error
}

func NewErrCode(errCode int32) *PlatformError {
	return &PlatformError{
		errCode: errCode,
		errMsg:  message[errCode],
	}
}

func NewErrCodeMsg(errCode int32, errMsg string) *PlatformError {
	return &PlatformError{
		errCode: errCode,
		errMsg:  errMsg,
	}
}
func NewErrCodeMsgData(errCode int32, errMsg string, data interface{}) *PlatformError {
	return &PlatformError{
		errCode: errCode,
		errMsg:  errMsg,
		data:    data,
	}
}
func NewErrCodeMsgCause(errCode int32, errMsg string, cause error) *PlatformError {
	return &PlatformError{
		errCode: errCode,
		errMsg:  errMsg,
		data:    nil,
		cause:   cause,
	}
}
func IsCodeErr(errCode int32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}

func GetMsgByCode(errCode int32) string {
	return message[errCode]
}

func (e *PlatformError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func (e *PlatformError) GetErrCode() int32 {
	return e.errCode
}
func (e *PlatformError) CodeIs(code int32) bool {
	return e.errCode == code
}

func (e *PlatformError) GetErrMsg() string {
	return e.errMsg
}
func (e *PlatformError) GetErrData() interface{} {
	return e.data
}

func IsPlatformErr(err error) (*PlatformError, bool) {
	if err == nil {
		return nil, false
	}

	platformError, ok := err.(*PlatformError)
	if ok {
		return platformError, true
	}
	if errWithCause, ok := err.(interface {
		Cause() error
	}); ok {
		causeErr := errWithCause.Cause()
		if causeErr == nil || err == causeErr {
			return nil, false
		}
		return IsPlatformErr(causeErr)
	}
	return nil, false
}
