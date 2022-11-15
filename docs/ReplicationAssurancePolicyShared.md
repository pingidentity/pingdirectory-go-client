# ReplicationAssurancePolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumreplicationAssurancePolicySchemaUrn**](EnumreplicationAssurancePolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the Replication Assurance Policy. | [optional] 
**Enabled** | **bool** | Indicates whether this Replication Assurance Policy is enabled for use in the server. If a Replication Assurance Policy is disabled, then no new operations will be associated with it. | 
**EvaluationOrderIndex** | **int32** | When multiple Replication Assurance Policies are defined, this property determines the evaluation order for finding a Replication Assurance Policy match against an operation. Policies are evaluated based on this index from least to greatest. Values of this property must be unique but not necessarily contiguous. | 
**LocalLevel** | [**EnumreplicationAssurancePolicyLocalLevelProp**](EnumreplicationAssurancePolicyLocalLevelProp.md) |  | 
**RemoteLevel** | [**EnumreplicationAssurancePolicyRemoteLevelProp**](EnumreplicationAssurancePolicyRemoteLevelProp.md) |  | 
**Timeout** | **string** | Specifies the maximum length of time to wait for the replication assurance requirements to be met before timing out and replying to the client. | 
**ConnectionCriteria** | Pointer to **string** | Specifies a connection criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned. | [optional] 

## Methods

### NewReplicationAssurancePolicyShared

`func NewReplicationAssurancePolicyShared(enabled bool, evaluationOrderIndex int32, localLevel EnumreplicationAssurancePolicyLocalLevelProp, remoteLevel EnumreplicationAssurancePolicyRemoteLevelProp, timeout string, ) *ReplicationAssurancePolicyShared`

NewReplicationAssurancePolicyShared instantiates a new ReplicationAssurancePolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationAssurancePolicySharedWithDefaults

`func NewReplicationAssurancePolicySharedWithDefaults() *ReplicationAssurancePolicyShared`

NewReplicationAssurancePolicySharedWithDefaults instantiates a new ReplicationAssurancePolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ReplicationAssurancePolicyShared) GetSchemas() []EnumreplicationAssurancePolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationAssurancePolicyShared) GetSchemasOk() (*[]EnumreplicationAssurancePolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationAssurancePolicyShared) SetSchemas(v []EnumreplicationAssurancePolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ReplicationAssurancePolicyShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ReplicationAssurancePolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReplicationAssurancePolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReplicationAssurancePolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReplicationAssurancePolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ReplicationAssurancePolicyShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReplicationAssurancePolicyShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReplicationAssurancePolicyShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEvaluationOrderIndex

`func (o *ReplicationAssurancePolicyShared) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *ReplicationAssurancePolicyShared) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *ReplicationAssurancePolicyShared) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetLocalLevel

`func (o *ReplicationAssurancePolicyShared) GetLocalLevel() EnumreplicationAssurancePolicyLocalLevelProp`

GetLocalLevel returns the LocalLevel field if non-nil, zero value otherwise.

### GetLocalLevelOk

`func (o *ReplicationAssurancePolicyShared) GetLocalLevelOk() (*EnumreplicationAssurancePolicyLocalLevelProp, bool)`

GetLocalLevelOk returns a tuple with the LocalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLevel

`func (o *ReplicationAssurancePolicyShared) SetLocalLevel(v EnumreplicationAssurancePolicyLocalLevelProp)`

SetLocalLevel sets LocalLevel field to given value.


### GetRemoteLevel

`func (o *ReplicationAssurancePolicyShared) GetRemoteLevel() EnumreplicationAssurancePolicyRemoteLevelProp`

GetRemoteLevel returns the RemoteLevel field if non-nil, zero value otherwise.

### GetRemoteLevelOk

`func (o *ReplicationAssurancePolicyShared) GetRemoteLevelOk() (*EnumreplicationAssurancePolicyRemoteLevelProp, bool)`

GetRemoteLevelOk returns a tuple with the RemoteLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLevel

`func (o *ReplicationAssurancePolicyShared) SetRemoteLevel(v EnumreplicationAssurancePolicyRemoteLevelProp)`

SetRemoteLevel sets RemoteLevel field to given value.


### GetTimeout

`func (o *ReplicationAssurancePolicyShared) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ReplicationAssurancePolicyShared) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ReplicationAssurancePolicyShared) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetConnectionCriteria

`func (o *ReplicationAssurancePolicyShared) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *ReplicationAssurancePolicyShared) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *ReplicationAssurancePolicyShared) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *ReplicationAssurancePolicyShared) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *ReplicationAssurancePolicyShared) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *ReplicationAssurancePolicyShared) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *ReplicationAssurancePolicyShared) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *ReplicationAssurancePolicyShared) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


