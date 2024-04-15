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

// QueryEdgeInstanceDevice invokes the iot.QueryEdgeInstanceDevice API synchronously
func (client *Client) QueryEdgeInstanceDevice(request *QueryEdgeInstanceDeviceRequest) (response *QueryEdgeInstanceDeviceResponse, err error) {
	response = CreateQueryEdgeInstanceDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEdgeInstanceDeviceWithChan invokes the iot.QueryEdgeInstanceDevice API asynchronously
func (client *Client) QueryEdgeInstanceDeviceWithChan(request *QueryEdgeInstanceDeviceRequest) (<-chan *QueryEdgeInstanceDeviceResponse, <-chan error) {
	responseChan := make(chan *QueryEdgeInstanceDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEdgeInstanceDevice(request)
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

// QueryEdgeInstanceDeviceWithCallback invokes the iot.QueryEdgeInstanceDevice API asynchronously
func (client *Client) QueryEdgeInstanceDeviceWithCallback(request *QueryEdgeInstanceDeviceRequest, callback func(response *QueryEdgeInstanceDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEdgeInstanceDeviceResponse
		var err error
		defer close(result)
		response, err = client.QueryEdgeInstanceDevice(request)
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

// QueryEdgeInstanceDeviceRequest is the request struct for api QueryEdgeInstanceDevice
type QueryEdgeInstanceDeviceRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryEdgeInstanceDeviceResponse is the response struct for api QueryEdgeInstanceDevice
type QueryEdgeInstanceDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryEdgeInstanceDevice `json:"Data" xml:"Data"`
}

// CreateQueryEdgeInstanceDeviceRequest creates a request to invoke QueryEdgeInstanceDevice API
func CreateQueryEdgeInstanceDeviceRequest() (request *QueryEdgeInstanceDeviceRequest) {
	request = &QueryEdgeInstanceDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryEdgeInstanceDevice", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEdgeInstanceDeviceResponse creates a response to parse from QueryEdgeInstanceDevice response
func CreateQueryEdgeInstanceDeviceResponse() (response *QueryEdgeInstanceDeviceResponse) {
	response = &QueryEdgeInstanceDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
