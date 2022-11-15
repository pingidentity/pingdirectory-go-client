# DigestMd5SaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdigestMd5SaslMechanismHandlerSchemaUrn**](EnumdigestMd5SaslMechanismHandlerSchemaUrn.md) |  | 
**Realm** | Pointer to **string** | Specifies the realm that is to be used by the server for DIGEST-MD5 authentication. | [optional] 
**IdentityMapper** | **string** | Specifies the name of the identity mapper that is to be used with this SASL mechanism handler to match the authentication or authorization ID included in the SASL bind request to the corresponding user in the directory. | 
**ServerFqdn** | Pointer to **string** | Specifies the DNS-resolvable fully-qualified domain name for the server that is used when validating the digest-uri parameter during the authentication process. | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewDigestMd5SaslMechanismHandlerResponse

`func NewDigestMd5SaslMechanismHandlerResponse(schemas []EnumdigestMd5SaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, ) *DigestMd5SaslMechanismHandlerResponse`

NewDigestMd5SaslMechanismHandlerResponse instantiates a new DigestMd5SaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigestMd5SaslMechanismHandlerResponseWithDefaults

`func NewDigestMd5SaslMechanismHandlerResponseWithDefaults() *DigestMd5SaslMechanismHandlerResponse`

NewDigestMd5SaslMechanismHandlerResponseWithDefaults instantiates a new DigestMd5SaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DigestMd5SaslMechanismHandlerResponse) GetSchemas() []EnumdigestMd5SaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DigestMd5SaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumdigestMd5SaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DigestMd5SaslMechanismHandlerResponse) SetSchemas(v []EnumdigestMd5SaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRealm

`func (o *DigestMd5SaslMechanismHandlerResponse) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *DigestMd5SaslMechanismHandlerResponse) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *DigestMd5SaslMechanismHandlerResponse) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *DigestMd5SaslMechanismHandlerResponse) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *DigestMd5SaslMechanismHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *DigestMd5SaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *DigestMd5SaslMechanismHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetServerFqdn

`func (o *DigestMd5SaslMechanismHandlerResponse) GetServerFqdn() string`

GetServerFqdn returns the ServerFqdn field if non-nil, zero value otherwise.

### GetServerFqdnOk

`func (o *DigestMd5SaslMechanismHandlerResponse) GetServerFqdnOk() (*string, bool)`

GetServerFqdnOk returns a tuple with the ServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFqdn

`func (o *DigestMd5SaslMechanismHandlerResponse) SetServerFqdn(v string)`

SetServerFqdn sets ServerFqdn field to given value.

### HasServerFqdn

`func (o *DigestMd5SaslMechanismHandlerResponse) HasServerFqdn() bool`

HasServerFqdn returns a boolean if a field has been set.

### GetDescription

`func (o *DigestMd5SaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DigestMd5SaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DigestMd5SaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DigestMd5SaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DigestMd5SaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DigestMd5SaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DigestMd5SaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


