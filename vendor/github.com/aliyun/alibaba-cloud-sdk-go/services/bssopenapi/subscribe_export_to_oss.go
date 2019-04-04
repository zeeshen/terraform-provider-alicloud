package bssopenapi

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

// SubscribeExportToOSS invokes the bssopenapi.SubscribeExportToOSS API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/subscribeexporttooss.html
func (client *Client) SubscribeExportToOSS(request *SubscribeExportToOSSRequest) (response *SubscribeExportToOSSResponse, err error) {
	response = CreateSubscribeExportToOSSResponse()
	err = client.DoAction(request, response)
	return
}

// SubscribeExportToOSSWithChan invokes the bssopenapi.SubscribeExportToOSS API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/subscribeexporttooss.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubscribeExportToOSSWithChan(request *SubscribeExportToOSSRequest) (<-chan *SubscribeExportToOSSResponse, <-chan error) {
	responseChan := make(chan *SubscribeExportToOSSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubscribeExportToOSS(request)
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

// SubscribeExportToOSSWithCallback invokes the bssopenapi.SubscribeExportToOSS API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/subscribeexporttooss.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubscribeExportToOSSWithCallback(request *SubscribeExportToOSSRequest, callback func(response *SubscribeExportToOSSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubscribeExportToOSSResponse
		var err error
		defer close(result)
		response, err = client.SubscribeExportToOSS(request)
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

// SubscribeExportToOSSRequest is the request struct for api SubscribeExportToOSS
type SubscribeExportToOSSRequest struct {
	*requests.RpcRequest
	BucketOwnerId   requests.Integer `position:"Query" name:"BucketOwnerId"`
	SubscribeType   *[]string        `position:"Query" name:"SubscribeType"  type:"Repeated"`
	SubscribeBucket string           `position:"Query" name:"SubscribeBucket"`
}

// SubscribeExportToOSSResponse is the response struct for api SubscribeExportToOSS
type SubscribeExportToOSSResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSubscribeExportToOSSRequest creates a request to invoke SubscribeExportToOSS API
func CreateSubscribeExportToOSSRequest() (request *SubscribeExportToOSSRequest) {
	request = &SubscribeExportToOSSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "SubscribeExportToOSS", "", "")
	return
}

// CreateSubscribeExportToOSSResponse creates a response to parse from SubscribeExportToOSS response
func CreateSubscribeExportToOSSResponse() (response *SubscribeExportToOSSResponse) {
	response = &SubscribeExportToOSSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}