package iot

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyOTAFirmware invokes the iot.ModifyOTAFirmware API synchronously
func (client *Client) ModifyOTAFirmware(request *ModifyOTAFirmwareRequest) (response *ModifyOTAFirmwareResponse, err error) {
	response = CreateModifyOTAFirmwareResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyOTAFirmwareWithChan invokes the iot.ModifyOTAFirmware API asynchronously
func (client *Client) ModifyOTAFirmwareWithChan(request *ModifyOTAFirmwareRequest) (<-chan *ModifyOTAFirmwareResponse, <-chan error) {
	responseChan := make(chan *ModifyOTAFirmwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyOTAFirmware(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyOTAFirmwareWithCallback invokes the iot.ModifyOTAFirmware API asynchronously
func (client *Client) ModifyOTAFirmwareWithCallback(request *ModifyOTAFirmwareRequest, callback func(response *ModifyOTAFirmwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyOTAFirmwareResponse
		var err error
		defer close(result)
		response, err = client.ModifyOTAFirmware(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyOTAFirmwareRequest is the request struct for api ModifyOTAFirmware
type ModifyOTAFirmwareRequest struct {
	*requests.RpcRequest
	FirmwareUdi   string `position:"Query" name:"FirmwareUdi"`
	FirmwareDesc  string `position:"Query" name:"FirmwareDesc"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	FirmwareName  string `position:"Query" name:"FirmwareName"`
	FirmwareId    string `position:"Query" name:"FirmwareId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// ModifyOTAFirmwareResponse is the response struct for api ModifyOTAFirmware
type ModifyOTAFirmwareResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateModifyOTAFirmwareRequest creates a request to invoke ModifyOTAFirmware API
func CreateModifyOTAFirmwareRequest() (request *ModifyOTAFirmwareRequest) {
	request = &ModifyOTAFirmwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ModifyOTAFirmware", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyOTAFirmwareResponse creates a response to parse from ModifyOTAFirmware response
func CreateModifyOTAFirmwareResponse() (response *ModifyOTAFirmwareResponse) {
	response = &ModifyOTAFirmwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}