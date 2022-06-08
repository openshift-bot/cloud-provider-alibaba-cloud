package ess

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

// CreateScalingGroup invokes the ess.CreateScalingGroup API synchronously
func (client *Client) CreateScalingGroup(request *CreateScalingGroupRequest) (response *CreateScalingGroupResponse, err error) {
	response = CreateCreateScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScalingGroupWithChan invokes the ess.CreateScalingGroup API asynchronously
func (client *Client) CreateScalingGroupWithChan(request *CreateScalingGroupRequest) (<-chan *CreateScalingGroupResponse, <-chan error) {
	responseChan := make(chan *CreateScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScalingGroup(request)
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

// CreateScalingGroupWithCallback invokes the ess.CreateScalingGroup API asynchronously
func (client *Client) CreateScalingGroupWithCallback(request *CreateScalingGroupRequest, callback func(response *CreateScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateScalingGroup(request)
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

// CreateScalingGroupRequest is the request struct for api CreateScalingGroup
type CreateScalingGroupRequest struct {
	*requests.RpcRequest
	VSwitchIds                          *[]string                                   `position:"Query" name:"VSwitchIds"  type:"Repeated"`
	SpotInstanceRemedy                  requests.Boolean                            `position:"Query" name:"SpotInstanceRemedy"`
	ScaleOutAmountCheck                 requests.Boolean                            `position:"Query" name:"ScaleOutAmountCheck"`
	GroupType                           string                                      `position:"Query" name:"GroupType"`
	Tag                                 *[]CreateScalingGroupTag                    `position:"Query" name:"Tag"  type:"Repeated"`
	DefaultCooldown                     requests.Integer                            `position:"Query" name:"DefaultCooldown"`
	ContainerGroupId                    string                                      `position:"Query" name:"ContainerGroupId"`
	MultiAZPolicy                       string                                      `position:"Query" name:"MultiAZPolicy"`
	DBInstanceIds                       string                                      `position:"Query" name:"DBInstanceIds"`
	LaunchTemplateId                    string                                      `position:"Query" name:"LaunchTemplateId"`
	DesiredCapacity                     requests.Integer                            `position:"Query" name:"DesiredCapacity"`
	LaunchTemplateOverride              *[]CreateScalingGroupLaunchTemplateOverride `position:"Query" name:"LaunchTemplateOverride"  type:"Repeated"`
	CompensateWithOnDemand              requests.Boolean                            `position:"Query" name:"CompensateWithOnDemand"`
	MinSize                             requests.Integer                            `position:"Query" name:"MinSize"`
	OwnerId                             requests.Integer                            `position:"Query" name:"OwnerId"`
	AlbServerGroup                      *[]CreateScalingGroupAlbServerGroup         `position:"Query" name:"AlbServerGroup"  type:"Repeated"`
	VSwitchId                           string                                      `position:"Query" name:"VSwitchId"`
	InstanceId                          string                                      `position:"Query" name:"InstanceId"`
	MaxSize                             requests.Integer                            `position:"Query" name:"MaxSize"`
	LifecycleHook                       *[]CreateScalingGroupLifecycleHook          `position:"Query" name:"LifecycleHook"  type:"Repeated"`
	LoadBalancerIds                     string                                      `position:"Query" name:"LoadBalancerIds"`
	ClientToken                         string                                      `position:"Query" name:"ClientToken"`
	OnDemandBaseCapacity                requests.Integer                            `position:"Query" name:"OnDemandBaseCapacity"`
	OnDemandPercentageAboveBaseCapacity requests.Integer                            `position:"Query" name:"OnDemandPercentageAboveBaseCapacity"`
	RemovalPolicy1                      string                                      `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2                      string                                      `position:"Query" name:"RemovalPolicy.2"`
	HealthCheckType                     string                                      `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount                string                                      `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName                    string                                      `position:"Query" name:"ScalingGroupName"`
	OwnerAccount                        string                                      `position:"Query" name:"OwnerAccount"`
	SpotInstancePools                   requests.Integer                            `position:"Query" name:"SpotInstancePools"`
	GroupDeletionProtection             requests.Boolean                            `position:"Query" name:"GroupDeletionProtection"`
	LaunchTemplateVersion               string                                      `position:"Query" name:"LaunchTemplateVersion"`
	ScalingPolicy                       string                                      `position:"Query" name:"ScalingPolicy"`
	VServerGroup                        *[]CreateScalingGroupVServerGroup           `position:"Query" name:"VServerGroup"  type:"Repeated"`
}

// CreateScalingGroupTag is a repeated param struct in CreateScalingGroupRequest
type CreateScalingGroupTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateScalingGroupLaunchTemplateOverride is a repeated param struct in CreateScalingGroupRequest
type CreateScalingGroupLaunchTemplateOverride struct {
	WeightedCapacity string `name:"WeightedCapacity"`
	InstanceType     string `name:"InstanceType"`
}

// CreateScalingGroupAlbServerGroup is a repeated param struct in CreateScalingGroupRequest
type CreateScalingGroupAlbServerGroup struct {
	AlbServerGroupId string `name:"AlbServerGroupId"`
	Port             string `name:"Port"`
	Weight           string `name:"Weight"`
}

// CreateScalingGroupLifecycleHook is a repeated param struct in CreateScalingGroupRequest
type CreateScalingGroupLifecycleHook struct {
	DefaultResult        string `name:"DefaultResult"`
	LifecycleHookName    string `name:"LifecycleHookName"`
	HeartbeatTimeout     string `name:"HeartbeatTimeout"`
	NotificationArn      string `name:"NotificationArn"`
	NotificationMetadata string `name:"NotificationMetadata"`
	LifecycleTransition  string `name:"LifecycleTransition"`
}

// CreateScalingGroupVServerGroup is a repeated param struct in CreateScalingGroupRequest
type CreateScalingGroupVServerGroup struct {
	LoadBalancerId        string                                                 `name:"LoadBalancerId"`
	VServerGroupAttribute *[]CreateScalingGroupVServerGroupVServerGroupAttribute `name:"VServerGroupAttribute" type:"Repeated"`
}

// CreateScalingGroupVServerGroupVServerGroupAttribute is a repeated param struct in CreateScalingGroupRequest
type CreateScalingGroupVServerGroupVServerGroupAttribute struct {
	VServerGroupId string `name:"VServerGroupId"`
	Port           string `name:"Port"`
	Weight         string `name:"Weight"`
}

// CreateScalingGroupResponse is the response struct for api CreateScalingGroup
type CreateScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ScalingGroupId string `json:"ScalingGroupId" xml:"ScalingGroupId"`
}

// CreateCreateScalingGroupRequest creates a request to invoke CreateScalingGroup API
func CreateCreateScalingGroupRequest() (request *CreateScalingGroupRequest) {
	request = &CreateScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingGroup", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateScalingGroupResponse creates a response to parse from CreateScalingGroup response
func CreateCreateScalingGroupResponse() (response *CreateScalingGroupResponse) {
	response = &CreateScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
