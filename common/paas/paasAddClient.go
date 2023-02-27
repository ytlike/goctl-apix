package paas

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"qbq-open-platform/common/errorsEnums"
)

type (
	ClientAddReq struct {
		PaasAuthorization
		AccessTokenValidity  int32    `json:"accessTokenValidity"  validate:"required"`
		AppId                string   `json:"appId" validate:"required"`
		ClientName           string   `json:"clientName" validate:"required"`
		GrantTypes           []string `json:"grantTypes"  validate:"required"`
		PId                  int32    `json:"pid"  validate:"required"`
		RefreshTokenValidity int32    `json:"refreshTokenValidity"  validate:"required"`
		ResourceIds          []string `json:"resourceIds"  validate:"required"`
	}
)

type (
	ClientAddRespData struct {
		AppId        string `JSON:"appId"`
		ClientId     string `JSON:"clientId"`
		ClientName   string `JSON:"clientName"`
		ClientSecret string `JSON:"clientSecret"`
		TenantId     string `JSON:"tenantId"`
	}
)

func (paas *Paas) PaasAddClient(ctx context.Context, clientAddReq *ClientAddReq) (*ClientAddRespData, error) {
	paasToken, err := paas.GetPaasClientToken(ctx)
	if err != nil {
		return nil, err
	}
	var clientAddResp ClientAddRespData
	authorization := fmt.Sprintf("%s %s", paasToken.TokenType, paasToken.AccessToken)
	clientAddReq.SetAuthorization(authorization)
	respData, err := paas.callPaas(ctx, http.MethodPost, paas.PaasUrl+"/paas-permission/client/permission/addClient", clientAddReq)
	if err != nil {
		return nil, err
	}
	addClientData, err := json.Marshal(respData)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.ADD_CLIENT_ERROR), "%v", err)
	}
	err = json.Unmarshal(addClientData, &clientAddResp)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.ADD_CLIENT_ERROR), "%v", err)
	}
	return &clientAddResp, nil
}
