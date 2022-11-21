# ReplicationAssuranceResultCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Result Criteria | 
**Schemas** | [**[]EnumreplicationAssuranceResultCriteriaSchemaUrn**](EnumreplicationAssuranceResultCriteriaSchemaUrn.md) |  | 
**LocalAssuranceLevel** | [**[]EnumresultCriteriaLocalAssuranceLevelProp**](EnumresultCriteriaLocalAssuranceLevelProp.md) |  | 
**RemoteAssuranceLevel** | [**[]EnumresultCriteriaRemoteAssuranceLevelProp**](EnumresultCriteriaRemoteAssuranceLevelProp.md) |  | 
**AssuranceTimeoutCriteria** | Pointer to [**EnumresultCriteriaAssuranceTimeoutCriteriaProp**](EnumresultCriteriaAssuranceTimeoutCriteriaProp.md) |  | [optional] 
**AssuranceTimeoutValue** | Pointer to **string** | The value to use for performing matching based on the assurance timeout. This will be ignored if the assurance-timeout-criteria is \&quot;any\&quot;. | [optional] 
**ResponseDelayedByAssurance** | Pointer to [**EnumresultCriteriaResponseDelayedByAssuranceProp**](EnumresultCriteriaResponseDelayedByAssuranceProp.md) |  | [optional] 
**AssuranceBehaviorAlteredByControl** | Pointer to [**EnumresultCriteriaAssuranceBehaviorAlteredByControlProp**](EnumresultCriteriaAssuranceBehaviorAlteredByControlProp.md) |  | [optional] 
**AssuranceSatisfied** | Pointer to [**EnumresultCriteriaAssuranceSatisfiedProp**](EnumresultCriteriaAssuranceSatisfiedProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewReplicationAssuranceResultCriteriaResponse

`func NewReplicationAssuranceResultCriteriaResponse(id string, schemas []EnumreplicationAssuranceResultCriteriaSchemaUrn, localAssuranceLevel []EnumresultCriteriaLocalAssuranceLevelProp, remoteAssuranceLevel []EnumresultCriteriaRemoteAssuranceLevelProp, ) *ReplicationAssuranceResultCriteriaResponse`

NewReplicationAssuranceResultCriteriaResponse instantiates a new ReplicationAssuranceResultCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationAssuranceResultCriteriaResponseWithDefaults

`func NewReplicationAssuranceResultCriteriaResponseWithDefaults() *ReplicationAssuranceResultCriteriaResponse`

NewReplicationAssuranceResultCriteriaResponseWithDefaults instantiates a new ReplicationAssuranceResultCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicationAssuranceResultCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicationAssuranceResultCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ReplicationAssuranceResultCriteriaResponse) GetSchemas() []EnumreplicationAssuranceResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetSchemasOk() (*[]EnumreplicationAssuranceResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationAssuranceResultCriteriaResponse) SetSchemas(v []EnumreplicationAssuranceResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLocalAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaResponse) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp`

GetLocalAssuranceLevel returns the LocalAssuranceLevel field if non-nil, zero value otherwise.

### GetLocalAssuranceLevelOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetLocalAssuranceLevelOk() (*[]EnumresultCriteriaLocalAssuranceLevelProp, bool)`

GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaResponse) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp)`

SetLocalAssuranceLevel sets LocalAssuranceLevel field to given value.


### GetRemoteAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaResponse) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp`

GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field if non-nil, zero value otherwise.

### GetRemoteAssuranceLevelOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetRemoteAssuranceLevelOk() (*[]EnumresultCriteriaRemoteAssuranceLevelProp, bool)`

GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaResponse) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp)`

SetRemoteAssuranceLevel sets RemoteAssuranceLevel field to given value.


### GetAssuranceTimeoutCriteria

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp`

GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field if non-nil, zero value otherwise.

### GetAssuranceTimeoutCriteriaOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool)`

GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutCriteria

`func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp)`

SetAssuranceTimeoutCriteria sets AssuranceTimeoutCriteria field to given value.

### HasAssuranceTimeoutCriteria

`func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceTimeoutCriteria() bool`

HasAssuranceTimeoutCriteria returns a boolean if a field has been set.

### GetAssuranceTimeoutValue

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutValue() string`

GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field if non-nil, zero value otherwise.

### GetAssuranceTimeoutValueOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceTimeoutValueOk() (*string, bool)`

GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutValue

`func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceTimeoutValue(v string)`

SetAssuranceTimeoutValue sets AssuranceTimeoutValue field to given value.

### HasAssuranceTimeoutValue

`func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceTimeoutValue() bool`

HasAssuranceTimeoutValue returns a boolean if a field has been set.

### GetResponseDelayedByAssurance

`func (o *ReplicationAssuranceResultCriteriaResponse) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp`

GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field if non-nil, zero value otherwise.

### GetResponseDelayedByAssuranceOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool)`

GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDelayedByAssurance

`func (o *ReplicationAssuranceResultCriteriaResponse) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp)`

SetResponseDelayedByAssurance sets ResponseDelayedByAssurance field to given value.

### HasResponseDelayedByAssurance

`func (o *ReplicationAssuranceResultCriteriaResponse) HasResponseDelayedByAssurance() bool`

HasResponseDelayedByAssurance returns a boolean if a field has been set.

### GetAssuranceBehaviorAlteredByControl

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp`

GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field if non-nil, zero value otherwise.

### GetAssuranceBehaviorAlteredByControlOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool)`

GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceBehaviorAlteredByControl

`func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp)`

SetAssuranceBehaviorAlteredByControl sets AssuranceBehaviorAlteredByControl field to given value.

### HasAssuranceBehaviorAlteredByControl

`func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceBehaviorAlteredByControl() bool`

HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.

### GetAssuranceSatisfied

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp`

GetAssuranceSatisfied returns the AssuranceSatisfied field if non-nil, zero value otherwise.

### GetAssuranceSatisfiedOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool)`

GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceSatisfied

`func (o *ReplicationAssuranceResultCriteriaResponse) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp)`

SetAssuranceSatisfied sets AssuranceSatisfied field to given value.

### HasAssuranceSatisfied

`func (o *ReplicationAssuranceResultCriteriaResponse) HasAssuranceSatisfied() bool`

HasAssuranceSatisfied returns a boolean if a field has been set.

### GetDescription

`func (o *ReplicationAssuranceResultCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReplicationAssuranceResultCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReplicationAssuranceResultCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *ReplicationAssuranceResultCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReplicationAssuranceResultCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReplicationAssuranceResultCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ReplicationAssuranceResultCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


