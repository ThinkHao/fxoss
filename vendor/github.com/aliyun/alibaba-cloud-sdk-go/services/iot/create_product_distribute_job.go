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

// CreateProductDistributeJob invokes the iot.CreateProductDistributeJob API synchronously
func (client *Client) CreateProductDistributeJob(request *CreateProductDistributeJobRequest) (response *CreateProductDistributeJobResponse, err error) {
	response = CreateCreateProductDistributeJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProductDistributeJobWithChan invokes the iot.CreateProductDistributeJob API asynchronously
func (client *Client) CreateProductDistributeJobWithChan(request *CreateProductDistributeJobRequest) (<-chan *CreateProductDistributeJobResponse, <-chan error) {
	responseChan := make(chan *CreateProductDistributeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProductDistributeJob(request)
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

// CreateProductDistributeJobWithCallback invokes the iot.CreateProductDistributeJob API asynchronously
func (client *Client) CreateProductDistributeJobWithCallback(request *CreateProductDistributeJobRequest, callback func(response *CreateProductDistributeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProductDistributeJobResponse
		var err error
		defer close(result)
		response, err = client.CreateProductDistributeJob(request)
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

// CreateProductDistributeJobRequest is the request struct for api CreateProductDistributeJob
type CreateProductDistributeJobRequest struct {
	*requests.RpcRequest
	Captcha          string `position:"Query" name:"Captcha"`
	SourceInstanceId string `position:"Query" name:"SourceInstanceId"`
	TargetAliyunId   string `position:"Query" name:"TargetAliyunId"`
	ProductKey       string `position:"Query" name:"ProductKey"`
	TargetInstanceId string `position:"Query" name:"TargetInstanceId"`
	ApiProduct       string `position:"Body" name:"ApiProduct"`
	ApiRevision      string `position:"Body" name:"ApiRevision"`
	TargetUid        string `position:"Query" name:"TargetUid"`
}

// CreateProductDistributeJobResponse is the response struct for api CreateProductDistributeJob
type CreateProductDistributeJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	JobId        string `json:"JobId" xml:"JobId"`
}

// CreateCreateProductDistributeJobRequest creates a request to invoke CreateProductDistributeJob API
func CreateCreateProductDistributeJobRequest() (request *CreateProductDistributeJobRequest) {
	request = &CreateProductDistributeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateProductDistributeJob", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProductDistributeJobResponse creates a response to parse from CreateProductDistributeJob response
func CreateCreateProductDistributeJobResponse() (response *CreateProductDistributeJobResponse) {
	response = &CreateProductDistributeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
