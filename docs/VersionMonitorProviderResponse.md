# VersionMonitorProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumversionMonitorProviderSchemaUrn**](EnumversionMonitorProviderSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 

## Methods

### NewVersionMonitorProviderResponse

`func NewVersionMonitorProviderResponse(schemas []EnumversionMonitorProviderSchemaUrn, enabled bool, ) *VersionMonitorProviderResponse`

NewVersionMonitorProviderResponse instantiates a new VersionMonitorProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionMonitorProviderResponseWithDefaults

`func NewVersionMonitorProviderResponseWithDefaults() *VersionMonitorProviderResponse`

NewVersionMonitorProviderResponseWithDefaults instantiates a new VersionMonitorProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *VersionMonitorProviderResponse) GetSchemas() []EnumversionMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VersionMonitorProviderResponse) GetSchemasOk() (*[]EnumversionMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VersionMonitorProviderResponse) SetSchemas(v []EnumversionMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *VersionMonitorProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionMonitorProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionMonitorProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionMonitorProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *VersionMonitorProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VersionMonitorProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VersionMonitorProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


