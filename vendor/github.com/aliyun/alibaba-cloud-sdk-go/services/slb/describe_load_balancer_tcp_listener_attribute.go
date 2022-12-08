package slb

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

// DescribeLoadBalancerTCPListenerAttribute invokes the slb.DescribeLoadBalancerTCPListenerAttribute API synchronously
func (client *Client) DescribeLoadBalancerTCPListenerAttribute(request *DescribeLoadBalancerTCPListenerAttributeRequest) (response *DescribeLoadBalancerTCPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerTCPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerTCPListenerAttributeWithChan invokes the slb.DescribeLoadBalancerTCPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithChan(request *DescribeLoadBalancerTCPListenerAttributeRequest) (<-chan *DescribeLoadBalancerTCPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerTCPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerTCPListenerAttribute(request)
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

// DescribeLoadBalancerTCPListenerAttributeWithCallback invokes the slb.DescribeLoadBalancerTCPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithCallback(request *DescribeLoadBalancerTCPListenerAttributeRequest, callback func(response *DescribeLoadBalancerTCPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerTCPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerTCPListenerAttribute(request)
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

// DescribeLoadBalancerTCPListenerAttributeRequest is the request struct for api DescribeLoadBalancerTCPListenerAttribute
type DescribeLoadBalancerTCPListenerAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerTCPListenerAttributeResponse is the response struct for api DescribeLoadBalancerTCPListenerAttribute
type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	*responses.BaseResponse
	VServerGroupId                 string                                               `json:"VServerGroupId" xml:"VServerGroupId"`
	Status                         string                                               `json:"Status" xml:"Status"`
	AclType                        string                                               `json:"AclType" xml:"AclType"`
	ConnectionDrainTimeout         int                                                  `json:"ConnectionDrainTimeout" xml:"ConnectionDrainTimeout"`
	FailoverStrategy               string                                               `json:"FailoverStrategy" xml:"FailoverStrategy"`
	WorkingServerGroupId           string                                               `json:"WorkingServerGroupId" xml:"WorkingServerGroupId"`
	HealthCheckTcpFastCloseEnabled bool                                                 `json:"HealthCheckTcpFastCloseEnabled" xml:"HealthCheckTcpFastCloseEnabled"`
	FullNatEnabled                 bool                                                 `json:"FullNatEnabled" xml:"FullNatEnabled"`
	ServiceManagedMode             string                                               `json:"ServiceManagedMode" xml:"ServiceManagedMode"`
	RequestId                      string                                               `json:"RequestId" xml:"RequestId"`
	HealthCheckConnectPort         int                                                  `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	Description                    string                                               `json:"Description" xml:"Description"`
	Bandwidth                      int                                                  `json:"Bandwidth" xml:"Bandwidth"`
	HealthCheckType                string                                               `json:"HealthCheckType" xml:"HealthCheckType"`
	MasterSlaveServerGroupId       string                                               `json:"MasterSlaveServerGroupId" xml:"MasterSlaveServerGroupId"`
	BackendServerPort              int                                                  `json:"BackendServerPort" xml:"BackendServerPort"`
	AclStatus                      string                                               `json:"AclStatus" xml:"AclStatus"`
	HealthCheckDomain              string                                               `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	UnhealthyThreshold             int                                                  `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	MasterServerGroupId            string                                               `json:"MasterServerGroupId" xml:"MasterServerGroupId"`
	HealthCheckHttpCode            string                                               `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	MaxConnection                  int                                                  `json:"MaxConnection" xml:"MaxConnection"`
	ProxyProtocolV2Enabled         bool                                                 `json:"ProxyProtocolV2Enabled" xml:"ProxyProtocolV2Enabled"`
	SlaveServerGroupId             string                                               `json:"SlaveServerGroupId" xml:"SlaveServerGroupId"`
	PersistenceTimeout             int                                                  `json:"PersistenceTimeout" xml:"PersistenceTimeout"`
	ListenerPort                   int                                                  `json:"ListenerPort" xml:"ListenerPort"`
	HealthCheckInterval            int                                                  `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckURI                 string                                               `json:"HealthCheckURI" xml:"HealthCheckURI"`
	FailoverThreshold              int                                                  `json:"FailoverThreshold" xml:"FailoverThreshold"`
	AclId                          string                                               `json:"AclId" xml:"AclId"`
	SynProxy                       string                                               `json:"SynProxy" xml:"SynProxy"`
	Scheduler                      string                                               `json:"Scheduler" xml:"Scheduler"`
	EstablishedTimeout             int                                                  `json:"EstablishedTimeout" xml:"EstablishedTimeout"`
	VpcIds                         string                                               `json:"VpcIds" xml:"VpcIds"`
	HealthCheckConnectTimeout      int                                                  `json:"HealthCheckConnectTimeout" xml:"HealthCheckConnectTimeout"`
	MasterSlaveModeEnabled         bool                                                 `json:"MasterSlaveModeEnabled" xml:"MasterSlaveModeEnabled"`
	HealthyThreshold               int                                                  `json:"HealthyThreshold" xml:"HealthyThreshold"`
	ConnectionDrain                string                                               `json:"ConnectionDrain" xml:"ConnectionDrain"`
	HealthCheckMethod              string                                               `json:"HealthCheckMethod" xml:"HealthCheckMethod"`
	HealthCheck                    string                                               `json:"HealthCheck" xml:"HealthCheck"`
	AclIds                         AclIdsInDescribeLoadBalancerTCPListenerAttribute     `json:"AclIds" xml:"AclIds"`
	PortRanges                     PortRangesInDescribeLoadBalancerTCPListenerAttribute `json:"PortRanges" xml:"PortRanges"`
}

// CreateDescribeLoadBalancerTCPListenerAttributeRequest creates a request to invoke DescribeLoadBalancerTCPListenerAttribute API
func CreateDescribeLoadBalancerTCPListenerAttributeRequest() (request *DescribeLoadBalancerTCPListenerAttributeRequest) {
	request = &DescribeLoadBalancerTCPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerTCPListenerAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerTCPListenerAttributeResponse creates a response to parse from DescribeLoadBalancerTCPListenerAttribute response
func CreateDescribeLoadBalancerTCPListenerAttributeResponse() (response *DescribeLoadBalancerTCPListenerAttributeResponse) {
	response = &DescribeLoadBalancerTCPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}