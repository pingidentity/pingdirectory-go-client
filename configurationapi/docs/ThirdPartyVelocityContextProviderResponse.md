# ThirdPartyVelocityContextProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Velocity Context Provider | 
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

### NewThirdPartyVelocityContextProviderResponse

`func NewThirdPartyVelocityContextProviderResponse(id string, schemas []EnumthirdPartyVelocityContextProviderSchemaUrn, extensionClass string, ) *ThirdPartyVelocityContextProviderResponse`

NewThirdPartyVelocityContextProviderResponse instantiates a new ThirdPartyVelocityContextProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyVelocityContextProviderResponseWithDefaults

`func NewThirdPartyVelocityContextProviderResponseWithDefaults() *ThirdPartyVelocityContextProviderResponse`

NewThirdPartyVelocityContextProviderResponseWithDefaults instantiates a new ThirdPartyVelocityContextProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ThirdPartyVelocityContextProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ThirdPartyVelocityContextProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ThirdPartyVelocityContextProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ThirdPartyVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyVelocityContextProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ThirdPartyVelocityContextProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ThirdPartyVelocityContextProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyVelocityContextProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ThirdPartyVelocityContextProviderResponse) GetSchemas() []EnumthirdPartyVelocityContextProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetSchemasOk() (*[]EnumthirdPartyVelocityContextProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ThirdPartyVelocityContextProviderResponse) SetSchemas(v []EnumthirdPartyVelocityContextProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *ThirdPartyVelocityContextProviderResponse) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ThirdPartyVelocityContextProviderResponse) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ThirdPartyVelocityContextProviderResponse) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ThirdPartyVelocityContextProviderResponse) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ThirdPartyVelocityContextProviderResponse) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetEnabled

`func (o *ThirdPartyVelocityContextProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ThirdPartyVelocityContextProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ThirdPartyVelocityContextProviderResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectScope

`func (o *ThirdPartyVelocityContextProviderResponse) GetObjectScope() EnumvelocityContextProviderObjectScopeProp`

GetObjectScope returns the ObjectScope field if non-nil, zero value otherwise.

### GetObjectScopeOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool)`

GetObjectScopeOk returns a tuple with the ObjectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScope

`func (o *ThirdPartyVelocityContextProviderResponse) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp)`

SetObjectScope sets ObjectScope field to given value.

### HasObjectScope

`func (o *ThirdPartyVelocityContextProviderResponse) HasObjectScope() bool`

HasObjectScope returns a boolean if a field has been set.

### GetIncludedView

`func (o *ThirdPartyVelocityContextProviderResponse) GetIncludedView() []string`

GetIncludedView returns the IncludedView field if non-nil, zero value otherwise.

### GetIncludedViewOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetIncludedViewOk() (*[]string, bool)`

GetIncludedViewOk returns a tuple with the IncludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedView

`func (o *ThirdPartyVelocityContextProviderResponse) SetIncludedView(v []string)`

SetIncludedView sets IncludedView field to given value.

### HasIncludedView

`func (o *ThirdPartyVelocityContextProviderResponse) HasIncludedView() bool`

HasIncludedView returns a boolean if a field has been set.

### GetExcludedView

`func (o *ThirdPartyVelocityContextProviderResponse) GetExcludedView() []string`

GetExcludedView returns the ExcludedView field if non-nil, zero value otherwise.

### GetExcludedViewOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetExcludedViewOk() (*[]string, bool)`

GetExcludedViewOk returns a tuple with the ExcludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedView

`func (o *ThirdPartyVelocityContextProviderResponse) SetExcludedView(v []string)`

SetExcludedView sets ExcludedView field to given value.

### HasExcludedView

`func (o *ThirdPartyVelocityContextProviderResponse) HasExcludedView() bool`

HasExcludedView returns a boolean if a field has been set.

### GetHttpMethod

`func (o *ThirdPartyVelocityContextProviderResponse) GetHttpMethod() []string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetHttpMethodOk() (*[]string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ThirdPartyVelocityContextProviderResponse) SetHttpMethod(v []string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *ThirdPartyVelocityContextProviderResponse) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetResponseHeader

`func (o *ThirdPartyVelocityContextProviderResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *ThirdPartyVelocityContextProviderResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *ThirdPartyVelocityContextProviderResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *ThirdPartyVelocityContextProviderResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


