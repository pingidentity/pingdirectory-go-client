# DirectoryRestApiHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdirectoryRestApiHttpServletExtensionSchemaUrn**](EnumdirectoryRestApiHttpServletExtensionSchemaUrn.md) |  | 
**BasicAuthEnabled** | Pointer to **bool** | Enables HTTP Basic authentication, using a username and password. The Identity Mapper specified by the identity-mapper property will be used to map the username to a DN. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the Identity Mapper that is to be used for associating user entries with basic authentication usernames. | [optional] 
**AccessTokenValidator** | Pointer to **[]string** |  | [optional] 
**AccessTokenScope** | Pointer to **string** | The name of a scope that must be present in an access token accepted by the Directory REST API HTTP Servlet Extension. | [optional] 
**Audience** | Pointer to **string** | A string or URI that identifies the Directory REST API HTTP Servlet Extension in the context of OAuth2 authorization. | [optional] 
**MaxPageSize** | Pointer to **int32** | The maximum number of entries to be returned in one page of search results. | [optional] 
**SchemasEndpointObjectclass** | Pointer to **[]string** |  | [optional] 
**DefaultOperationalAttribute** | Pointer to **[]string** |  | [optional] 
**RejectExpansionAttribute** | Pointer to **[]string** |  | [optional] 
**AllowedControl** | Pointer to [**[]EnumhttpServletExtensionAllowedControlProp**](EnumhttpServletExtensionAllowedControlProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** |  | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewDirectoryRestApiHttpServletExtensionResponse

`func NewDirectoryRestApiHttpServletExtensionResponse(schemas []EnumdirectoryRestApiHttpServletExtensionSchemaUrn, ) *DirectoryRestApiHttpServletExtensionResponse`

NewDirectoryRestApiHttpServletExtensionResponse instantiates a new DirectoryRestApiHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryRestApiHttpServletExtensionResponseWithDefaults

`func NewDirectoryRestApiHttpServletExtensionResponseWithDefaults() *DirectoryRestApiHttpServletExtensionResponse`

NewDirectoryRestApiHttpServletExtensionResponseWithDefaults instantiates a new DirectoryRestApiHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetSchemas() []EnumdirectoryRestApiHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetSchemasOk() (*[]EnumdirectoryRestApiHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetSchemas(v []EnumdirectoryRestApiHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBasicAuthEnabled

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetAccessTokenScope

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAccessTokenScope() string`

GetAccessTokenScope returns the AccessTokenScope field if non-nil, zero value otherwise.

### GetAccessTokenScopeOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAccessTokenScopeOk() (*string, bool)`

GetAccessTokenScopeOk returns a tuple with the AccessTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenScope

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetAccessTokenScope(v string)`

SetAccessTokenScope sets AccessTokenScope field to given value.

### HasAccessTokenScope

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasAccessTokenScope() bool`

HasAccessTokenScope returns a boolean if a field has been set.

### GetAudience

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetMaxPageSize

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetMaxPageSize() int32`

GetMaxPageSize returns the MaxPageSize field if non-nil, zero value otherwise.

### GetMaxPageSizeOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetMaxPageSizeOk() (*int32, bool)`

GetMaxPageSizeOk returns a tuple with the MaxPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPageSize

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetMaxPageSize(v int32)`

SetMaxPageSize sets MaxPageSize field to given value.

### HasMaxPageSize

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasMaxPageSize() bool`

HasMaxPageSize returns a boolean if a field has been set.

### GetSchemasEndpointObjectclass

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetSchemasEndpointObjectclass() []string`

GetSchemasEndpointObjectclass returns the SchemasEndpointObjectclass field if non-nil, zero value otherwise.

### GetSchemasEndpointObjectclassOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetSchemasEndpointObjectclassOk() (*[]string, bool)`

GetSchemasEndpointObjectclassOk returns a tuple with the SchemasEndpointObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemasEndpointObjectclass

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetSchemasEndpointObjectclass(v []string)`

SetSchemasEndpointObjectclass sets SchemasEndpointObjectclass field to given value.

### HasSchemasEndpointObjectclass

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasSchemasEndpointObjectclass() bool`

HasSchemasEndpointObjectclass returns a boolean if a field has been set.

### GetDefaultOperationalAttribute

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetDefaultOperationalAttribute() []string`

GetDefaultOperationalAttribute returns the DefaultOperationalAttribute field if non-nil, zero value otherwise.

### GetDefaultOperationalAttributeOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetDefaultOperationalAttributeOk() (*[]string, bool)`

GetDefaultOperationalAttributeOk returns a tuple with the DefaultOperationalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOperationalAttribute

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetDefaultOperationalAttribute(v []string)`

SetDefaultOperationalAttribute sets DefaultOperationalAttribute field to given value.

### HasDefaultOperationalAttribute

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasDefaultOperationalAttribute() bool`

HasDefaultOperationalAttribute returns a boolean if a field has been set.

### GetRejectExpansionAttribute

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetRejectExpansionAttribute() []string`

GetRejectExpansionAttribute returns the RejectExpansionAttribute field if non-nil, zero value otherwise.

### GetRejectExpansionAttributeOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetRejectExpansionAttributeOk() (*[]string, bool)`

GetRejectExpansionAttributeOk returns a tuple with the RejectExpansionAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectExpansionAttribute

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetRejectExpansionAttribute(v []string)`

SetRejectExpansionAttribute sets RejectExpansionAttribute field to given value.

### HasRejectExpansionAttribute

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasRejectExpansionAttribute() bool`

HasRejectExpansionAttribute returns a boolean if a field has been set.

### GetAllowedControl

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAllowedControl() []EnumhttpServletExtensionAllowedControlProp`

GetAllowedControl returns the AllowedControl field if non-nil, zero value otherwise.

### GetAllowedControlOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetAllowedControlOk() (*[]EnumhttpServletExtensionAllowedControlProp, bool)`

GetAllowedControlOk returns a tuple with the AllowedControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedControl

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetAllowedControl(v []EnumhttpServletExtensionAllowedControlProp)`

SetAllowedControl sets AllowedControl field to given value.

### HasAllowedControl

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasAllowedControl() bool`

HasAllowedControl returns a boolean if a field has been set.

### GetDescription

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *DirectoryRestApiHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *DirectoryRestApiHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *DirectoryRestApiHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


