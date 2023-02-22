# AddJvmDefaultTrustManagerProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Trust Manager Provider | 
**Schemas** | [**[]EnumjvmDefaultTrustManagerProviderSchemaUrn**](EnumjvmDefaultTrustManagerProviderSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicate whether the Trust Manager Provider is enabled for use. | 

## Methods

### NewAddJvmDefaultTrustManagerProviderRequest

`func NewAddJvmDefaultTrustManagerProviderRequest(providerName string, schemas []EnumjvmDefaultTrustManagerProviderSchemaUrn, enabled bool, ) *AddJvmDefaultTrustManagerProviderRequest`

NewAddJvmDefaultTrustManagerProviderRequest instantiates a new AddJvmDefaultTrustManagerProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJvmDefaultTrustManagerProviderRequestWithDefaults

`func NewAddJvmDefaultTrustManagerProviderRequestWithDefaults() *AddJvmDefaultTrustManagerProviderRequest`

NewAddJvmDefaultTrustManagerProviderRequestWithDefaults instantiates a new AddJvmDefaultTrustManagerProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddJvmDefaultTrustManagerProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddJvmDefaultTrustManagerProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddJvmDefaultTrustManagerProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddJvmDefaultTrustManagerProviderRequest) GetSchemas() []EnumjvmDefaultTrustManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddJvmDefaultTrustManagerProviderRequest) GetSchemasOk() (*[]EnumjvmDefaultTrustManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddJvmDefaultTrustManagerProviderRequest) SetSchemas(v []EnumjvmDefaultTrustManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *AddJvmDefaultTrustManagerProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddJvmDefaultTrustManagerProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddJvmDefaultTrustManagerProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


