# HttpConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumhttpConfigurationSchemaUrn**](EnumhttpConfigurationSchemaUrn.md) |  | [optional] 
**IncludeStackTracesInErrorPages** | Pointer to **bool** | Indicates whether exceptions thrown by servlet or web application extensions will be included in the resulting error page response. Stack traces can be helpful in diagnosing application errors, but in production they may reveal information that might be useful to a malicious attacker. | [optional] 
**IncludeServletInformationInErrorPages** | Pointer to **bool** | Indicates whether to expose servlet information in the error page response. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewHttpConfigurationResponse

`func NewHttpConfigurationResponse() *HttpConfigurationResponse`

NewHttpConfigurationResponse instantiates a new HttpConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpConfigurationResponseWithDefaults

`func NewHttpConfigurationResponseWithDefaults() *HttpConfigurationResponse`

NewHttpConfigurationResponseWithDefaults instantiates a new HttpConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *HttpConfigurationResponse) GetSchemas() []EnumhttpConfigurationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpConfigurationResponse) GetSchemasOk() (*[]EnumhttpConfigurationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpConfigurationResponse) SetSchemas(v []EnumhttpConfigurationSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *HttpConfigurationResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetIncludeStackTracesInErrorPages

`func (o *HttpConfigurationResponse) GetIncludeStackTracesInErrorPages() bool`

GetIncludeStackTracesInErrorPages returns the IncludeStackTracesInErrorPages field if non-nil, zero value otherwise.

### GetIncludeStackTracesInErrorPagesOk

`func (o *HttpConfigurationResponse) GetIncludeStackTracesInErrorPagesOk() (*bool, bool)`

GetIncludeStackTracesInErrorPagesOk returns a tuple with the IncludeStackTracesInErrorPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStackTracesInErrorPages

`func (o *HttpConfigurationResponse) SetIncludeStackTracesInErrorPages(v bool)`

SetIncludeStackTracesInErrorPages sets IncludeStackTracesInErrorPages field to given value.

### HasIncludeStackTracesInErrorPages

`func (o *HttpConfigurationResponse) HasIncludeStackTracesInErrorPages() bool`

HasIncludeStackTracesInErrorPages returns a boolean if a field has been set.

### GetIncludeServletInformationInErrorPages

`func (o *HttpConfigurationResponse) GetIncludeServletInformationInErrorPages() bool`

GetIncludeServletInformationInErrorPages returns the IncludeServletInformationInErrorPages field if non-nil, zero value otherwise.

### GetIncludeServletInformationInErrorPagesOk

`func (o *HttpConfigurationResponse) GetIncludeServletInformationInErrorPagesOk() (*bool, bool)`

GetIncludeServletInformationInErrorPagesOk returns a tuple with the IncludeServletInformationInErrorPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeServletInformationInErrorPages

`func (o *HttpConfigurationResponse) SetIncludeServletInformationInErrorPages(v bool)`

SetIncludeServletInformationInErrorPages sets IncludeServletInformationInErrorPages field to given value.

### HasIncludeServletInformationInErrorPages

`func (o *HttpConfigurationResponse) HasIncludeServletInformationInErrorPages() bool`

HasIncludeServletInformationInErrorPages returns a boolean if a field has been set.

### GetMeta

`func (o *HttpConfigurationResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HttpConfigurationResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HttpConfigurationResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HttpConfigurationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpConfigurationResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HttpConfigurationResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpConfigurationResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HttpConfigurationResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


