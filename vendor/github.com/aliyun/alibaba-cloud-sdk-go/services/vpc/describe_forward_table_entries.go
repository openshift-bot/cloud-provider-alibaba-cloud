package vpc

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

// DescribeForwardTableEntries invokes the vpc.DescribeForwardTableEntries API synchronously
func (client *Client) DescribeForwardTableEntries(request *DescribeForwardTableEntriesRequest) (response *DescribeForwardTableEntriesResponse, err error) {
	response = CreateDescribeForwardTableEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeForwardTableEntriesWithChan invokes the vpc.DescribeForwardTableEntries API asynchronously
func (client *Client) DescribeForwardTableEntriesWithChan(request *DescribeForwardTableEntriesRequest) (<-chan *DescribeForwardTableEntriesResponse, <-chan error) {
	responseChan := make(chan *DescribeForwardTableEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeForwardTableEntries(request)
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

// DescribeForwardTableEntriesWithCallback invokes the vpc.DescribeForwardTableEntries API asynchronously
func (client *Client) DescribeForwardTableEntriesWithCallback(request *DescribeForwardTableEntriesRequest, callback func(response *DescribeForwardTableEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeForwardTableEntriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeForwardTableEntries(request)
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

// DescribeForwardTableEntriesRequest is the request struct for api DescribeForwardTableEntries
type DescribeForwardTableEntriesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ForwardTableId       string           `position:"Query" name:"ForwardTableId"`
	InternalIp           string           `position:"Query" name:"InternalIp"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ForwardEntryId       string           `position:"Query" name:"ForwardEntryId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
	ExternalIp           string           `position:"Query" name:"ExternalIp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string           `position:"Query" name:"IpProtocol"`
	ForwardEntryName     string           `position:"Query" name:"ForwardEntryName"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InternalPort         string           `position:"Query" name:"InternalPort"`
	ExternalPort         string           `position:"Query" name:"ExternalPort"`
}

// DescribeForwardTableEntriesResponse is the response struct for api DescribeForwardTableEntries
type DescribeForwardTableEntriesResponse struct {
	*responses.BaseResponse
	PageSize            int                 `json:"PageSize" xml:"PageSize"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	PageNumber          int                 `json:"PageNumber" xml:"PageNumber"`
	TotalCount          int                 `json:"TotalCount" xml:"TotalCount"`
	ForwardTableEntries ForwardTableEntries `json:"ForwardTableEntries" xml:"ForwardTableEntries"`
}

// CreateDescribeForwardTableEntriesRequest creates a request to invoke DescribeForwardTableEntries API
func CreateDescribeForwardTableEntriesRequest() (request *DescribeForwardTableEntriesRequest) {
	request = &DescribeForwardTableEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeForwardTableEntries", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeForwardTableEntriesResponse creates a response to parse from DescribeForwardTableEntries response
func CreateDescribeForwardTableEntriesResponse() (response *DescribeForwardTableEntriesResponse) {
	response = &DescribeForwardTableEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
