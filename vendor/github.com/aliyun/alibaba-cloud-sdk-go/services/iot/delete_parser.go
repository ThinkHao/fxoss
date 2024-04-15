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

// DeleteParser invokes the iot.DeleteParser API synchronously
func (client *Client) DeleteParser(request *DeleteParserRequest) (response *DeleteParserResponse, err error) {
	response = CreateDeleteParserResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteParserWithChan invokes the iot.DeleteParser API asynchronously
func (client *Client) DeleteParserWithChan(request *DeleteParserRequest) (<-chan *DeleteParserResponse, <-chan error) {
	responseChan := make(chan *DeleteParserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteParser(request)
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

// DeleteParserWithCallback invokes the iot.DeleteParser API asynchronously
func (client *Client) DeleteParserWithCallback(request *DeleteParserRequest, callback func(response *DeleteParserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteParserResponse
		var err error
		defer close(result)
		response, err = client.DeleteParser(request)
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

// DeleteParserRequest is the request struct for api DeleteParser
type DeleteParserRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ParserId      requests.Integer `position:"Query" name:"ParserId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// DeleteParserResponse is the response struct for api DeleteParser
type DeleteParserResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteParserRequest creates a request to invoke DeleteParser API
func CreateDeleteParserRequest() (request *DeleteParserRequest) {
	request = &DeleteParserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteParser", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteParserResponse creates a response to parse from DeleteParser response
func CreateDeleteParserResponse() (response *DeleteParserResponse) {
	response = &DeleteParserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}