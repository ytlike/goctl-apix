package vo

type ApplicationMessageVo struct {
	RequestId         string           `json:"requestId"`
	TaskCode          string           `json:"taskCode"`
	PublishAppCode    string           `json:"publishAppCode"`
	OperateType       int32            `json:"operateType"`
	AcceptAppCodeList []string         `json:"acceptAppCodeList"`
	OperateUserInfo   *OperateUserInfo `json:"operateUserInfo"`
	AddTaskInfo       *AddTaskInfo     `json:"addTaskInfo,omitempty"`
	ExpiredInfo       *ExpiredInfo     `json:"ExpiredInfo,omitempty"`
	ChangeTaskInfo    *ChangeTaskInfo  `json:"changeTaskInfo,omitempty"`
	AddSignerInfo     *AddSignerInfo   `json:"addSignerInfo,omitempty"`
}

type OperateUserInfo struct {
	SignerName string `json:"signerName"`
	SignerId   string `json:"signerId"`
}

type AddTaskInfo struct {
	TaskName                string           `json:"taskName"`
	TaskCode                string           `json:"taskCode"`
	TaskBid                 string           `json:"taskBid"`
	TaskFlowId              *string          `json:"taskFlowId"`
	TaskType                int32            `json:"taskType"`
	TaskCategory            int32            `json:"taskCategory"`
	TaskVersion             string           `json:"taskVersion"`
	SignType                int32            `json:"signType"`
	PublisherId             string           `json:"publisherId"`
	PublisherName           string           `json:"publisherName"`
	PublisherTel            *string          `json:"publisherTel"`
	PublisherCardType       *string          `json:"publisherCardType"`
	PublisherCardNo         *string          `json:"publisherCardNo"`
	PublisherEnterpriseCode *string          `json:"publisherEnterpriseCode"`
	PublisherEnterpriseName *string          `json:"publisherEnterpriseName"`
	PublishTime             string           `json:"publishTime"`
	Deadline                string           `json:"deadline"`
	CreateTime              string           `json:"createTime"`
	SignatoryList           []*Signatory     `json:"signatoryList"`
	SignFileList            []*SignFile      `json:"signFileList"`
	SignAnnexFileList       []*SignAnnexFile `json:"signAnnexFileList"`
}

type ExpiredInfo struct {
	TaskCode string `json:"taskCode"`
}

type Signatory struct {
	SignatoryId             string        `json:"signatoryId"`
	SignatoryName           string        `json:"signatoryName"`
	SignatoryEnterpriseName string        `json:"signatoryEnterpriseName"`
	SignatoryEnterpriseCode string        `json:"signatoryEnterpriseCode"`
	SignerInfoList          []*SignerInfo `json:"signerInfoList"`
}

type SignerInfo struct {
	SignerId       string `json:"signerId"`
	SignerPaasId   string `json:"signerPaasId"`
	SignerTel      string `json:"signerTel"`
	SignerName     string `json:"signerName"`
	SignerCardType string `json:"signerCardType"`
	SignerCardNo   string `json:"signerCardNo"`
}
type SignFile struct {
	FileId        string         `json:"fileId"`
	FileName      string         `json:"fileName"`
	FileSize      int64          `json:"fileSize"`
	SignJson      string         `json:"signJson"`
	FileExt       string         `json:"fileExt"`
	SignatoryList []*SignatoryId `json:"signatoryList"`
}

type SignatoryId struct {
	SignatoryId string `json:"signatoryId"`
}

type SignAnnexFile struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	FileExt  string `json:"fileExt"`
}
type ChangeTaskInfo struct {
	TaskCode        string            `json:"taskCode"`
	TaskStatus      int32             `json:"taskStatus"`
	SignerId        string            `json:"signerId"`
	SignerStatus    int32             `json:"signerStatus"`
	SignatoryId     string            `json:"signatoryId"`
	SignatoryStatus int32             `json:"signatoryStatus"`
	Reason          string            `json:"reason"`
	OperateTime     string            `json:"operateTime"`
	SignFileList    []*ChangeSignFile `json:"signFileList"`
	FinishTime      string            `json:"finishTime"`
}

type ChangeSignFile struct {
	FileId         string `json:"fileId"`
	SignedFileId   string `json:"signedFileId"`
	FileStatus     int32  `json:"fileStatus"`
	SignedFileSize int32  `json:"signedFileSize"`
}
type AddSignerInfo struct {
	TaskCode    string       `json:"taskCode"`
	SignatoryId string       `json:"signatoryId"`
	SignerInfo  []*AddSigner `json:"signerInfo"`
}

type AddSigner struct {
	SignerId       string `json:"signerId"`
	SignerPaasId   string `json:"signerPaasId"`
	SignerTel      string `json:"signerTel"`
	SignerName     string `json:"signerName"`
	SignerCardType string `json:"signerCardType"`
	SignerCardNo   string `json:"signerCardNo"`
}
