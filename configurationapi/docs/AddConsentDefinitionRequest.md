# AddConsentDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumconsentDefinitionSchemaUrn**](EnumconsentDefinitionSchemaUrn.md) |  | [optional] 
**UniqueID** | **string** | A version-independent unique identifier for this Consent Definition. | 
**DisplayName** | Pointer to **string** | A human-readable display name for this Consent Definition. | [optional] 
**Parameter** | Pointer to **[]string** | Optional parameters for this Consent Definition. | [optional] 
**Description** | Pointer to **string** | A description for this Consent Definition | [optional] 
**DefinitionName** | **string** | Name of the new Consent Definition | 

## Methods

### NewAddConsentDefinitionRequest

`func NewAddConsentDefinitionRequest(uniqueID string, definitionName string, ) *AddConsentDefinitionRequest`

NewAddConsentDefinitionRequest instantiates a new AddConsentDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConsentDefinitionRequestWithDefaults

`func NewAddConsentDefinitionRequestWithDefaults() *AddConsentDefinitionRequest`

NewAddConsentDefinitionRequestWithDefaults instantiates a new AddConsentDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddConsentDefinitionRequest) GetSchemas() []EnumconsentDefinitionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddConsentDefinitionRequest) GetSchemasOk() (*[]EnumconsentDefinitionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddConsentDefinitionRequest) SetSchemas(v []EnumconsentDefinitionSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddConsentDefinitionRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetUniqueID

`func (o *AddConsentDefinitionRequest) GetUniqueID() string`

GetUniqueID returns the UniqueID field if non-nil, zero value otherwise.

### GetUniqueIDOk

`func (o *AddConsentDefinitionRequest) GetUniqueIDOk() (*string, bool)`

GetUniqueIDOk returns a tuple with the UniqueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueID

`func (o *AddConsentDefinitionRequest) SetUniqueID(v string)`

SetUniqueID sets UniqueID field to given value.


### GetDisplayName

`func (o *AddConsentDefinitionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddConsentDefinitionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddConsentDefinitionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddConsentDefinitionRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetParameter

`func (o *AddConsentDefinitionRequest) GetParameter() []string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *AddConsentDefinitionRequest) GetParameterOk() (*[]string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *AddConsentDefinitionRequest) SetParameter(v []string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *AddConsentDefinitionRequest) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetDescription

`func (o *AddConsentDefinitionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddConsentDefinitionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddConsentDefinitionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddConsentDefinitionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefinitionName

`func (o *AddConsentDefinitionRequest) GetDefinitionName() string`

GetDefinitionName returns the DefinitionName field if non-nil, zero value otherwise.

### GetDefinitionNameOk

`func (o *AddConsentDefinitionRequest) GetDefinitionNameOk() (*string, bool)`

GetDefinitionNameOk returns a tuple with the DefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionName

`func (o *AddConsentDefinitionRequest) SetDefinitionName(v string)`

SetDefinitionName sets DefinitionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


