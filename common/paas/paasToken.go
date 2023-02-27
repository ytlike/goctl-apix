package paas

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/global"
	"qbq-open-platform/common/httpc"
	"qbq-open-platform/common/utils"
)

const (
	AOS_CLIENT_CREDENTIALS    = "aos_client_credentials"
	AOS_PASSWORD              = "aos_password"
	AOS_EXCHANGE_TOKEN        = "aos_exchange_token"
	AOS_SIMPLE_EXCHANGE_TOKEN = "aos_simple_exchange_token"
)

type (
	PaasToken struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int32  `json:"expires_in"`
	}

	paasTokenReq struct {
		GrantType string `form:"grant_type"`
	}
)

func (paas *Paas) GetPaasClientToken(ctx context.Context) (*PaasToken, error) {
	paasTokenJson, err := global.Config().RedisClient.GetCtx(ctx, global.PAAS_CLIENT_TOKEN)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
	}
	var paasToken PaasToken
	if paasTokenJson == "" {
		data := &paasTokenReq{
			GrantType: AOS_CLIENT_CREDENTIALS,
		}
		req, err := httpc.BuildRequest(ctx, http.MethodPost, paas.PaasUrl+"/paas-auth/oauth/token", data)
		if err != nil {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "%v", err)
		}
		req.SetBasicAuth(paas.ClientId, paas.ClientSecret)
		resp, err := httpc.DoRequest(req)
		if err != nil {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "%v", err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "%v", err)
		}
		var paasResp PaasCommonResult
		err = json.Unmarshal(body, &paasResp)
		if err != nil {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "json反序列化失败, %v", err)
		}
		if resp.StatusCode == http.StatusOK {
			if paasResp.Code != 200 {
				return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "%s", paasResp.Msg)
			}
		} else {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "%s", paasResp.Msg)
		}

		paasTokenData, err := json.Marshal(paasResp.Data)
		if err != nil {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "json序列化失败, %v", err)
		}
		err = json.Unmarshal(paasTokenData, &paasToken)
		if err != nil {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_TOKEN_ERROR), "json反序列化失败, %v", err)
		}
		paasTokenJson, err = utils.Struct2json(&paasToken)
		if err != nil {
			return nil, err
		}
		err = global.Config().RedisClient.SetexCtx(ctx, global.PAAS_CLIENT_TOKEN, paasTokenJson, int(paasToken.ExpiresIn))
		if err != nil {
			return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
		}
	} else {
		err = utils.Json2struct(paasTokenJson, &paasToken)
		if err != nil {
			return nil, err
		}
	}

	return &paasToken, nil
}
