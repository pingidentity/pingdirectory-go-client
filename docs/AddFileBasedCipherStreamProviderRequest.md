# AddFileBasedCipherStreamProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the new Cipher Stream Provider | 
**Schemas** | [**[]EnumfileBasedCipherStreamProviderSchemaUrn**](EnumfileBasedCipherStreamProviderSchemaUrn.md) |  | 
**PasswordFile** | **string** | The path to the file containing the password to use when generating ciphers. | 
**WaitForPasswordFile** | Pointer to **bool** | Indicates whether the server should wait for the password file to become available if it does not exist. | [optional] 
**Description** | Pointer to **string** | A description for this Cipher Stream Provider | [optional] 
**Enabled** | **bool** | Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server. | 

## Methods

### NewAddFileBasedCipherStreamProviderRequest

`func NewAddFileBasedCipherStreamProviderRequest(providerName string, schemas []EnumfileBasedCipherStreamProviderSchemaUrn, passwordFile string, enabled bool, ) *AddFileBasedCipherStreamProviderRequest`

NewAddFileBasedCipherStreamProviderRequest instantiates a new AddFileBasedCipherStreamProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFileBasedCipherStreamProviderRequestWithDefaults

`func NewAddFileBasedCipherStreamProviderRequestWithDefaults() *AddFileBasedCipherStreamProviderRequest`

NewAddFileBasedCipherStreamProviderRequestWithDefaults instantiates a new AddFileBasedCipherStreamProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *AddFileBasedCipherStreamProviderRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *AddFileBasedCipherStreamProviderRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *AddFileBasedCipherStreamProviderRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetSchemas

`func (o *AddFileBasedCipherStreamProviderRequest) GetSchemas() []EnumfileBasedCipherStreamProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFileBasedCipherStreamProviderRequest) GetSchemasOk() (*[]EnumfileBasedCipherStreamProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFileBasedCipherStreamProviderRequest) SetSchemas(v []EnumfileBasedCipherStreamProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPasswordFile

`func (o *AddFileBasedCipherStreamProviderRequest) GetPasswordFile() string`

GetPasswordFile returns the PasswordFile field if non-nil, zero value otherwise.

### GetPasswordFileOk

`func (o *AddFileBasedCipherStreamProviderRequest) GetPasswordFileOk() (*string, bool)`

GetPasswordFileOk returns a tuple with the PasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFile

`func (o *AddFileBasedCipherStreamProviderRequest) SetPasswordFile(v string)`

SetPasswordFile sets PasswordFile field to given value.


### GetWaitForPasswordFile

`func (o *AddFileBasedCipherStreamProviderRequest) GetWaitForPasswordFile() bool`

GetWaitForPasswordFile returns the WaitForPasswordFile field if non-nil, zero value otherwise.

### GetWaitForPasswordFileOk

`func (o *AddFileBasedCipherStreamProviderRequest) GetWaitForPasswordFileOk() (*bool, bool)`

GetWaitForPasswordFileOk returns a tuple with the WaitForPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForPasswordFile

`func (o *AddFileBasedCipherStreamProviderRequest) SetWaitForPasswordFile(v bool)`

SetWaitForPasswordFile sets WaitForPasswordFile field to given value.

### HasWaitForPasswordFile

`func (o *AddFileBasedCipherStreamProviderRequest) HasWaitForPasswordFile() bool`

HasWaitForPasswordFile returns a boolean if a field has been set.

### GetDescription

`func (o *AddFileBasedCipherStreamProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFileBasedCipherStreamProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFileBasedCipherStreamProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFileBasedCipherStreamProviderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddFileBasedCipherStreamProviderRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddFileBasedCipherStreamProviderRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddFileBasedCipherStreamProviderRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


