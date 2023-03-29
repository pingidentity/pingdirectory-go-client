# CustomVelocityContextProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcustomVelocityContextProviderSchemaUrn**](EnumcustomVelocityContextProviderSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**Enabled** | Pointer to **bool** | Indicates whether this Velocity Context Provider is enabled. If set to &#39;false&#39; this Velocity Context Provider will not contribute context content for any requests. | [optional] 
**ObjectScope** | Pointer to [**EnumvelocityContextProviderObjectScopeProp**](EnumvelocityContextProviderObjectScopeProp.md) |  | [optional] 
**IncludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will contribute content. | [optional] 
**ExcludedView** | Pointer to **[]string** | The name of a view for which this Velocity Context Provider will not contribute content. | [optional] 
**HttpMethod** | Pointer to **[]string** | Specifies the set of HTTP methods handled by this Velocity Context Provider, which will perform actions necessary to fulfill the request before updating the context for the response. The values of this property are not case-sensitive. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for template page requests to which this Velocity Context Provider contributes content. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewCustomVelocityContextProviderResponse

`func NewCustomVelocityContextProviderResponse(schemas []EnumcustomVelocityContextProviderSchemaUrn, id string, ) *CustomVelocityContextProviderResponse`

NewCustomVelocityContextProviderResponse instantiates a new CustomVelocityContextProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomVelocityContextProviderResponseWithDefaults

`func NewCustomVelocityContextProviderResponseWithDefaults() *CustomVelocityContextProviderResponse`

NewCustomVelocityContextProviderResponseWithDefaults instantiates a new CustomVelocityContextProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *CustomVelocityContextProviderResponse) GetSchemas() []EnumcustomVelocityContextProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CustomVelocityContextProviderResponse) GetSchemasOk() (*[]EnumcustomVelocityContextProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CustomVelocityContextProviderResponse) SetSchemas(v []EnumcustomVelocityContextProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *CustomVelocityContextProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomVelocityContextProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomVelocityContextProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *CustomVelocityContextProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomVelocityContextProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomVelocityContextProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomVelocityContextProviderResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectScope

`func (o *CustomVelocityContextProviderResponse) GetObjectScope() EnumvelocityContextProviderObjectScopeProp`

GetObjectScope returns the ObjectScope field if non-nil, zero value otherwise.

### GetObjectScopeOk

`func (o *CustomVelocityContextProviderResponse) GetObjectScopeOk() (*EnumvelocityContextProviderObjectScopeProp, bool)`

GetObjectScopeOk returns a tuple with the ObjectScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScope

`func (o *CustomVelocityContextProviderResponse) SetObjectScope(v EnumvelocityContextProviderObjectScopeProp)`

SetObjectScope sets ObjectScope field to given value.

### HasObjectScope

`func (o *CustomVelocityContextProviderResponse) HasObjectScope() bool`

HasObjectScope returns a boolean if a field has been set.

### GetIncludedView

`func (o *CustomVelocityContextProviderResponse) GetIncludedView() []string`

GetIncludedView returns the IncludedView field if non-nil, zero value otherwise.

### GetIncludedViewOk

`func (o *CustomVelocityContextProviderResponse) GetIncludedViewOk() (*[]string, bool)`

GetIncludedViewOk returns a tuple with the IncludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedView

`func (o *CustomVelocityContextProviderResponse) SetIncludedView(v []string)`

SetIncludedView sets IncludedView field to given value.

### HasIncludedView

`func (o *CustomVelocityContextProviderResponse) HasIncludedView() bool`

HasIncludedView returns a boolean if a field has been set.

### GetExcludedView

`func (o *CustomVelocityContextProviderResponse) GetExcludedView() []string`

GetExcludedView returns the ExcludedView field if non-nil, zero value otherwise.

### GetExcludedViewOk

`func (o *CustomVelocityContextProviderResponse) GetExcludedViewOk() (*[]string, bool)`

GetExcludedViewOk returns a tuple with the ExcludedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedView

`func (o *CustomVelocityContextProviderResponse) SetExcludedView(v []string)`

SetExcludedView sets ExcludedView field to given value.

### HasExcludedView

`func (o *CustomVelocityContextProviderResponse) HasExcludedView() bool`

HasExcludedView returns a boolean if a field has been set.

### GetHttpMethod

`func (o *CustomVelocityContextProviderResponse) GetHttpMethod() []string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *CustomVelocityContextProviderResponse) GetHttpMethodOk() (*[]string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *CustomVelocityContextProviderResponse) SetHttpMethod(v []string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *CustomVelocityContextProviderResponse) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetResponseHeader

`func (o *CustomVelocityContextProviderResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *CustomVelocityContextProviderResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *CustomVelocityContextProviderResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *CustomVelocityContextProviderResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetMeta

`func (o *CustomVelocityContextProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CustomVelocityContextProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CustomVelocityContextProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CustomVelocityContextProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CustomVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CustomVelocityContextProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CustomVelocityContextProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CustomVelocityContextProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


