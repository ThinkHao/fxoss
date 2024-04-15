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

// QuerySoundCodeScheduleList invokes the iot.QuerySoundCodeScheduleList API synchronously
func (client *Client) QuerySoundCodeScheduleList(request *QuerySoundCodeScheduleListRequest) (response *QuerySoundCodeScheduleListResponse, err error) {
	response = CreateQuerySoundCodeScheduleListResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySoundCodeScheduleListWithChan invokes the iot.QuerySoundCodeScheduleList API asynchronously
func (client *Client) QuerySoundCodeScheduleListWithChan(request *QuerySoundCodeScheduleListRequest) (<-chan *QuerySoundCodeScheduleListResponse, <-chan error) {
	responseChan := make(chan *QuerySoundCodeScheduleListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySoundCodeScheduleList(request)
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

// QuerySoundCodeScheduleListWithCallback invokes the iot.QuerySoundCodeScheduleList API asynchronously
func (client *Client) QuerySoundCodeScheduleListWithCallback(request *QuerySoundCodeScheduleListRequest, callback func(response *QuerySoundCodeScheduleListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySoundCodeScheduleListResponse
		var err error
		defer close(result)
		response, err = client.QuerySoundCodeScheduleList(request)
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

// QuerySoundCodeScheduleListRequest is the request struct for api QuerySoundCodeScheduleList
type QuerySoundCodeScheduleListRequest struct {
	*requests.RpcRequest
	PageId        requests.Integer `position:"Body" name:"PageId"`
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QuerySoundCodeScheduleListResponse is the response struct for api QuerySoundCodeScheduleList
type QuerySoundCodeScheduleListResponse struct {
	*responses.BaseResponse
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	Success      bool                             `json:"Success" xml:"Success"`
	Code         string                           `json:"Code" xml:"Code"`
	ErrorMessage string                           `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQuerySoundCodeScheduleList `json:"Data" xml:"Data"`
}

// CreateQuerySoundCodeScheduleListRequest creates a request to invoke QuerySoundCodeScheduleList API
func CreateQuerySoundCodeScheduleListRequest() (request *QuerySoundCodeScheduleListRequest) {
	request = &QuerySoundCodeScheduleListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySoundCodeScheduleList", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySoundCodeScheduleListResponse creates a response to parse from QuerySoundCodeScheduleList response
func CreateQuerySoundCodeScheduleListResponse() (response *QuerySoundCodeScheduleListResponse) {
	response = &QuerySoundCodeScheduleListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
