# AddThirdPartyTrustManagerProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Trust Manager Provider | 
**Schemas** | [**[]EnumthirdPartyTrustManagerProviderSchemaUrn**](EnumthirdPartyTrustManagerProviderSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Trust Manager Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Trust Manager Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Enabled** | **bool** | Indicate whether the Trust Manager Provider is enabled for use. | 
**IncludeJVMDefaultIssuers** | Pointer to **bool** | Indicates whether certificates issued by an authority included in the JVM&#39;s set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider. | [optional] 

## Methods

### NewAddThirdPartyTrustManagerProviderRequest

`func NewAddThirdPartyTrustManagerProviderRequest(providerName string, schemas []EnumthirdPartyTrustManagerProviderSchemaUrn, extensionClass string, enabled bool, ) *AddThirdPartyTrustManagerProviderRequest`

NewAddThirdPartyTrustManagerProviderRequest instantiates a new AddThirdPartyTrustManagerProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyTrustManagerProviderRequestWithDefaults

`func NewAddThirdPartyTrustManagerProviderRequestWithDefaults() *AddThirdPartyTrustManagerProviderRequest`

NewAddThirdPartyTrustManagerProviderRequestWithDefaults instantiates a new AddThirdPartyTrustManagerProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddThirdPartyTrustManagerProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddThirdPartyTrustManagerProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddThirdPartyTrustManagerProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddThirdPartyTrustManagerProviderRequest) GetSchemas() []EnumthirdPartyTrustManagerProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyTrustManagerProviderRequest) GetSchemasOk() (*[]EnumthirdPartyTrustManagerProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyTrustManagerProviderRequest) SetSchemas(v []EnumthirdPartyTrustManagerProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyTrustManagerProviderRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyTrustManagerProviderRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyTrustManagerProviderRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyTrustManagerProviderRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyTrustManagerProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyTrustManagerProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyTrustManagerProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIncludeJVMDefaultIssuers

`func (o *AddThirdPartyTrustManagerProviderRequest) GetIncludeJVMDefaultIssuers() bool`

GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field if non-nil, zero value otherwise.

### GetIncludeJVMDefaultIssuersOk

`func (o *AddThirdPartyTrustManagerProviderRequest) GetIncludeJVMDefaultIssuersOk() (*bool, bool)`

GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeJVMDefaultIssuers

`func (o *AddThirdPartyTrustManagerProviderRequest) SetIncludeJVMDefaultIssuers(v bool)`

SetIncludeJVMDefaultIssuers sets IncludeJVMDefaultIssuers field to given value.

### HasIncludeJVMDefaultIssuers

`func (o *AddThirdPartyTrustManagerProviderRequest) HasIncludeJVMDefaultIssuers() bool`

HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


