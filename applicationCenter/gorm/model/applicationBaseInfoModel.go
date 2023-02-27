package model

import "time"

type ApplicationBaseInfoModel struct {
	Id                int64      `gorm:"primaryKey"` // 主键
	AppCode           string     // 应用编码
	AppDeveloperId    int64      `gorm:"type:int64"` // 开发者用户id
	AppDeveloperName  *string    // 开发者名称
	AppIcon           *string    // 应用图标
	AppName           string     // 应用名称
	AppDesc           *string    // 应用描述
	AppKey            *string    // 应用key
	AppSecret         *string    // 应用secret
	AppAbilityUrl     *string    // 应用对应能力平台地址
	PaasAppId         *string    // 应用对应能力平台的appId
	Watermark         *string    // 水印内容
	WatermarkPosition *int32     // 水印位置 1-首页；2-每一页；
	AppStatus         int32      `gorm:"default:1"`           // 应用状态 0-停用；1-启用；
	Deleted           int32      `gorm:"default:0"`           // 是否删除 0-未删除；1-已删除；
	CreateTime        *time.Time `gorm:"column:create_time;"` // 任务创建时间
	UpdateTime        *time.Time `gorm:"column:update_time;"` // 任务更新时间
}

func (ApplicationBaseInfoModel) TableName() string {
	return "application_base_info"
}
