package server_api_params

//type TencentCloudStorageCredentialReq struct {
//	OperationID string `json:"operationID"`
//}
//
//type TencentCloudStorageCredentialRespData struct {
//	*sts.CredentialResult
//	Region string `json:"region"`
//	Bucket string `json:"bucket"`
//}
//
//type TencentCloudStorageCredentialResp struct {
//	CommResp
//	CosData TencentCloudStorageCredentialRespData `json:"-"`
//	Data    map[string]interface{}                `json:"data"`
//}
type FcmUpdateTokenReq struct {
	OperationID string `json:"operationID" binding:"required"`
	Platform    int    `json:"platform" binding:"required,min=1,max=2"` //only for ios + android
	FcmToken    string `json:"fcmToken" binding:"required"`
}

type FcmUpdateTokenResp struct {
	CommResp
}
