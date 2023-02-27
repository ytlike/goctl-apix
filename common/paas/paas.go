package paas

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/global"
	"qbq-open-platform/common/httpc"
	"qbq-open-platform/common/httpc/header"
)

type Option func(paas *Paas)

type Paas struct {
	ClientId     string
	ClientSecret string
	PaasUrl      string
	Retry        int32
}

func WithClientId(clientId string) Option {
	return func(p *Paas) {
		p.ClientId = clientId
	}
}

func WithClientSecret(clientSecret string) Option {
	return func(p *Paas) {
		p.ClientSecret = clientSecret
	}
}

func WithPaasUrl(paasUrl string) Option {
	return func(p *Paas) {
		p.PaasUrl = paasUrl
	}
}

func WithRetry(retry int32) Option {
	return func(p *Paas) {
		p.Retry = retry
	}
}

func New(opts ...Option) *Paas {
	paas := &Paas{
		ClientId:     "",
		ClientSecret: "",
		PaasUrl:      "",
		Retry:        3,
	}

	for _, opt := range opts {
		opt(paas)
	}

	return paas
}

type (
	PaasData interface {
		GetAuthorization() string
		SetAuthorization(string)
	}

	PaasAuthorization struct {
		Authorization string `header:"Authorization"`
	}

	PaasCommonResult struct {
		Code    int32                  `json:"code"`
		Success bool                   `json:"success"`
		Msg     string                 `json:"msg"`
		Data    map[string]interface{} `json:"data"`
	}

	PaasSpecialResult struct {
		Code    int32  `json:"code"`
		Success bool   `json:"success"`
		Msg     string `json:"msg"`
		Data    string `json:"data"`
	}
)

func (auth *PaasAuthorization) GetAuthorization() string {
	return auth.Authorization
}

func (auth *PaasAuthorization) SetAuthorization(authorization string) {
	auth.Authorization = authorization
}

func (paas *Paas) callPaas(ctx context.Context, method, url string, data PaasData) (map[string]interface{}, error) {
	var paasResp PaasCommonResult
	var err error
	for i := 0; i <= int(paas.Retry); i++ {
		req, err2 := httpc.BuildRequest(ctx, method, url, data)
		if err2 != nil {
			err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.CALL_PAAS_ERROR), "%v", err2)
			continue
		}
		if len(req.Header.Values("Authorization")) == 0 {
			req.Header.Add("Authorization", data.GetAuthorization())
		}

		resp, err2 := httpc.DoRequest(req)
		if err2 != nil {
			err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.CALL_PAAS_ERROR), "%v", err2)
			continue
		}
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.CALL_PAAS_ERROR), "%v", err2)
			continue
		}
		if resp.Header.Get(header.ContentType) == header.StreamContentType {
			paasResp.Data = make(map[string]interface{})
			paasResp.Data["fileData"] = body
			return paasResp.Data, nil
		} else if resp.Header.Get(header.ContentType) == header.ApplicationJson || resp.Header.Get(header.ContentType) == header.JsonContentType {
			err = json.Unmarshal(body, &paasResp)
			if err != nil {
				var paasSpecialResp PaasSpecialResult
				err = json.Unmarshal(body, &paasSpecialResp)
				if err != nil {
					err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.CALL_PAAS_ERROR), "json反序列化失败, %v", err)
					continue
				} else {
					responseData := make(map[string]interface{})
					responseData["data"] = paasSpecialResp.Data
					paasResp = PaasCommonResult{
						Code:    paasSpecialResp.Code,
						Success: paasSpecialResp.Success,
						Msg:     paasSpecialResp.Msg,
						Data:    responseData,
					}
				}
			}

			if resp.StatusCode == http.StatusOK {
				if paasResp.Code != 200 {
					err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.CALL_PAAS_ERROR), "%s", paasResp.Msg)
					continue
				}
			} else {
				err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.CALL_PAAS_ERROR), "%s", paasResp.Msg)
				if resp.StatusCode == http.StatusUnauthorized && paasResp.Code == 40001 {
					_, err3 := global.Config().RedisClient.DelCtx(ctx, global.PAAS_CLIENT_TOKEN)
					if err3 != nil {
						err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err3)
						continue
					}
					paasToken, err4 := paas.GetPaasClientToken(ctx)
					if err4 != nil {
						err = err4
						continue
					}
					data.SetAuthorization(fmt.Sprintf("%s %s", paasToken.TokenType, paasToken.AccessToken))
				}
				continue
			}
			return paasResp.Data, nil
		}
	}
	return nil, err

}
