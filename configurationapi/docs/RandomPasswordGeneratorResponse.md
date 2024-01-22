# RandomPasswordGeneratorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumrandomPasswordGeneratorSchemaUrn**](EnumrandomPasswordGeneratorSchemaUrn.md) |  | 
**PasswordCharacterSet** | **[]string** | Specifies one or more named character sets. | 
**PasswordFormat** | **string** | Specifies the format to use for the generated password. | 
**Description** | Pointer to **string** | A description for this Password Generator | [optional] 
**Enabled** | **bool** | Indicates whether the Password Generator is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Password Generator | 

## Methods

### NewRandomPasswordGeneratorResponse

`func NewRandomPasswordGeneratorResponse(schemas []EnumrandomPasswordGeneratorSchemaUrn, passwordCharacterSet []string, passwordFormat string, enabled bool, id string, ) *RandomPasswordGeneratorResponse`

NewRandomPasswordGeneratorResponse instantiates a new RandomPasswordGeneratorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomPasswordGeneratorResponseWithDefaults

`func NewRandomPasswordGeneratorResponseWithDefaults() *RandomPasswordGeneratorResponse`

NewRandomPasswordGeneratorResponseWithDefaults instantiates a new RandomPasswordGeneratorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *RandomPasswordGeneratorResponse) GetSchemas() []EnumrandomPasswordGeneratorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RandomPasswordGeneratorResponse) GetSchemasOk() (*[]EnumrandomPasswordGeneratorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RandomPasswordGeneratorResponse) SetSchemas(v []EnumrandomPasswordGeneratorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordCharacterSet

`func (o *RandomPasswordGeneratorResponse) GetPasswordCharacterSet() []string`

GetPasswordCharacterSet returns the PasswordCharacterSet field if non-nil, zero value otherwise.

### GetPasswordCharacterSetOk

`func (o *RandomPasswordGeneratorResponse) GetPasswordCharacterSetOk() (*[]string, bool)`

GetPasswordCharacterSetOk returns a tuple with the PasswordCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCharacterSet

`func (o *RandomPasswordGeneratorResponse) SetPasswordCharacterSet(v []string)`

SetPasswordCharacterSet sets PasswordCharacterSet field to given value.


### GetPasswordFormat

`func (o *RandomPasswordGeneratorResponse) GetPasswordFormat() string`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *RandomPasswordGeneratorResponse) GetPasswordFormatOk() (*string, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *RandomPasswordGeneratorResponse) SetPasswordFormat(v string)`

SetPasswordFormat sets PasswordFormat field to given value.


### GetDescription

`func (o *RandomPasswordGeneratorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RandomPasswordGeneratorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RandomPasswordGeneratorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RandomPasswordGeneratorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RandomPasswordGeneratorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RandomPasswordGeneratorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RandomPasswordGeneratorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *RandomPasswordGeneratorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RandomPasswordGeneratorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RandomPasswordGeneratorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RandomPasswordGeneratorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *RandomPasswordGeneratorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *RandomPasswordGeneratorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *RandomPasswordGeneratorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *RandomPasswordGeneratorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *RandomPasswordGeneratorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RandomPasswordGeneratorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RandomPasswordGeneratorResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


