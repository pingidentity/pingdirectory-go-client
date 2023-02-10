# ChangelogPasswordEncryptionPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumchangelogPasswordEncryptionPluginSchemaUrn**](EnumchangelogPasswordEncryptionPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**ChangelogPasswordEncryptionKey** | Pointer to **string** | A passphrase that may be used to generate the key for encrypting passwords stored in the changelog. The same passphrase also needs to be set (either through the \&quot;changelog-password-decryption-key\&quot; property or the \&quot;changelog-password-decryption-key-passphrase-provider\&quot; property) in the Global Sync Configuration in the Data Sync Server. | [optional] 
**ChangelogPasswordEncryptionKeyPassphraseProvider** | Pointer to **string** | A passphrase provider that may be used to obtain the passphrase that will be used to generate the key for encrypting passwords stored in the changelog. The same passphrase also needs to be set (either through the \&quot;changelog-password-decryption-key\&quot; property or the \&quot;changelog-password-decryption-key-passphrase-provider\&quot; property) in the Global Sync Configuration in the Data Sync Server. | [optional] 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewChangelogPasswordEncryptionPluginResponse

`func NewChangelogPasswordEncryptionPluginResponse(schemas []EnumchangelogPasswordEncryptionPluginSchemaUrn, id string, pluginType []EnumpluginPluginTypeProp, enabled bool, ) *ChangelogPasswordEncryptionPluginResponse`

NewChangelogPasswordEncryptionPluginResponse instantiates a new ChangelogPasswordEncryptionPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogPasswordEncryptionPluginResponseWithDefaults

`func NewChangelogPasswordEncryptionPluginResponseWithDefaults() *ChangelogPasswordEncryptionPluginResponse`

NewChangelogPasswordEncryptionPluginResponseWithDefaults instantiates a new ChangelogPasswordEncryptionPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ChangelogPasswordEncryptionPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChangelogPasswordEncryptionPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChangelogPasswordEncryptionPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ChangelogPasswordEncryptionPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ChangelogPasswordEncryptionPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ChangelogPasswordEncryptionPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ChangelogPasswordEncryptionPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *ChangelogPasswordEncryptionPluginResponse) GetSchemas() []EnumchangelogPasswordEncryptionPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetSchemasOk() (*[]EnumchangelogPasswordEncryptionPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ChangelogPasswordEncryptionPluginResponse) SetSchemas(v []EnumchangelogPasswordEncryptionPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ChangelogPasswordEncryptionPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangelogPasswordEncryptionPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetChangelogPasswordEncryptionKey

`func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKey() string`

GetChangelogPasswordEncryptionKey returns the ChangelogPasswordEncryptionKey field if non-nil, zero value otherwise.

### GetChangelogPasswordEncryptionKeyOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKeyOk() (*string, bool)`

GetChangelogPasswordEncryptionKeyOk returns a tuple with the ChangelogPasswordEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPasswordEncryptionKey

`func (o *ChangelogPasswordEncryptionPluginResponse) SetChangelogPasswordEncryptionKey(v string)`

SetChangelogPasswordEncryptionKey sets ChangelogPasswordEncryptionKey field to given value.

### HasChangelogPasswordEncryptionKey

`func (o *ChangelogPasswordEncryptionPluginResponse) HasChangelogPasswordEncryptionKey() bool`

HasChangelogPasswordEncryptionKey returns a boolean if a field has been set.

### GetChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKeyPassphraseProvider() string`

GetChangelogPasswordEncryptionKeyPassphraseProvider returns the ChangelogPasswordEncryptionKeyPassphraseProvider field if non-nil, zero value otherwise.

### GetChangelogPasswordEncryptionKeyPassphraseProviderOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetChangelogPasswordEncryptionKeyPassphraseProviderOk() (*string, bool)`

GetChangelogPasswordEncryptionKeyPassphraseProviderOk returns a tuple with the ChangelogPasswordEncryptionKeyPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *ChangelogPasswordEncryptionPluginResponse) SetChangelogPasswordEncryptionKeyPassphraseProvider(v string)`

SetChangelogPasswordEncryptionKeyPassphraseProvider sets ChangelogPasswordEncryptionKeyPassphraseProvider field to given value.

### HasChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *ChangelogPasswordEncryptionPluginResponse) HasChangelogPasswordEncryptionKeyPassphraseProvider() bool`

HasChangelogPasswordEncryptionKeyPassphraseProvider returns a boolean if a field has been set.

### GetPluginType

`func (o *ChangelogPasswordEncryptionPluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *ChangelogPasswordEncryptionPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetDescription

`func (o *ChangelogPasswordEncryptionPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangelogPasswordEncryptionPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangelogPasswordEncryptionPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ChangelogPasswordEncryptionPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ChangelogPasswordEncryptionPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *ChangelogPasswordEncryptionPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *ChangelogPasswordEncryptionPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *ChangelogPasswordEncryptionPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *ChangelogPasswordEncryptionPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


