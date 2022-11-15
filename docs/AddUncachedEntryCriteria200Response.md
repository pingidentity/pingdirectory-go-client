# AddUncachedEntryCriteria200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Uncached Entry Criteria | 
**Schemas** | [**[]EnumthirdPartyUncachedEntryCriteriaSchemaUrn**](EnumthirdPartyUncachedEntryCriteriaSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Uncached Entry Criteria | [optional] 
**Enabled** | **bool** | Indicates whether this Uncached Entry Criteria is enabled for use in the server. | 
**AccessTimeThreshold** | **string** | Specifies the maximum length of time that has passed since an entry was last accessed that it should still be included in the id2entry database. Entries that have not been accessed in more than this length of time may be written into the uncached-id2entry database. | 
**Filter** | **string** | Specifies the search filter that should be used to differentiate entries into cached and uncached sets. | 
**FilterIdentifiesUncachedEntries** | **bool** | Indicates whether the associated filter identifies those entries which should be stored in the uncached-id2entry database (if true) or entries which should be stored in the id2entry database (if false). | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Uncached Entry Criteria. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Uncached Entry Criteria. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddUncachedEntryCriteria200Response

`func NewAddUncachedEntryCriteria200Response(id string, schemas []EnumthirdPartyUncachedEntryCriteriaSchemaUrn, enabled bool, accessTimeThreshold string, filter string, filterIdentifiesUncachedEntries bool, scriptClass string, extensionClass string, ) *AddUncachedEntryCriteria200Response`

NewAddUncachedEntryCriteria200Response instantiates a new AddUncachedEntryCriteria200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUncachedEntryCriteria200ResponseWithDefaults

`func NewAddUncachedEntryCriteria200ResponseWithDefaults() *AddUncachedEntryCriteria200Response`

NewAddUncachedEntryCriteria200ResponseWithDefaults instantiates a new AddUncachedEntryCriteria200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddUncachedEntryCriteria200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddUncachedEntryCriteria200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddUncachedEntryCriteria200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddUncachedEntryCriteria200Response) GetSchemas() []EnumthirdPartyUncachedEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUncachedEntryCriteria200Response) GetSchemasOk() (*[]EnumthirdPartyUncachedEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUncachedEntryCriteria200Response) SetSchemas(v []EnumthirdPartyUncachedEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddUncachedEntryCriteria200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUncachedEntryCriteria200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUncachedEntryCriteria200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUncachedEntryCriteria200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUncachedEntryCriteria200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUncachedEntryCriteria200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUncachedEntryCriteria200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAccessTimeThreshold

`func (o *AddUncachedEntryCriteria200Response) GetAccessTimeThreshold() string`

GetAccessTimeThreshold returns the AccessTimeThreshold field if non-nil, zero value otherwise.

### GetAccessTimeThresholdOk

`func (o *AddUncachedEntryCriteria200Response) GetAccessTimeThresholdOk() (*string, bool)`

GetAccessTimeThresholdOk returns a tuple with the AccessTimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTimeThreshold

`func (o *AddUncachedEntryCriteria200Response) SetAccessTimeThreshold(v string)`

SetAccessTimeThreshold sets AccessTimeThreshold field to given value.


### GetFilter

`func (o *AddUncachedEntryCriteria200Response) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddUncachedEntryCriteria200Response) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddUncachedEntryCriteria200Response) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetFilterIdentifiesUncachedEntries

`func (o *AddUncachedEntryCriteria200Response) GetFilterIdentifiesUncachedEntries() bool`

GetFilterIdentifiesUncachedEntries returns the FilterIdentifiesUncachedEntries field if non-nil, zero value otherwise.

### GetFilterIdentifiesUncachedEntriesOk

`func (o *AddUncachedEntryCriteria200Response) GetFilterIdentifiesUncachedEntriesOk() (*bool, bool)`

GetFilterIdentifiesUncachedEntriesOk returns a tuple with the FilterIdentifiesUncachedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIdentifiesUncachedEntries

`func (o *AddUncachedEntryCriteria200Response) SetFilterIdentifiesUncachedEntries(v bool)`

SetFilterIdentifiesUncachedEntries sets FilterIdentifiesUncachedEntries field to given value.


### GetScriptClass

`func (o *AddUncachedEntryCriteria200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddUncachedEntryCriteria200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddUncachedEntryCriteria200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddUncachedEntryCriteria200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddUncachedEntryCriteria200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddUncachedEntryCriteria200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddUncachedEntryCriteria200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddUncachedEntryCriteria200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddUncachedEntryCriteria200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddUncachedEntryCriteria200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddUncachedEntryCriteria200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddUncachedEntryCriteria200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddUncachedEntryCriteria200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddUncachedEntryCriteria200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


