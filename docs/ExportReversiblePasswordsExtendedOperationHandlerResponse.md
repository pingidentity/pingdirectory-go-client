# ExportReversiblePasswordsExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Extended Operation Handler | 
**Schemas** | [**[]EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn**](EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewExportReversiblePasswordsExtendedOperationHandlerResponse

`func NewExportReversiblePasswordsExtendedOperationHandlerResponse(id string, schemas []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, enabled bool, ) *ExportReversiblePasswordsExtendedOperationHandlerResponse`

NewExportReversiblePasswordsExtendedOperationHandlerResponse instantiates a new ExportReversiblePasswordsExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportReversiblePasswordsExtendedOperationHandlerResponseWithDefaults

`func NewExportReversiblePasswordsExtendedOperationHandlerResponseWithDefaults() *ExportReversiblePasswordsExtendedOperationHandlerResponse`

NewExportReversiblePasswordsExtendedOperationHandlerResponseWithDefaults instantiates a new ExportReversiblePasswordsExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetSchemas() []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetSchemas(v []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ExportReversiblePasswordsExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


