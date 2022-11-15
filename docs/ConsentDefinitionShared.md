# ConsentDefinitionShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumconsentDefinitionSchemaUrn**](EnumconsentDefinitionSchemaUrn.md) |  | [optional] 
**UniqueID** | **string** | A version-independent unique identifier for this Consent Definition. | 
**DisplayName** | Pointer to **string** | A human-readable display name for this Consent Definition. | [optional] 
**Parameter** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Consent Definition | [optional] 

## Methods

### NewConsentDefinitionShared

`func NewConsentDefinitionShared(uniqueID string, ) *ConsentDefinitionShared`

NewConsentDefinitionShared instantiates a new ConsentDefinitionShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentDefinitionSharedWithDefaults

`func NewConsentDefinitionSharedWithDefaults() *ConsentDefinitionShared`

NewConsentDefinitionSharedWithDefaults instantiates a new ConsentDefinitionShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ConsentDefinitionShared) GetSchemas() []EnumconsentDefinitionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ConsentDefinitionShared) GetSchemasOk() (*[]EnumconsentDefinitionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ConsentDefinitionShared) SetSchemas(v []EnumconsentDefinitionSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ConsentDefinitionShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetUniqueID

`func (o *ConsentDefinitionShared) GetUniqueID() string`

GetUniqueID returns the UniqueID field if non-nil, zero value otherwise.

### GetUniqueIDOk

`func (o *ConsentDefinitionShared) GetUniqueIDOk() (*string, bool)`

GetUniqueIDOk returns a tuple with the UniqueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueID

`func (o *ConsentDefinitionShared) SetUniqueID(v string)`

SetUniqueID sets UniqueID field to given value.


### GetDisplayName

`func (o *ConsentDefinitionShared) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConsentDefinitionShared) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConsentDefinitionShared) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConsentDefinitionShared) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetParameter

`func (o *ConsentDefinitionShared) GetParameter() []string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ConsentDefinitionShared) GetParameterOk() (*[]string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ConsentDefinitionShared) SetParameter(v []string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ConsentDefinitionShared) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetDescription

`func (o *ConsentDefinitionShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsentDefinitionShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsentDefinitionShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsentDefinitionShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


