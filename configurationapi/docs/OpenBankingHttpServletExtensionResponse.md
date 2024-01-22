# OpenBankingHttpServletExtensionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumopenBankingHttpServletExtensionSchemaUrn**](EnumopenBankingHttpServletExtensionSchemaUrn.md) |  | 
**Id** | **string** | Name of the HTTP Servlet Extension | 
**CorrelationIDResponseHeader** | Pointer to **string** | Specifies the name of the HTTP header that will contain a correlation ID value. This header will be used both in downstream requests to the Consent API and in responses to the Open Banking client. The value \&quot;x-fapi-interaction-id\&quot; is recommended. | [optional] 
**FapiFinancialID** | Pointer to **string** | The unique financial id of the ASPSP associated with this Open Banking service. | [optional] 
**AccessTokenValidator** | Pointer to **[]string** | If specified, the Access Token Validator(s) that may be used to validate access tokens for requests submitted to this Open Banking HTTP Servlet Extension. | [optional] 
**PathPrefix** | Pointer to **string** | An optional ASPSP-specific path prefix to include in the base URI path. If specified with the value \&quot;myPrefix\&quot;, for example, the base path will be /myPrefix/open-banking/v1.1/. | [optional] 
**ConsentServer** | Pointer to **string** | Specifies the PingDirectory instance that is hosting the Consent API for storage and retrieval of Account Requests. | [optional] 
**ConsentDefinitionID** | **string** | The name/id of the consent definition, as defined in the consent-server, that is used for storing Account Requests. | 
**Description** | Pointer to **string** | A description for this HTTP Servlet Extension | [optional] 
**CrossOriginPolicy** | Pointer to **string** | The cross-origin request policy to use for the HTTP Servlet Extension. | [optional] 
**ResponseHeader** | Pointer to **[]string** | Specifies HTTP header fields and values added to response headers for all requests. | [optional] 

## Methods

### NewOpenBankingHttpServletExtensionResponse

`func NewOpenBankingHttpServletExtensionResponse(schemas []EnumopenBankingHttpServletExtensionSchemaUrn, id string, consentDefinitionID string, ) *OpenBankingHttpServletExtensionResponse`

NewOpenBankingHttpServletExtensionResponse instantiates a new OpenBankingHttpServletExtensionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenBankingHttpServletExtensionResponseWithDefaults

`func NewOpenBankingHttpServletExtensionResponseWithDefaults() *OpenBankingHttpServletExtensionResponse`

NewOpenBankingHttpServletExtensionResponseWithDefaults instantiates a new OpenBankingHttpServletExtensionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *OpenBankingHttpServletExtensionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OpenBankingHttpServletExtensionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OpenBankingHttpServletExtensionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OpenBankingHttpServletExtensionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *OpenBankingHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *OpenBankingHttpServletExtensionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *OpenBankingHttpServletExtensionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *OpenBankingHttpServletExtensionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *OpenBankingHttpServletExtensionResponse) GetSchemas() []EnumopenBankingHttpServletExtensionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *OpenBankingHttpServletExtensionResponse) GetSchemasOk() (*[]EnumopenBankingHttpServletExtensionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *OpenBankingHttpServletExtensionResponse) SetSchemas(v []EnumopenBankingHttpServletExtensionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *OpenBankingHttpServletExtensionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenBankingHttpServletExtensionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenBankingHttpServletExtensionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCorrelationIDResponseHeader

`func (o *OpenBankingHttpServletExtensionResponse) GetCorrelationIDResponseHeader() string`

GetCorrelationIDResponseHeader returns the CorrelationIDResponseHeader field if non-nil, zero value otherwise.

### GetCorrelationIDResponseHeaderOk

`func (o *OpenBankingHttpServletExtensionResponse) GetCorrelationIDResponseHeaderOk() (*string, bool)`

GetCorrelationIDResponseHeaderOk returns a tuple with the CorrelationIDResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationIDResponseHeader

`func (o *OpenBankingHttpServletExtensionResponse) SetCorrelationIDResponseHeader(v string)`

SetCorrelationIDResponseHeader sets CorrelationIDResponseHeader field to given value.

### HasCorrelationIDResponseHeader

`func (o *OpenBankingHttpServletExtensionResponse) HasCorrelationIDResponseHeader() bool`

HasCorrelationIDResponseHeader returns a boolean if a field has been set.

### GetFapiFinancialID

`func (o *OpenBankingHttpServletExtensionResponse) GetFapiFinancialID() string`

GetFapiFinancialID returns the FapiFinancialID field if non-nil, zero value otherwise.

### GetFapiFinancialIDOk

`func (o *OpenBankingHttpServletExtensionResponse) GetFapiFinancialIDOk() (*string, bool)`

GetFapiFinancialIDOk returns a tuple with the FapiFinancialID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFapiFinancialID

`func (o *OpenBankingHttpServletExtensionResponse) SetFapiFinancialID(v string)`

SetFapiFinancialID sets FapiFinancialID field to given value.

### HasFapiFinancialID

`func (o *OpenBankingHttpServletExtensionResponse) HasFapiFinancialID() bool`

HasFapiFinancialID returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *OpenBankingHttpServletExtensionResponse) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *OpenBankingHttpServletExtensionResponse) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *OpenBankingHttpServletExtensionResponse) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *OpenBankingHttpServletExtensionResponse) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetPathPrefix

