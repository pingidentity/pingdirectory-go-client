# SubtreeAccessibilityExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn**](EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Extended Operation Handler | 
**NonAccessibleSubtreeUnauthorizedBindResultCode** | Pointer to **int64** | The integer value for the result code that the server should return to clients that attempt to perform an unauthorized bind operation in a non-accessible subtree. | [optional] 
**NonAccessibleSubtreeUnauthorizedReadResultCode** | Pointer to **int64** | The integer value for the result code that the server should return to clients that attempt to perform an unauthorized read (e.g., search or compare) operation in a non-accessible subtree. | [optional] 
**NonAccessibleSubtreeUnauthorizedWriteResultCode** | Pointer to **int64** | The integer value for the result code that the server should return to clients that attempt to perform an unauthorized write (e.g., add, delete, modify or modify DN) operation in a non-accessible subtree. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewSubtreeAccessibilityExtendedOperationHandlerResponse

`func NewSubtreeAccessibilityExtendedOperationHandlerResponse(schemas []EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn, id string, enabled bool, ) *SubtreeAccessibilityExtendedOperationHandlerResponse`

NewSubtreeAccessibilityExtendedOperationHandlerResponse instantiates a new SubtreeAccessibilityExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubtreeAccessibilityExtendedOperationHandlerResponseWithDefaults

`func NewSubtreeAccessibilityExtendedOperationHandlerResponseWithDefaults() *SubtreeAccessibilityExtendedOperationHandlerResponse`

NewSubtreeAccessibilityExtendedOperationHandlerResponseWithDefaults instantiates a new SubtreeAccessibilityExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetSchemas() []EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetSchemas(v []EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetNonAccessibleSubtreeUnauthorizedBindResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetNonAccessibleSubtreeUnauthorizedBindResultCode() int64`

GetNonAccessibleSubtreeUnauthorizedBindResultCode returns the NonAccessibleSubtreeUnauthorizedBindResultCode field if non-nil, zero value otherwise.

### GetNonAccessibleSubtreeUnauthorizedBindResultCodeOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetNonAccessibleSubtreeUnauthorizedBindResultCodeOk() (*int64, bool)`

GetNonAccessibleSubtreeUnauthorizedBindResultCodeOk returns a tuple with the NonAccessibleSubtreeUnauthorizedBindResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonAccessibleSubtreeUnauthorizedBindResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetNonAccessibleSubtreeUnauthorizedBindResultCode(v int64)`

SetNonAccessibleSubtreeUnauthorizedBindResultCode sets NonAccessibleSubtreeUnauthorizedBindResultCode field to given value.

### HasNonAccessibleSubtreeUnauthorizedBindResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) HasNonAccessibleSubtreeUnauthorizedBindResultCode() bool`

HasNonAccessibleSubtreeUnauthorizedBindResultCode returns a boolean if a field has been set.

### GetNonAccessibleSubtreeUnauthorizedReadResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetNonAccessibleSubtreeUnauthorizedReadResultCode() int64`

GetNonAccessibleSubtreeUnauthorizedReadResultCode returns the NonAccessibleSubtreeUnauthorizedReadResultCode field if non-nil, zero value otherwise.

### GetNonAccessibleSubtreeUnauthorizedReadResultCodeOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetNonAccessibleSubtreeUnauthorizedReadResultCodeOk() (*int64, bool)`

GetNonAccessibleSubtreeUnauthorizedReadResultCodeOk returns a tuple with the NonAccessibleSubtreeUnauthorizedReadResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonAccessibleSubtreeUnauthorizedReadResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetNonAccessibleSubtreeUnauthorizedReadResultCode(v int64)`

SetNonAccessibleSubtreeUnauthorizedReadResultCode sets NonAccessibleSubtreeUnauthorizedReadResultCode field to given value.

### HasNonAccessibleSubtreeUnauthorizedReadResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) HasNonAccessibleSubtreeUnauthorizedReadResultCode() bool`

HasNonAccessibleSubtreeUnauthorizedReadResultCode returns a boolean if a field has been set.

### GetNonAccessibleSubtreeUnauthorizedWriteResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetNonAccessibleSubtreeUnauthorizedWriteResultCode() int64`

GetNonAccessibleSubtreeUnauthorizedWriteResultCode returns the NonAccessibleSubtreeUnauthorizedWriteResultCode field if non-nil, zero value otherwise.

### GetNonAccessibleSubtreeUnauthorizedWriteResultCodeOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetNonAccessibleSubtreeUnauthorizedWriteResultCodeOk() (*int64, bool)`

GetNonAccessibleSubtreeUnauthorizedWriteResultCodeOk returns a tuple with the NonAccessibleSubtreeUnauthorizedWriteResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonAccessibleSubtreeUnauthorizedWriteResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetNonAccessibleSubtreeUnauthorizedWriteResultCode(v int64)`

SetNonAccessibleSubtreeUnauthorizedWriteResultCode sets NonAccessibleSubtreeUnauthorizedWriteResultCode field to given value.

### HasNonAccessibleSubtreeUnauthorizedWriteResultCode

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) HasNonAccessibleSubtreeUnauthorizedWriteResultCode() bool`

HasNonAccessibleSubtreeUnauthorizedWriteResultCode returns a boolean if a field has been set.

### GetDescription

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SubtreeAccessibilityExtendedOperationHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


