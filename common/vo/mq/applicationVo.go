package vo

import "time"

type ApplicationMQVo struct {
	Id             int64      `json:"id"`
	AppCode        string     `json:"appCode"`
	AppName        string     `json:"appName"`
	AppDeveloperId int64      `gorm:"type:int64"`
	AppKey         *string    `json:"appKey"`
	AppSecret      *string    `json:"appSecret"`
	AppAbilityUrl  *string    `json:"appAbilityUrl"`
	AppStatus      int32      `json:"appStatus"`
	Deleted        int32      `json:"deleted"`
	CreateTime     *time.Time `json:"createTime"`
	UpdateTime     *time.Time `json:"updateTime"`
}
