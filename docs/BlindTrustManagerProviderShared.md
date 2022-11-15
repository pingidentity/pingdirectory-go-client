# BlindTrustManagerProviderShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumblindTrustManagerProviderSchemaUrn**](EnumblindTrustManagerProviderSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicate whether the Trust Manager Provider is enabled for use. | 
**IncludeJVMDefaultIssuers** | Pointer to **bool** | Indicates whether certificates issued by an authority included in the JVM&#39;s set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider. | [optional] 

## Methods

### NewBlindTrustManagerProviderShared

`func NewBlindTrustManagerProviderShared(schemas []EnumblindTrustManagerProviderSchemaUrn, enabled bool, ) *BlindTrustManagerProviderShared`

NewBlindTrustManagerProviderShared instantiates a new BlindTrustManagerProviderShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlindTrustManagerProviderSharedWithDefaults

`func NewBlindTrustManagerProviderSharedWithDefaults() *BlindTrustManagerProviderShared`

NewBlindTrustManagerProviderSharedWithDefaults instantiates a new BlindTrustManagerProviderShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BlindTrustManagerProviderShared) GetSchemas() []EnumblindTrustManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BlindTrustManagerProviderShared) GetSchemasOk() (*[]EnumblindTrustManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BlindTrustManagerProviderShared) SetSchemas(v []EnumblindTrustManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *BlindTrustManagerProviderShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlindTrustManagerProviderShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlindTrustManagerProviderShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeJVMDefaultIssuers

`func (o *BlindTrustManagerProviderShared) GetIncludeJVMDefaultIssuers() bool`

GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field if non-nil, zero value otherwise.

### GetIncludeJVMDefaultIssuersOk

`func (o *BlindTrustManagerProviderShared) GetIncludeJVMDefaultIssuersOk() (*bool, bool)`

GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJVMDefaultIssuers

`func (o *BlindTrustManagerProviderShared) SetIncludeJVMDefaultIssuers(v bool)`

SetIncludeJVMDefaultIssuers sets IncludeJVMDefaultIssuers field to given value.

### HasIncludeJVMDefaultIssuers

`func (o *BlindTrustManagerProviderShared) HasIncludeJVMDefaultIssuers() bool`

HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


