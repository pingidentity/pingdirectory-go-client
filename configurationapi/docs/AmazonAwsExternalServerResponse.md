# AmazonAwsExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumamazonAwsExternalServerSchemaUrn**](EnumamazonAwsExternalServerSchemaUrn.md) |  | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the AWS service. | [optional] 
**AuthenticationMethod** | Pointer to [**EnumexternalServerAmazonAwsAuthenticationMethodProp**](EnumexternalServerAmazonAwsAuthenticationMethodProp.md) |  | [optional] 
**AwsAccessKeyID** | Pointer to **string** | The access key ID that will be used if authentication should use an access key. If this is provided, then an aws-secret-access-key must also be provided. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | The secret access key that will be used if authentication should use an access key. If this is provided, then an aws-access-key-id must also be provided. | [optional] 
**AwsRegionName** | **string** | The name of the AWS region containing the resources that will be accessed. | 
**Description** | Pointer to **string** | A description for this External Server | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the External Server | 

## Methods

### NewAmazonAwsExternalServerResponse

`func NewAmazonAwsExternalServerResponse(schemas []EnumamazonAwsExternalServerSchemaUrn, awsRegionName string, id string, ) *AmazonAwsExternalServerResponse`

NewAmazonAwsExternalServerResponse instantiates a new AmazonAwsExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAwsExternalServerResponseWithDefaults

`func NewAmazonAwsExternalServerResponseWithDefaults() *AmazonAwsExternalServerResponse`

NewAmazonAwsExternalServerResponseWithDefaults instantiates a new AmazonAwsExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AmazonAwsExternalServerResponse) GetSchemas() []EnumamazonAwsExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AmazonAwsExternalServerResponse) GetSchemasOk() (*[]EnumamazonAwsExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AmazonAwsExternalServerResponse) SetSchemas(v []EnumamazonAwsExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHttpProxyExternalServer

`func (o *AmazonAwsExternalServerResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *AmazonAwsExternalServerResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *AmazonAwsExternalServerResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *AmazonAwsExternalServerResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AmazonAwsExternalServerResponse) GetAuthenticationMethod() EnumexternalServerAmazonAwsAuthenticationMethodProp`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AmazonAwsExternalServerResponse) GetAuthenticationMethodOk() (*EnumexternalServerAmazonAwsAuthenticationMethodProp, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AmazonAwsExternalServerResponse) SetAuthenticationMethod(v EnumexternalServerAmazonAwsAuthenticationMethodProp)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *AmazonAwsExternalServerResponse) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetAwsAccessKeyID

`func (o *AmazonAwsExternalServerResponse) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AmazonAwsExternalServerResponse) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AmazonAwsExternalServerResponse) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AmazonAwsExternalServerResponse) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *AmazonAwsExternalServerResponse) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AmazonAwsExternalServerResponse) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AmazonAwsExternalServerResponse) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *AmazonAwsExternalServerResponse) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsRegionName

`func (o *AmazonAwsExternalServerResponse) GetAwsRegionName() string`

GetAwsRegionName returns the AwsRegionName field if non-nil, zero value otherwise.

### GetAwsRegionNameOk

`func (o *AmazonAwsExternalServerResponse) GetAwsRegionNameOk() (*string, bool)`

GetAwsRegionNameOk returns a tuple with the AwsRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionName

`func (o *AmazonAwsExternalServerResponse) SetAwsRegionName(v string)`

SetAwsRegionName sets AwsRegionName field to given value.


### GetDescription

`func (o *AmazonAwsExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonAwsExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonAwsExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AmazonAwsExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AmazonAwsExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AmazonAwsExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AmazonAwsExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AmazonAwsExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AmazonAwsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AmazonAwsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AmazonAwsExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AmazonAwsExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AmazonAwsExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonAwsExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonAwsExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


