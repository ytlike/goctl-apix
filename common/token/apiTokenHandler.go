package token

//var (
//	ApiTokenHandler TokenHandle
//)
//
//func init() {
//	ApiTokenHandler = NewApiTokenHandler()
//}
//
//// apiTokenHandler Api端token处理类
//type apiTokenHandler struct {
//}
//
//func NewApiTokenHandler() *apiTokenHandler {
//	return &apiTokenHandler{}
//}
//
//func (t *apiTokenHandler) ApiCheckToken(ctx *context.Context, accessToken string) error {
//	tokenKey := global.API_TOKEN_PREFIX + accessToken
//	appJson, err := global.Config().RedisClient.GetCtx(*ctx, tokenKey)
//	if err != nil {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	if appJson == "" {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//	}
//	var application model.ApplicationBaseInfoModel
//	err = utils.Json2struct(appJson, &application)
//	if err != nil {
//		return err
//	}
//	t.RenewToken(*ctx, &application, accessToken)
//	*ctx = context.WithValue(*ctx, global.CACHE_API_TOKEN_KEY, accessToken)
//	*ctx = context.WithValue(*ctx, global.CACHE_API_TOKEN_VALUE, appJson)
//	return nil
//}
//
//func (t *apiTokenHandler) RenewToken(ctx context.Context, value interface{}, accessToken string) (bool, error) {
//	app, _ := value.(*model.ApplicationBaseInfoModel)
//	tokenKey := global.API_TOKEN_PREFIX + accessToken
//	appKey := global.TOKEN_API_PREFIX + *app.AppKey
//	err := global.Config().RedisClient.ExpireCtx(ctx, tokenKey, global.TOKE_EXPIRATION_TIME)
//	if err != nil {
//		return false, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	err = global.Config().RedisClient.ExpireCtx(ctx, appKey, global.TOKE_EXPIRATION_TIME)
//	if err != nil {
//		return false, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	return true, nil
//}
//
//func (t *apiTokenHandler) DeleteToken(ctx context.Context, accessToken string) error {
//	tokenKey := global.API_TOKEN_PREFIX + accessToken
//	appJson, err := global.Config().RedisClient.GetCtx(ctx, tokenKey)
//	if err != nil {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	if appJson == "" {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//	}
//	var app model.ApplicationBaseInfoModel
//	err = utils.Json2struct(appJson, &app)
//	if err != nil {
//		return err
//	}
//	appKey := global.TOKEN_API_PREFIX + *app.AppKey
//	_, err = global.Config().RedisClient.DelCtx(ctx, tokenKey, appKey)
//	if err != nil {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	return nil
//}
//
//func (t *apiTokenHandler) RpcCheckToken(ctx context.Context) error {
//	if md, ok := metadata.FromIncomingContext(ctx); ok {
//		if len(md.Get(global.CACHE_API_TOKEN_KEY)) == 0 {
//			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//		} else {
//			accessToken := md.Get(global.CACHE_API_TOKEN_KEY)[0]
//			key := global.API_TOKEN_PREFIX + accessToken
//			if ok, err := global.Config().RedisClient.ExistsCtx(ctx, key); err == nil {
//				if !ok {
//					return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//				}
//			} else {
//				return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.REDIS_ERROR))
//			}
//			return nil
//		}
//	} else {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//	}
//}
//
//func (t *apiTokenHandler) RpcGetToken(ctx context.Context) (string, error) {
//	if md, ok := metadata.FromIncomingContext(ctx); ok {
//		if len(md.Get(global.CACHE_API_TOKEN_KEY)) == 0 {
//			return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//		} else {
//			accessToken := md.Get(global.CACHE_API_TOKEN_KEY)[0]
//			key := global.API_TOKEN_PREFIX + accessToken
//			if ok, err := global.Config().RedisClient.ExistsCtx(ctx, key); err == nil {
//				if !ok {
//					return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//				} else {
//					return accessToken, nil
//				}
//			} else {
//				return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.REDIS_ERROR))
//			}
//		}
//	} else {
//		return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//	}
//}
//
//func (t *apiTokenHandler) RpcGetValue(ctx context.Context) (string, error) {
//	if md, ok := metadata.FromIncomingContext(ctx); ok {
//		if len(md.Get(global.CACHE_API_TOKEN_KEY)) == 0 {
//			return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//		} else {
//			accessToken := md.Get(global.CACHE_API_TOKEN_KEY)[0]
//			key := global.API_TOKEN_PREFIX + accessToken
//			if ok, err := global.Config().RedisClient.ExistsCtx(ctx, key); err == nil {
//				if !ok {
//					return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//				} else {
//					json := md.Get(global.CACHE_API_TOKEN_VALUE)[0]
//					return json, nil
//				}
//			} else {
//				return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.REDIS_ERROR))
//			}
//		}
//	} else {
//		return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//	}
//}
