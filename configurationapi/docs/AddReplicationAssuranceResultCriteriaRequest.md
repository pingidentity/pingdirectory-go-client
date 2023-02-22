# AddReplicationAssuranceResultCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Result Criteria | 
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

### NewAddReplicationAssuranceResultCriteriaRequest

`func NewAddReplicationAssuranceResultCriteriaRequest(criteriaName string, schemas []EnumreplicationAssuranceResultCriteriaSchemaUrn, localAssuranceLevel []EnumresultCriteriaLocalAssuranceLevelProp, remoteAssuranceLevel []EnumresultCriteriaRemoteAssuranceLevelProp, ) *AddReplicationAssuranceResultCriteriaRequest`

NewAddReplicationAssuranceResultCriteriaRequest instantiates a new AddReplicationAssuranceResultCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddReplicationAssuranceResultCriteriaRequestWithDefaults

`func NewAddReplicationAssuranceResultCriteriaRequestWithDefaults() *AddReplicationAssuranceResultCriteriaRequest`

NewAddReplicationAssuranceResultCriteriaRequestWithDefaults instantiates a new AddReplicationAssuranceResultCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetSchemas() []EnumreplicationAssuranceResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetSchemasOk() (*[]EnumreplicationAssuranceResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetSchemas(v []EnumreplicationAssuranceResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLocalAssuranceLevel

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetLocalAssuranceLevel() []EnumresultCriteriaLocalAssuranceLevelProp`

GetLocalAssuranceLevel returns the LocalAssuranceLevel field if non-nil, zero value otherwise.

### GetLocalAssuranceLevelOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetLocalAssuranceLevelOk() (*[]EnumresultCriteriaLocalAssuranceLevelProp, bool)`

GetLocalAssuranceLevelOk returns a tuple with the LocalAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAssuranceLevel

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetLocalAssuranceLevel(v []EnumresultCriteriaLocalAssuranceLevelProp)`

SetLocalAssuranceLevel sets LocalAssuranceLevel field to given value.


### GetRemoteAssuranceLevel

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetRemoteAssuranceLevel() []EnumresultCriteriaRemoteAssuranceLevelProp`

GetRemoteAssuranceLevel returns the RemoteAssuranceLevel field if non-nil, zero value otherwise.

### GetRemoteAssuranceLevelOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetRemoteAssuranceLevelOk() (*[]EnumresultCriteriaRemoteAssuranceLevelProp, bool)`

GetRemoteAssuranceLevelOk returns a tuple with the RemoteAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssuranceLevel

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetRemoteAssuranceLevel(v []EnumresultCriteriaRemoteAssuranceLevelProp)`

SetRemoteAssuranceLevel sets RemoteAssuranceLevel field to given value.


### GetAssuranceTimeoutCriteria

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutCriteria() EnumresultCriteriaAssuranceTimeoutCriteriaProp`

GetAssuranceTimeoutCriteria returns the AssuranceTimeoutCriteria field if non-nil, zero value otherwise.

### GetAssuranceTimeoutCriteriaOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutCriteriaOk() (*EnumresultCriteriaAssuranceTimeoutCriteriaProp, bool)`

GetAssuranceTimeoutCriteriaOk returns a tuple with the AssuranceTimeoutCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutCriteria

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceTimeoutCriteria(v EnumresultCriteriaAssuranceTimeoutCriteriaProp)`

SetAssuranceTimeoutCriteria sets AssuranceTimeoutCriteria field to given value.

### HasAssuranceTimeoutCriteria

`func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceTimeoutCriteria() bool`

HasAssuranceTimeoutCriteria returns a boolean if a field has been set.

### GetAssuranceTimeoutValue

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutValue() string`

GetAssuranceTimeoutValue returns the AssuranceTimeoutValue field if non-nil, zero value otherwise.

### GetAssuranceTimeoutValueOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceTimeoutValueOk() (*string, bool)`

GetAssuranceTimeoutValueOk returns a tuple with the AssuranceTimeoutValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceTimeoutValue

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceTimeoutValue(v string)`

SetAssuranceTimeoutValue sets AssuranceTimeoutValue field to given value.

### HasAssuranceTimeoutValue

`func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceTimeoutValue() bool`

HasAssuranceTimeoutValue returns a boolean if a field has been set.

### GetResponseDelayedByAssurance

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetResponseDelayedByAssurance() EnumresultCriteriaResponseDelayedByAssuranceProp`

GetResponseDelayedByAssurance returns the ResponseDelayedByAssurance field if non-nil, zero value otherwise.

### GetResponseDelayedByAssuranceOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetResponseDelayedByAssuranceOk() (*EnumresultCriteriaResponseDelayedByAssuranceProp, bool)`

GetResponseDelayedByAssuranceOk returns a tuple with the ResponseDelayedByAssurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDelayedByAssurance

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetResponseDelayedByAssurance(v EnumresultCriteriaResponseDelayedByAssuranceProp)`

SetResponseDelayedByAssurance sets ResponseDelayedByAssurance field to given value.

### HasResponseDelayedByAssurance

`func (o *AddReplicationAssuranceResultCriteriaRequest) HasResponseDelayedByAssurance() bool`

HasResponseDelayedByAssurance returns a boolean if a field has been set.

### GetAssuranceBehaviorAlteredByControl

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceBehaviorAlteredByControl() EnumresultCriteriaAssuranceBehaviorAlteredByControlProp`

GetAssuranceBehaviorAlteredByControl returns the AssuranceBehaviorAlteredByControl field if non-nil, zero value otherwise.

### GetAssuranceBehaviorAlteredByControlOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceBehaviorAlteredByControlOk() (*EnumresultCriteriaAssuranceBehaviorAlteredByControlProp, bool)`

GetAssuranceBehaviorAlteredByControlOk returns a tuple with the AssuranceBehaviorAlteredByControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceBehaviorAlteredByControl

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceBehaviorAlteredByControl(v EnumresultCriteriaAssuranceBehaviorAlteredByControlProp)`

SetAssuranceBehaviorAlteredByControl sets AssuranceBehaviorAlteredByControl field to given value.

### HasAssuranceBehaviorAlteredByControl

`func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceBehaviorAlteredByControl() bool`

HasAssuranceBehaviorAlteredByControl returns a boolean if a field has been set.

### GetAssuranceSatisfied

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceSatisfied() EnumresultCriteriaAssuranceSatisfiedProp`

GetAssuranceSatisfied returns the AssuranceSatisfied field if non-nil, zero value otherwise.

### GetAssuranceSatisfiedOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetAssuranceSatisfiedOk() (*EnumresultCriteriaAssuranceSatisfiedProp, bool)`

GetAssuranceSatisfiedOk returns a tuple with the AssuranceSatisfied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceSatisfied

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetAssuranceSatisfied(v EnumresultCriteriaAssuranceSatisfiedProp)`

SetAssuranceSatisfied sets AssuranceSatisfied field to given value.

### HasAssuranceSatisfied

`func (o *AddReplicationAssuranceResultCriteriaRequest) HasAssuranceSatisfied() bool`

HasAssuranceSatisfied returns a boolean if a field has been set.

### GetDescription

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddReplicationAssuranceResultCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddReplicationAssuranceResultCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddReplicationAssuranceResultCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


