# ReplicationAssuranceResultCriteriaShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumreplicationAssuranceResultCriteriaSchemaUrn**](EnumreplicationAssuranceResultCriteriaSchemaUrn.md) |  | 
**LocalAssuranceLevel** | [**[]EnumresultCriteriaLocalAssuranceLevelProp**](EnumresultCriteriaLocalAssuranceLevelProp.md) |  | 
**RemoteAssuranceLevel** | [**[]EnumresultCriteriaRemoteAssuranceLevelProp**](EnumresultCriteriaRemoteAssuranceLevelProp.md) |  | 
**AssuranceTimeoutCriteria** | Pointer to [**EnumresultCriteriaAssuranceTimeoutCriteriaProp**](EnumresultCriteriaAssuranceTimeoutCriteriaProp.md) |  | [optional] 
**AssuranceTimeoutValue** | Pointer to **string** | The value to use for performing matching based on the assurance timeout. This will be ignored if the assurance-timeout-criteria is \&quot;any\&quot;. | [optional] 
**ResponseDelayedByAssurance** | Pointer to [**EnumresultCriteriaResponseDelayedByAssuranceProp**](EnumresultCriteriaResponseDelayedByAssuranceProp.md) |  | [optional] 
**AssuranceBehaviorAlteredByControl** | Pointer to [**EnumresultCriteriaAssuranceBehaviorAlteredByControlProp**](EnumresultCriteriaAssuranceBehaviorAlteredByControlProp.md) |  | [optional] 
**AssuranceSatisfied** | Pointer to [**EnumresultCriteriaAssuranceSatisfiedProp**](EnumresultCriteriaAssuranceSatisfiedProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 

## Methods

### NewReplicationAssuranceResultCriteriaShared

`func NewReplicationAssuranceResultCriteriaShared(schemas []EnumreplicationAssuranceResultCriteriaSchemaUrn, localAssuranceLevel []EnumresultCriteriaLocalAssuranceLevelProp, remoteAssuranceLevel []EnumresultCriteriaRemoteAssuranceLevelProp, ) *ReplicationAssuranceResultCriteriaShared`

NewReplicationAssuranceResultCriteriaShared instantiates a new ReplicationAssuranceResultCriteriaShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationAssuranceResultCriteriaSharedWithDefaults

`func NewReplicationAssuranceResultCriteriaSharedWithDefaults() *ReplicationAssuranceResultCriteriaShared`

NewReplicationAssuranceResultCriteriaSharedWithDefaults instantiates a new ReplicationAssuranceResultCriteriaShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ReplicationAssuranceResultCriteriaShared) GetSchemas() []EnumreplicationAssuranceResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetSchemasOk() (*[]EnumreplicationAssuranceResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReplicationAssuranceResultCriteriaShared) SetSchemas(v []EnumreplicationAssuranceResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLocalAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaShared) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp`

GetLocalAssuranceLevel returns the LocalAssuranceLevel field if non-nil, zero value otherwise.

### GetLocalAssuranceLevelOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetLocalAssuranceLevelOk() (*[]EnumresultCriteriaLocalAssuranceLevelProp, bool)`

GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaShared) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp)`

SetLocalAssuranceLevel sets LocalAssuranceLevel field to given value.


### GetRemoteAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaShared) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp`

GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field if non-nil, zero value otherwise.

### GetRemoteAssuranceLevelOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetRemoteAssuranceLevelOk() (*[]EnumresultCriteriaRemoteAssuranceLevelProp, bool)`

GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssuranceLevel

`func (o *ReplicationAssuranceResultCriteriaShared) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp)`

SetRemoteAssuranceLevel sets RemoteAssuranceLevel field to given value.


### GetAssuranceTimeoutCriteria

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp`

GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field if non-nil, zero value otherwise.

### GetAssuranceTimeoutCriteriaOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool)`

GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutCriteria

`func (o *ReplicationAssuranceResultCriteriaShared) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp)`

SetAssuranceTimeoutCriteria sets AssuranceTimeoutCriteria field to given value.

### HasAssuranceTimeoutCriteria

`func (o *ReplicationAssuranceResultCriteriaShared) HasAssuranceTimeoutCriteria() bool`

HasAssuranceTimeoutCriteria returns a boolean if a field has been set.

### GetAssuranceTimeoutValue

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceTimeoutValue() string`

GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field if non-nil, zero value otherwise.

### GetAssuranceTimeoutValueOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceTimeoutValueOk() (*string, bool)`

GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutValue

`func (o *ReplicationAssuranceResultCriteriaShared) SetAssuranceTimeoutValue(v string)`

SetAssuranceTimeoutValue sets AssuranceTimeoutValue field to given value.

### HasAssuranceTimeoutValue

`func (o *ReplicationAssuranceResultCriteriaShared) HasAssuranceTimeoutValue() bool`

HasAssuranceTimeoutValue returns a boolean if a field has been set.

### GetResponseDelayedByAssurance

`func (o *ReplicationAssuranceResultCriteriaShared) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp`

GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field if non-nil, zero value otherwise.

### GetResponseDelayedByAssuranceOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool)`

GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDelayedByAssurance

`func (o *ReplicationAssuranceResultCriteriaShared) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp)`

SetResponseDelayedByAssurance sets ResponseDelayedByAssurance field to given value.

### HasResponseDelayedByAssurance

`func (o *ReplicationAssuranceResultCriteriaShared) HasResponseDelayedByAssurance() bool`

HasResponseDelayedByAssurance returns a boolean if a field has been set.

### GetAssuranceBehaviorAlteredByControl

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp`

GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field if non-nil, zero value otherwise.

### GetAssuranceBehaviorAlteredByControlOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool)`

GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceBehaviorAlteredByControl

`func (o *ReplicationAssuranceResultCriteriaShared) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp)`

SetAssuranceBehaviorAlteredByControl sets AssuranceBehaviorAlteredByControl field to given value.

### HasAssuranceBehaviorAlteredByControl

`func (o *ReplicationAssuranceResultCriteriaShared) HasAssuranceBehaviorAlteredByControl() bool`

HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.

### GetAssuranceSatisfied

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp`

GetAssuranceSatisfied returns the AssuranceSatisfied field if non-nil, zero value otherwise.

### GetAssuranceSatisfiedOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool)`

GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceSatisfied

`func (o *ReplicationAssuranceResultCriteriaShared) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp)`

SetAssuranceSatisfied sets AssuranceSatisfied field to given value.

### HasAssuranceSatisfied

`func (o *ReplicationAssuranceResultCriteriaShared) HasAssuranceSatisfied() bool`

HasAssuranceSatisfied returns a boolean if a field has been set.

### GetDescription

`func (o *ReplicationAssuranceResultCriteriaShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReplicationAssuranceResultCriteriaShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReplicationAssuranceResultCriteriaShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReplicationAssuranceResultCriteriaShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


