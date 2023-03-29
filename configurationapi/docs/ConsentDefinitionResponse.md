# ConsentDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Consent Definition | 
**Schemas** | Pointer to [**[]EnumconsentDefinitionSchemaUrn**](EnumconsentDefinitionSchemaUrn.md) |  | [optional] 
**UniqueID** | **string** | A version-independent unique identifier for this Consent Definition. | 
**DisplayName** | Pointer to **string** | A human-readable display name for this Consent Definition. | [optional] 
**Parameter** | Pointer to **[]string** | Optional parameters for this Consent Definition. | [optional] 
**Description** | Pointer to **string** | A description for this Consent Definition | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewConsentDefinitionResponse

`func NewConsentDefinitionResponse(id string, uniqueID string, ) *ConsentDefinitionResponse`

NewConsentDefinitionResponse instantiates a new ConsentDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentDefinitionResponseWithDefaults

`func NewConsentDefinitionResponseWithDefaults() *ConsentDefinitionResponse`

NewConsentDefinitionResponseWithDefaults instantiates a new ConsentDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsentDefinitionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsentDefinitionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsentDefinitionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ConsentDefinitionResponse) GetSchemas() []EnumconsentDefinitionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConsentDefinitionResponse) GetSchemasOk() (*[]EnumconsentDefinitionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConsentDefinitionResponse) SetSchemas(v []EnumconsentDefinitionSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConsentDefinitionResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetUniqueID

`func (o *ConsentDefinitionResponse) GetUniqueID() string`

GetUniqueID returns the UniqueID field if non-nil, zero value otherwise.

### GetUniqueIDOk

`func (o *ConsentDefinitionResponse) GetUniqueIDOk() (*string, bool)`

GetUniqueIDOk returns a tuple with the UniqueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueID

`func (o *ConsentDefinitionResponse) SetUniqueID(v string)`

SetUniqueID sets UniqueID field to given value.


### GetDisplayName

`func (o *ConsentDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConsentDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConsentDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConsentDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetParameter

`func (o *ConsentDefinitionResponse) GetParameter() []string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ConsentDefinitionResponse) GetParameterOk() (*[]string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ConsentDefinitionResponse) SetParameter(v []string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ConsentDefinitionResponse) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetDescription

`func (o *ConsentDefinitionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsentDefinitionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsentDefinitionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsentDefinitionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *ConsentDefinitionResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsentDefinitionResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsentDefinitionResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsentDefinitionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentDefinitionResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ConsentDefinitionResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentDefinitionResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ConsentDefinitionResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


