# AddThirdPartyVelocityContextProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Velocity Context Provider | 
**Schemas** | [**[]EnumthirdPartyVelocityContextProviderSchemaUrn**](EnumthirdPartyVelocityContextProviderSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Velocity Context Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Velocity Context Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Velocity Context Provider is enabled. If set to &#39;false&#39; this Velocity Context Provider will not contribute context content for any requests. | [optional] 
**ObjectScope** | Pointer to [**EnumvelocityContextProviderObjectScopeProp**](EnumvelocityContextProviderObjectScopeProp.md) |  | [optional] 
**IncludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will contribute content. | [optional] 
**ExcludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will not contribute content. | [optional] 
**HttpMethod** | Pointer to **[]string** | Specifies the set of HTTP methods handled by this Velocity Context Provider, which will perform actions necessary to fulfill the request before updating the context for the response. The values of this property are not case-sensitive. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for template page requests to which this Velocity Context Provider contributes content. | [optional] 

## Methods

### NewAddThirdPartyVelocityContextProviderRequest

`func NewAddThirdPartyVelocityContextProviderRequest(providerName string, schemas []EnumthirdPartyVelocityContextProviderSchemaUrn, extensionClass string, ) *AddThirdPartyVelocityContextProviderRequest`

NewAddThirdPartyVelocityContextProviderRequest instantiates a new AddThirdPartyVelocityContextProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyVelocityContextProviderRequestWithDefaults

`func NewAddThirdPartyVelocityContextProviderRequestWithDefaults() *AddThirdPartyVelocityContextProviderRequest`

NewAddThirdPartyVelocityContextProviderRequestWithDefaults instantiates a new AddThirdPartyVelocityContextProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddThirdPartyVelocityContextProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddThirdPartyVelocityContextProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddThirdPartyVelocityContextProviderRequest) GetSchemas() []EnumthirdPartyVelocityContextProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetSchemasOk() (*[]EnumthirdPartyVelocityContextProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyVelocityContextProviderRequest) SetSchemas(v []EnumthirdPartyVelocityContextProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyVelocityContextProviderRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyVelocityContextProviderRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyVelocityContextProviderRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyVelocityContextProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyVelocityContextProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddThirdPartyVelocityContextProviderRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectScope

`func (o *AddThirdPartyVelocityContextProviderRequest) GetObjectScope() EnumvelocityContextProviderObjectScopeProp`

GetObjectScope returns the ObjectScope field if non-nil, zero value otherwise.

### GetObjectScopeOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool)`

GetObjectScopeOk returns a tuple with the ObjectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScope

`func (o *AddThirdPartyVelocityContextProviderRequest) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp)`

SetObjectScope sets ObjectScope field to given value.

### HasObjectScope

`func (o *AddThirdPartyVelocityContextProviderRequest) HasObjectScope() bool`

HasObjectScope returns a boolean if a field has been set.

### GetIncludedView

`func (o *AddThirdPartyVelocityContextProviderRequest) GetIncludedView() []string`

GetIncludedView returns the IncludedView field if non-nil, zero value otherwise.

### GetIncludedViewOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetIncludedViewOk() (*[]string, bool)`

GetIncludedViewOk returns a tuple with the IncludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedView

`func (o *AddThirdPartyVelocityContextProviderRequest) SetIncludedView(v []string)`

SetIncludedView sets IncludedView field to given value.

### HasIncludedView

`func (o *AddThirdPartyVelocityContextProviderRequest) HasIncludedView() bool`

HasIncludedView returns a boolean if a field has been set.

### GetExcludedView

`func (o *AddThirdPartyVelocityContextProviderRequest) GetExcludedView() []string`

GetExcludedView returns the ExcludedView field if non-nil, zero value otherwise.

### GetExcludedViewOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetExcludedViewOk() (*[]string, bool)`

GetExcludedViewOk returns a tuple with the ExcludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedView

`func (o *AddThirdPartyVelocityContextProviderRequest) SetExcludedView(v []string)`

SetExcludedView sets ExcludedView field to given value.

### HasExcludedView

`func (o *AddThirdPartyVelocityContextProviderRequest) HasExcludedView() bool`

HasExcludedView returns a boolean if a field has been set.

### GetHttpMethod

`func (o *AddThirdPartyVelocityContextProviderRequest) GetHttpMethod() []string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetHttpMethodOk() (*[]string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *AddThirdPartyVelocityContextProviderRequest) SetHttpMethod(v []string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *AddThirdPartyVelocityContextProviderRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetResponseHeader

`func (o *AddThirdPartyVelocityContextProviderRequest) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *AddThirdPartyVelocityContextProviderRequest) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *AddThirdPartyVelocityContextProviderRequest) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *AddThirdPartyVelocityContextProviderRequest) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


