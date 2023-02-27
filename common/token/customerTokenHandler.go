package token

//var (
//	CustomerTokenHandler TokenHandle
//)
//
//func init() {
//	CustomerTokenHandler = NewCustomerTokenHandler()
//}
//
//// customerTokenHandler Customer端token处理类
//type customerTokenHandler struct {
//}
//
//func NewCustomerTokenHandler() *customerTokenHandler {
//	return &customerTokenHandler{}
//}
//
//func (t *customerTokenHandler) ApiCheckToken(ctx *context.Context, accessToken string) error {
//	tokenKey := global.CUSTOMER_TOKEN_PREFIX + accessToken
//	userJson, err := global.Config().RedisClient.GetCtx(*ctx, tokenKey)
//	if err != nil {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	if userJson == "" {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//	}
//	var user model.UserModel
//	err = utils.Json2struct(userJson, &user)
//	if err != nil {
//		return err
//	}
//	t.RenewToken(*ctx, &user, accessToken)
//	*ctx = context.WithValue(*ctx, global.CACHE_CUSTOMER_TOKEN_KEY, accessToken)
//	*ctx = context.WithValue(*ctx, global.CACHE_CUSTOMER_TOKEN_VALUE, userJson)
//	return nil
//}
//
//func (t *customerTokenHandler) RenewToken(ctx context.Context, value interface{}, accessToken string) (bool, error) {
//	user, _ := value.(*model.UserModel)
//	tokenKey := global.CUSTOMER_TOKEN_PREFIX + accessToken
//	userKey := global.TOKEN_CUSTOMER_PREFIX + user.Email
//	err := global.Config().RedisClient.ExpireCtx(ctx, userKey, global.TOKE_EXPIRATION_TIME)
//	if err != nil {
//		return false, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	err = global.Config().RedisClient.ExpireCtx(ctx, tokenKey, global.TOKE_EXPIRATION_TIME)
//	if err != nil {
//		return false, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	return true, nil
//}
//
//func (t *customerTokenHandler) DeleteToken(ctx context.Context, accessToken string) error {
//	tokenKey := global.CUSTOMER_TOKEN_PREFIX + accessToken
//	userJson, err := global.Config().RedisClient.GetCtx(ctx, tokenKey)
//	if err != nil {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	if userJson == "" {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//	}
//	var user model.UserModel
//	err = utils.Json2struct(userJson, &user)
//	if err != nil {
//		return err
//	}
//	userKey := global.TOKEN_CUSTOMER_PREFIX + user.Email
//	_, err = global.Config().RedisClient.DelCtx(ctx, tokenKey, userKey)
//	if err != nil {
//		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
//	}
//	return nil
//}
//
//func (t *customerTokenHandler) RpcCheckToken(ctx context.Context) error {
//	if md, ok := metadata.FromIncomingContext(ctx); ok {
//		if len(md.Get(global.CACHE_CUSTOMER_TOKEN_KEY)) == 0 {
//			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//		} else {
//			accessToken := md.Get(global.CACHE_CUSTOMER_TOKEN_KEY)[0]
//			key := global.CUSTOMER_TOKEN_PREFIX + accessToken
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
//func (t *customerTokenHandler) RpcGetToken(ctx context.Context) (string, error) {
//	if md, ok := metadata.FromIncomingContext(ctx); ok {
//		if len(md.Get(global.CACHE_CUSTOMER_TOKEN_KEY)) == 0 {
//			return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//		} else {
//			accessToken := md.Get(global.CACHE_CUSTOMER_TOKEN_KEY)[0]
//			key := global.CUSTOMER_TOKEN_PREFIX + accessToken
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
//func (t *customerTokenHandler) RpcGetValue(ctx context.Context) (string, error) {
//	if md, ok := metadata.FromIncomingContext(ctx); ok {
//		if len(md.Get(global.CACHE_CUSTOMER_TOKEN_KEY)) == 0 {
//			return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DEFECT_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.DEFECT_TOKEN_ERROR))
//		} else {
//			accessToken := md.Get(global.CACHE_CUSTOMER_TOKEN_KEY)[0]
//			key := global.CUSTOMER_TOKEN_PREFIX + accessToken
//			if ok, err := global.Config().RedisClient.ExistsCtx(ctx, key); err == nil {
//				if !ok {
//					return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INVALID_TOKEN_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.INVALID_TOKEN_ERROR))
//				} else {
//					json := md.Get(global.CACHE_CUSTOMER_TOKEN_VALUE)[0]
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
