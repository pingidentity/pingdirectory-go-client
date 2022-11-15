# JvmDefaultTrustManagerProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Trust Manager Provider | 
**Schemas** | [**[]EnumjvmDefaultTrustManagerProviderSchemaUrn**](EnumjvmDefaultTrustManagerProviderSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicate whether the Trust Manager Provider is enabled for use. | 
**IncludeJVMDefaultIssuers** | Pointer to **bool** | Indicates whether certificates issued by an authority included in the JVM&#39;s set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider. | [optional] 

## Methods

### NewJvmDefaultTrustManagerProviderResponse

`func NewJvmDefaultTrustManagerProviderResponse(id string, schemas []EnumjvmDefaultTrustManagerProviderSchemaUrn, enabled bool, ) *JvmDefaultTrustManagerProviderResponse`

NewJvmDefaultTrustManagerProviderResponse instantiates a new JvmDefaultTrustManagerProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJvmDefaultTrustManagerProviderResponseWithDefaults

`func NewJvmDefaultTrustManagerProviderResponseWithDefaults() *JvmDefaultTrustManagerProviderResponse`

NewJvmDefaultTrustManagerProviderResponseWithDefaults instantiates a new JvmDefaultTrustManagerProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JvmDefaultTrustManagerProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JvmDefaultTrustManagerProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JvmDefaultTrustManagerProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *JvmDefaultTrustManagerProviderResponse) GetSchemas() []EnumjvmDefaultTrustManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JvmDefaultTrustManagerProviderResponse) GetSchemasOk() (*[]EnumjvmDefaultTrustManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JvmDefaultTrustManagerProviderResponse) SetSchemas(v []EnumjvmDefaultTrustManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *JvmDefaultTrustManagerProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JvmDefaultTrustManagerProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JvmDefaultTrustManagerProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeJVMDefaultIssuers

`func (o *JvmDefaultTrustManagerProviderResponse) GetIncludeJVMDefaultIssuers() bool`

GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field if non-nil, zero value otherwise.

### GetIncludeJVMDefaultIssuersOk

`func (o *JvmDefaultTrustManagerProviderResponse) GetIncludeJVMDefaultIssuersOk() (*bool, bool)`

GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJVMDefaultIssuers

`func (o *JvmDefaultTrustManagerProviderResponse) SetIncludeJVMDefaultIssuers(v bool)`

SetIncludeJVMDefaultIssuers sets IncludeJVMDefaultIssuers field to given value.

### HasIncludeJVMDefaultIssuers

`func (o *JvmDefaultTrustManagerProviderResponse) HasIncludeJVMDefaultIssuers() bool`

HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


