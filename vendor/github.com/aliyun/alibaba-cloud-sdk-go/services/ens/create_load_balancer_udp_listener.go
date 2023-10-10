package ens

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

// CreateLoadBalancerUDPListener invokes the ens.CreateLoadBalancerUDPListener API synchronously
func (client *Client) CreateLoadBalancerUDPListener(request *CreateLoadBalancerUDPListenerRequest) (response *CreateLoadBalancerUDPListenerResponse, err error) {
	response = CreateCreateLoadBalancerUDPListenerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerUDPListenerWithChan invokes the ens.CreateLoadBalancerUDPListener API asynchronously
func (client *Client) CreateLoadBalancerUDPListenerWithChan(request *CreateLoadBalancerUDPListenerRequest) (<-chan *CreateLoadBalancerUDPListenerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerUDPListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancerUDPListener(request)
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

// CreateLoadBalancerUDPListenerWithCallback invokes the ens.CreateLoadBalancerUDPListener API asynchronously
func (client *Client) CreateLoadBalancerUDPListenerWithCallback(request *CreateLoadBalancerUDPListenerRequest, callback func(response *CreateLoadBalancerUDPListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerUDPListenerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancerUDPListener(request)
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

// CreateLoadBalancerUDPListenerRequest is the request struct for api CreateLoadBalancerUDPListener
type CreateLoadBalancerUDPListenerRequest struct {
	*requests.RpcRequest
	Protocol                  string           `position:"Query" name:"Protocol"`
	LoadBalancerId            string           `position:"Query" name:"LoadBalancerId"`
	HealthCheckReq            string           `position:"Query" name:"HealthCheckReq"`
	HealthCheckInterval       requests.Integer `position:"Query" name:"HealthCheckInterval"`
	BackendServerPort         requests.Integer `position:"Query" name:"BackendServerPort"`
	HealthCheckExp            string           `position:"Query" name:"HealthCheckExp"`
	HealthCheckConnectTimeout requests.Integer `position:"Query" name:"HealthCheckConnectTimeout"`
	Description               string           `position:"Query" name:"Description"`
	UnhealthyThreshold        requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          requests.Integer `position:"Query" name:"HealthyThreshold"`
	Scheduler                 string           `position:"Query" name:"Scheduler"`
	EipTransmit               string           `position:"Query" name:"EipTransmit"`
	ListenerPort              requests.Integer `position:"Query" name:"ListenerPort"`
	HealthCheckConnectPort    requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
}

// CreateLoadBalancerUDPListenerResponse is the response struct for api CreateLoadBalancerUDPListener
type CreateLoadBalancerUDPListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLoadBalancerUDPListenerRequest creates a request to invoke CreateLoadBalancerUDPListener API
func CreateCreateLoadBalancerUDPListenerRequest() (request *CreateLoadBalancerUDPListenerRequest) {
	request = &CreateLoadBalancerUDPListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateLoadBalancerUDPListener", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLoadBalancerUDPListenerResponse creates a response to parse from CreateLoadBalancerUDPListener response
func CreateCreateLoadBalancerUDPListenerResponse() (response *CreateLoadBalancerUDPListenerResponse) {
	response = &CreateLoadBalancerUDPListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