`func (o *OpenBankingHttpServletExtensionResponse) GetPathPrefix() string`

GetPathPrefix returns the PathPrefix field if non-nil, zero value otherwise.

### GetPathPrefixOk

`func (o *OpenBankingHttpServletExtensionResponse) GetPathPrefixOk() (*string, bool)`

GetPathPrefixOk returns a tuple with the PathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefix

`func (o *OpenBankingHttpServletExtensionResponse) SetPathPrefix(v string)`

SetPathPrefix sets PathPrefix field to given value.

### HasPathPrefix

`func (o *OpenBankingHttpServletExtensionResponse) HasPathPrefix() bool`

HasPathPrefix returns a boolean if a field has been set.

### GetConsentServer

`func (o *OpenBankingHttpServletExtensionResponse) GetConsentServer() string`

GetConsentServer returns the ConsentServer field if non-nil, zero value otherwise.

### GetConsentServerOk

`func (o *OpenBankingHttpServletExtensionResponse) GetConsentServerOk() (*string, bool)`

GetConsentServerOk returns a tuple with the ConsentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentServer

`func (o *OpenBankingHttpServletExtensionResponse) SetConsentServer(v string)`

SetConsentServer sets ConsentServer field to given value.

### HasConsentServer

`func (o *OpenBankingHttpServletExtensionResponse) HasConsentServer() bool`

HasConsentServer returns a boolean if a field has been set.

### GetConsentDefinitionID

`func (o *OpenBankingHttpServletExtensionResponse) GetConsentDefinitionID() string`

GetConsentDefinitionID returns the ConsentDefinitionID field if non-nil, zero value otherwise.

### GetConsentDefinitionIDOk

`func (o *OpenBankingHttpServletExtensionResponse) GetConsentDefinitionIDOk() (*string, bool)`

GetConsentDefinitionIDOk returns a tuple with the ConsentDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentDefinitionID

`func (o *OpenBankingHttpServletExtensionResponse) SetConsentDefinitionID(v string)`

SetConsentDefinitionID sets ConsentDefinitionID field to given value.


### GetDescription

`func (o *OpenBankingHttpServletExtensionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenBankingHttpServletExtensionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenBankingHttpServletExtensionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenBankingHttpServletExtensionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCrossOriginPolicy

`func (o *OpenBankingHttpServletExtensionResponse) GetCrossOriginPolicy() string`

GetCrossOriginPolicy returns the CrossOriginPolicy field if non-nil, zero value otherwise.

### GetCrossOriginPolicyOk

`func (o *OpenBankingHttpServletExtensionResponse) GetCrossOriginPolicyOk() (*string, bool)`

GetCrossOriginPolicyOk returns a tuple with the CrossOriginPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossOriginPolicy

`func (o *OpenBankingHttpServletExtensionResponse) SetCrossOriginPolicy(v string)`

SetCrossOriginPolicy sets CrossOriginPolicy field to given value.

### HasCrossOriginPolicy

`func (o *OpenBankingHttpServletExtensionResponse) HasCrossOriginPolicy() bool`

HasCrossOriginPolicy returns a boolean if a field has been set.

### GetResponseHeader

`func (o *OpenBankingHttpServletExtensionResponse) GetResponseHeader() []string`

GetResponseHeader returns the ResponseHeader field if non-nil, zero value otherwise.

### GetResponseHeaderOk

`func (o *OpenBankingHttpServletExtensionResponse) GetResponseHeaderOk() (*[]string, bool)`

GetResponseHeaderOk returns a tuple with the ResponseHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeader

`func (o *OpenBankingHttpServletExtensionResponse) SetResponseHeader(v []string)`

SetResponseHeader sets ResponseHeader field to given value.

### HasResponseHeader

`func (o *OpenBankingHttpServletExtensionResponse) HasResponseHeader() bool`

HasResponseHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


