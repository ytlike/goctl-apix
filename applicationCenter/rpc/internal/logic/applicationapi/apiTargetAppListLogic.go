package applicationapilogic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"qbq-open-platform/applicationCenter/gorm/model"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/utils"

	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	"qbq-open-platform/applicationCenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiTargetAppListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApiTargetAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiTargetAppListLogic {
	return &ApiTargetAppListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询可投递的应用列表
func (l *ApiTargetAppListLogic) ApiTargetAppList(in *pb.ApiTargetAppListReq) (*pb.ApiTargetAppListResp, error) {
	var totalSize int64
	result := l.svcCtx.DbEngin.WithContext(l.ctx).
		Table("application_base_info").
		Where("app_status = 1 and deleted = 0").
		Count(&totalSize)
	if result.Error != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
	}

	pager := utils.Paging(in.Page, in.Size, int32(totalSize))
	var appModelList []*model.ApplicationBaseInfoModel
	result = l.svcCtx.DbEngin.WithContext(l.ctx).
		Where("app_status = 1 and deleted = 0").
		Offset(int(pager.Offset)).
		Limit(int(pager.Limit)).
		Find(&appModelList)
	if result.Error != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
	}
	var targetAppList []*pb.ApiTargetApp
	err := copier.Copy(&targetAppList, &appModelList)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}

	for _, app := range targetAppList {
		paasFileUrl, err := l.svcCtx.PaasPlatform.GetPaasFileUrl(l.ctx, app.AppIcon)
		if err != nil {
			app.AppIcon = ""
		} else {
			app.AppIcon = paasFileUrl.Data
		}
	}

	resp := &pb.ApiTargetAppListResp{
		AppList: targetAppList,
	}
	return resp, nil
}
