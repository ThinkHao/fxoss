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

// QueryPageByApplyId invokes the iot.QueryPageByApplyId API synchronously
func (client *Client) QueryPageByApplyId(request *QueryPageByApplyIdRequest) (response *QueryPageByApplyIdResponse, err error) {
	response = CreateQueryPageByApplyIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPageByApplyIdWithChan invokes the iot.QueryPageByApplyId API asynchronously
func (client *Client) QueryPageByApplyIdWithChan(request *QueryPageByApplyIdRequest) (<-chan *QueryPageByApplyIdResponse, <-chan error) {
	responseChan := make(chan *QueryPageByApplyIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPageByApplyId(request)
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

// QueryPageByApplyIdWithCallback invokes the iot.QueryPageByApplyId API asynchronously
func (client *Client) QueryPageByApplyIdWithCallback(request *QueryPageByApplyIdRequest, callback func(response *QueryPageByApplyIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPageByApplyIdResponse
		var err error
		defer close(result)
		response, err = client.QueryPageByApplyId(request)
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

// QueryPageByApplyIdRequest is the request struct for api QueryPageByApplyId
type QueryPageByApplyIdRequest struct {
	*requests.RpcRequest
	RealTenantId      string           `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string           `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage       requests.Integer `position:"Query" name:"CurrentPage"`
	ApplyId           requests.Integer `position:"Query" name:"ApplyId"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
}

// QueryPageByApplyIdResponse is the response struct for api QueryPageByApplyId
type QueryPageByApplyIdResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Success         bool            `json:"Success" xml:"Success"`
	Code            string          `json:"Code" xml:"Code"`
	ErrorMessage    string          `json:"ErrorMessage" xml:"ErrorMessage"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	Page            int             `json:"Page" xml:"Page"`
	PageCount       int             `json:"PageCount" xml:"PageCount"`
	Total           int             `json:"Total" xml:"Total"`
	ApplyDeviceList ApplyDeviceList `json:"ApplyDeviceList" xml:"ApplyDeviceList"`
}

// CreateQueryPageByApplyIdRequest creates a request to invoke QueryPageByApplyId API
func CreateQueryPageByApplyIdRequest() (request *QueryPageByApplyIdRequest) {
	request = &QueryPageByApplyIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryPageByApplyId", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryPageByApplyIdResponse creates a response to parse from QueryPageByApplyId response
func CreateQueryPageByApplyIdResponse() (response *QueryPageByApplyIdResponse) {
	response = &QueryPageByApplyIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}