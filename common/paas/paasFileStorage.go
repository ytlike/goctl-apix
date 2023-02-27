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
	PaasFileMetadataReq struct {
		PaasAuthorization
		Id string `path:"id"`
	}

	PaasFileDownReq struct {
		PaasAuthorization
		Id string `path:"id"`
	}

	PaasFileUrlReq struct {
		PaasAuthorization
		Id string `path:"id"`
	}
)

type (
	PaasFileMetadata struct {
		FileName     string `json:"name"`
		FileFullName string `json:"fname"`
		FileExt      string `json:"ext"`
		FileSize     int64  `json:"size"`
		FileMD5      string `json:"md5"`
		FileSM3      string `json:"sm3"`
		FileUrl      string `json:"url"`
	}

	PaasFileUrl struct {
		Data string `json:"data"`
	}
)

func (paas *Paas) GetPaasFileMetadata(ctx context.Context, fileId string) (*PaasFileMetadata, error) {
	paasToken, err := paas.GetPaasClientToken(ctx)
	if err != nil {
		return nil, err
	}

	var paasFileMetadata PaasFileMetadata
	authorization := fmt.Sprintf("%s %s", paasToken.TokenType, paasToken.AccessToken)
	data := &PaasFileMetadataReq{
		PaasAuthorization: PaasAuthorization{Authorization: authorization},
		Id:                fileId,
	}

	respData, err := paas.callPaas(ctx, http.MethodGet, paas.PaasUrl+"/paas-file-storage/metadata/files/:id/info", data)
	if err != nil {
		return nil, err
	}
	paasFileMetadataData, err := json.Marshal(respData)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_FILE_METADATA_ERROR), "%v", err)
	}
	err = json.Unmarshal(paasFileMetadataData, &paasFileMetadata)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_FILE_METADATA_ERROR), "%v", err)
	}
	return &paasFileMetadata, nil
}

func (paas *Paas) PaasFileDown(ctx context.Context, fileId string) ([]byte, error) {
	paasToken, err := paas.GetPaasClientToken(ctx)
	if err != nil {
		return nil, err
	}
	authorization := fmt.Sprintf("%s %s", paasToken.TokenType, paasToken.AccessToken)
	data := &PaasFileDownReq{
		PaasAuthorization: PaasAuthorization{Authorization: authorization},
		Id:                fileId,
	}
	paasFileDownData, err := paas.callPaas(ctx, http.MethodGet, paas.PaasUrl+"/paas-file-storage/metadata/files/:id", data)
	if err != nil {
		return nil, err
	}
	fileData, ok := paasFileDownData["fileData"].([]byte)
	if ok {
		return fileData, nil
	} else {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
}

func (paas *Paas) GetPaasFileUrl(ctx context.Context, fileId string) (*PaasFileUrl, error) {
	paasToken, err := paas.GetPaasClientToken(ctx)
	if err != nil {
		return nil, err
	}
	var paasFileUrl PaasFileUrl
	authorization := fmt.Sprintf("%s %s", paasToken.TokenType, paasToken.AccessToken)
	data := &PaasFileUrlReq{
		PaasAuthorization: PaasAuthorization{Authorization: authorization},
		Id:                fileId,
	}

	respData, err := paas.callPaas(ctx, http.MethodGet, paas.PaasUrl+"/paas-file-storage/metadata/files/:id/url", data)
	if err != nil {
		return nil, err
	}
	paasUrlData, err := json.Marshal(respData)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_FILE_URL_ERROR), "%v", err)
	}
	err = json.Unmarshal(paasUrlData, &paasFileUrl)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.GET_PAAS_FILE_URL_ERROR), "%v", err)
	}
	return &paasFileUrl, nil
}
