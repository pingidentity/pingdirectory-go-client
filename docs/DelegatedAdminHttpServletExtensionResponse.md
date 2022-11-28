# DelegatedAdminHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumdelegatedAdminHttpServletExtensionSchemaUrn**](EnumdelegatedAdminHttpServletExtensionSchemaUrn.md) |  | 
**Id** | Pointer to **string** | Name of the HTTP Servlet Extension | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Enables HTTP Basic authentication, using a username and password. The Identity Mapper specified by the identity-mapper property will be used to map the username to a DN. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the Identity Mapper that is to be used for associating user entries with basic authentication user names. | [optional] 
**AccessTokenValidator** | Pointer to **[]string** | If specified, the Access Token Validator(s) that may be used to validate access tokens for requests submitted to this Delegated Admin HTTP Servlet Extension. | [optional] 
**AccessTokenScope** | Pointer to **string** | The name of a scope that must be present in an access token accepted by the Delegated Admin HTTP Servlet Extension. | [optional] 
**Audience** | Pointer to **string** | A string or URI that identifies the Delegated Admin HTTP Servlet Extension in the context of OAuth2 authorization. | [optional] 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are \&quot;Correlation-Id\&quot;, \&quot;X-Amzn-Trace-Id\&quot;, and \&quot;X-Request-Id\&quot;. | [optional] 

## Methods

### NewDelegatedAdminHttpServletExtensionResponse

`func NewDelegatedAdminHttpServletExtensionResponse(schemas []EnumdelegatedAdminHttpServletExtensionSchemaUrn, ) *DelegatedAdminHttpServletExtensionResponse`

NewDelegatedAdminHttpServletExtensionResponse instantiates a new DelegatedAdminHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminHttpServletExtensionResponseWithDefaults

`func NewDelegatedAdminHttpServletExtensionResponseWithDefaults() *DelegatedAdminHttpServletExtensionResponse`

NewDelegatedAdminHttpServletExtensionResponseWithDefaults instantiates a new DelegatedAdminHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DelegatedAdminHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DelegatedAdminHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DelegatedAdminHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DelegatedAdminHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *DelegatedAdminHttpServletExtensionResponse) GetSchemas() []EnumdelegatedAdminHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetSchemasOk() (*[]EnumdelegatedAdminHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelegatedAdminHttpServletExtensionResponse) SetSchemas(v []EnumdelegatedAdminHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *DelegatedAdminHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatedAdminHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DelegatedAdminHttpServletExtensionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *DelegatedAdminHttpServletExtensionResponse) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *DelegatedAdminHttpServletExtensionResponse) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *DelegatedAdminHttpServletExtensionResponse) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *DelegatedAdminHttpServletExtensionResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *DelegatedAdminHttpServletExtensionResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *DelegatedAdminHttpServletExtensionResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *DelegatedAdminHttpServletExtensionResponse) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *DelegatedAdminHttpServletExtensionResponse) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetAccessTokenScope

`func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenScope() string`

GetAccessTokenScope returns the AccessTokenScope field if non-nil, zero value otherwise.

### GetAccessTokenScopeOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetAccessTokenScopeOk() (*string, bool)`

GetAccessTokenScopeOk returns a tuple with the AccessTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenScope

`func (o *DelegatedAdminHttpServletExtensionResponse) SetAccessTokenScope(v string)`

SetAccessTokenScope sets AccessTokenScope field to given value.

### HasAccessTokenScope

`func (o *DelegatedAdminHttpServletExtensionResponse) HasAccessTokenScope() bool`

HasAccessTokenScope returns a boolean if a field has been set.

### GetAudience

`func (o *DelegatedAdminHttpServletExtensionResponse) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *DelegatedAdminHttpServletExtensionResponse) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *DelegatedAdminHttpServletExtensionResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetDescription

`func (o *DelegatedAdminHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAdminHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAdminHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *DelegatedAdminHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *DelegatedAdminHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *DelegatedAdminHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *DelegatedAdminHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *DelegatedAdminHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *DelegatedAdminHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.

### GetCorrelationIDResponseHeader

`func (o *DelegatedAdminHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *DelegatedAdminHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *DelegatedAdminHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *DelegatedAdminHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


