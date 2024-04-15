package utils

import (
	iot "github.com/alibabacloud-go/iot-api-gateway/client"
)

//func CreateAliIOTClient(accessKeyId *string, accessKeySecret *string) (_result *iot20180120.Client, _err error) {
//	config := &openapi.Config{
//		AccessKeyId:     accessKeyId,
//		AccessKeySecret: accessKeySecret,
//	}
//	config.Endpoint = tea.String("iot.cn-shanghai.aliyuncs.com")
//	_result = &iot20180120.Client{}
//	_result, _err = iot20180120.NewClient(config)
//
//	return _result, _err
//}
//
//func CreateAliIOTClientLegacy(accessKeyId, accessKeySecret string) (*legacySDK.Client, error) {
//	client, err := legacySDK.NewClientWithAccessKey("cn-shanghai", accessKeyId, accessKeySecret)
//	if err != nil {
//		return &legacySDK.Client{}, err
//	}
//
//	return client, nil
//}

func CreateAliIotApiClient(keyId, secretKey string) (*iot.Client, error) {
	config := new(iot.Config).
		SetAppKey(keyId).
		SetAppSecret(secretKey).
		SetDomain("api.link.aliyun.com")

	client, err := iot.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
