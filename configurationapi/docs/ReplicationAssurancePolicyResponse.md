# ReplicationAssurancePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Replication Assurance Policy | 
**Schemas** | Pointer to [**[]EnumreplicationAssurancePolicySchemaUrn**](EnumreplicationAssurancePolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the Replication Assurance Policy. | [optional] 
**Enabled** | **bool** | Indicates whether this Replication Assurance Policy is enabled for use in the server. If a Replication Assurance Policy is disabled, then no new operations will be associated with it. | 
**EvaluationOrderIndex** | **int32** | When multiple Replication Assurance Policies are defined, this property determines the evaluation order for finding a Replication Assurance Policy match against an operation. Policies are evaluated based on this index from least to greatest. Values of this property must be unique but not necessarily contiguous. | 
**LocalLevel** | [**EnumreplicationAssurancePolicyLocalLevelProp**](EnumreplicationAssurancePolicyLocalLevelProp.md) |  | 
**RemoteLevel** | [**EnumreplicationAssurancePolicyRemoteLevelProp**](EnumreplicationAssurancePolicyRemoteLevelProp.md) |  | 
**Timeout** | **string** | Specifies the maximum length of time to wait for the replication assurance requirements to be met before timing out and replying to the client. | 
**ConnectionCriteria** | Pointer to **string** | Specifies a connection criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria used to indicate which operations from clients matching this criteria use this policy. If both a connection criteria and a request criteria are specified for a policy, then both must match an operation for the policy to be assigned. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewReplicationAssurancePolicyResponse

`func NewReplicationAssurancePolicyResponse(id string, enabled bool, evaluationOrderIndex int32, localLevel EnumreplicationAssurancePolicyLocalLevelProp, remoteLevel EnumreplicationAssurancePolicyRemoteLevelProp, timeout string, ) *ReplicationAssurancePolicyResponse`

NewReplicationAssurancePolicyResponse instantiates a new ReplicationAssurancePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationAssurancePolicyResponseWithDefaults

`func NewReplicationAssurancePolicyResponseWithDefaults() *ReplicationAssurancePolicyResponse`

NewReplicationAssurancePolicyResponseWithDefaults instantiates a new ReplicationAssurancePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicationAssurancePolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicationAssurancePolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicationAssurancePolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ReplicationAssurancePolicyResponse) GetSchemas() []EnumreplicationAssurancePolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationAssurancePolicyResponse) GetSchemasOk() (*[]EnumreplicationAssurancePolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationAssurancePolicyResponse) SetSchemas(v []EnumreplicationAssurancePolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ReplicationAssurancePolicyResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ReplicationAssurancePolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReplicationAssurancePolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReplicationAssurancePolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReplicationAssurancePolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ReplicationAssurancePolicyResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReplicationAssurancePolicyResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReplicationAssurancePolicyResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEvaluationOrderIndex

`func (o *ReplicationAssurancePolicyResponse) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *ReplicationAssurancePolicyResponse) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *ReplicationAssurancePolicyResponse) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetLocalLevel

`func (o *ReplicationAssurancePolicyResponse) GetLocalLevel() EnumreplicationAssurancePolicyLocalLevelProp`

GetLocalLevel returns the LocalLevel field if non-nil, zero value otherwise.

### GetLocalLevelOk

`func (o *ReplicationAssurancePolicyResponse) GetLocalLevelOk() (*EnumreplicationAssurancePolicyLocalLevelProp, bool)`

GetLocalLevelOk returns a tuple with the LocalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLevel

`func (o *ReplicationAssurancePolicyResponse) SetLocalLevel(v EnumreplicationAssurancePolicyLocalLevelProp)`

SetLocalLevel sets LocalLevel field to given value.


### GetRemoteLevel

`func (o *ReplicationAssurancePolicyResponse) GetRemoteLevel() EnumreplicationAssurancePolicyRemoteLevelProp`

GetRemoteLevel returns the RemoteLevel field if non-nil, zero value otherwise.

### GetRemoteLevelOk

`func (o *ReplicationAssurancePolicyResponse) GetRemoteLevelOk() (*EnumreplicationAssurancePolicyRemoteLevelProp, bool)`

GetRemoteLevelOk returns a tuple with the RemoteLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLevel

`func (o *ReplicationAssurancePolicyResponse) SetRemoteLevel(v EnumreplicationAssurancePolicyRemoteLevelProp)`

SetRemoteLevel sets RemoteLevel field to given value.


### GetTimeout

`func (o *ReplicationAssurancePolicyResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ReplicationAssurancePolicyResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ReplicationAssurancePolicyResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### GetConnectionCriteria

`func (o *ReplicationAssurancePolicyResponse) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *ReplicationAssurancePolicyResponse) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *ReplicationAssurancePolicyResponse) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *ReplicationAssurancePolicyResponse) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *ReplicationAssurancePolicyResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *ReplicationAssurancePolicyResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *ReplicationAssurancePolicyResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *ReplicationAssurancePolicyResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetMeta

`func (o *ReplicationAssurancePolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReplicationAssurancePolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReplicationAssurancePolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ReplicationAssurancePolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationAssurancePolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ReplicationAssurancePolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationAssurancePolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ReplicationAssurancePolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


