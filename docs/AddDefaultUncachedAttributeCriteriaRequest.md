# AddDefaultUncachedAttributeCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Uncached Attribute Criteria | 
**Schemas** | [**[]EnumdefaultUncachedAttributeCriteriaSchemaUrn**](EnumdefaultUncachedAttributeCriteriaSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Uncached Attribute Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Attribute Criteria is enabled for use in the server. | 

## Methods

### NewAddDefaultUncachedAttributeCriteriaRequest

`func NewAddDefaultUncachedAttributeCriteriaRequest(criteriaName string, schemas []EnumdefaultUncachedAttributeCriteriaSchemaUrn, enabled bool, ) *AddDefaultUncachedAttributeCriteriaRequest`

NewAddDefaultUncachedAttributeCriteriaRequest instantiates a new AddDefaultUncachedAttributeCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDefaultUncachedAttributeCriteriaRequestWithDefaults

`func NewAddDefaultUncachedAttributeCriteriaRequestWithDefaults() *AddDefaultUncachedAttributeCriteriaRequest`

NewAddDefaultUncachedAttributeCriteriaRequestWithDefaults instantiates a new AddDefaultUncachedAttributeCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddDefaultUncachedAttributeCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetSchemas() []EnumdefaultUncachedAttributeCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetSchemasOk() (*[]EnumdefaultUncachedAttributeCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDefaultUncachedAttributeCriteriaRequest) SetSchemas(v []EnumdefaultUncachedAttributeCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDefaultUncachedAttributeCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDefaultUncachedAttributeCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDefaultUncachedAttributeCriteriaRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDefaultUncachedAttributeCriteriaRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


