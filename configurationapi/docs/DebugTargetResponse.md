# DebugTargetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Publisher | 
**Schemas** | Pointer to [**[]EnumdebugTargetSchemaUrn**](EnumdebugTargetSchemaUrn.md) |  | [optional] 
**DebugScope** | **string** | Specifies the fully-qualified Java package, class, or method affected by the settings in this target definition. Use the number character (#) to separate the class name and the method name (that is, com.unboundid.directory.server.core.DirectoryServer#startUp). | 
**DebugLevel** | [**EnumdebugTargetDebugLevelProp**](EnumdebugTargetDebugLevelProp.md) |  | 
**DebugCategory** | Pointer to [**[]EnumdebugTargetDebugCategoryProp**](EnumdebugTargetDebugCategoryProp.md) |  | [optional] 
**OmitMethodEntryArguments** | Pointer to **bool** | Specifies the property to indicate whether to include method arguments in debug messages. | [optional] 
**OmitMethodReturnValue** | Pointer to **bool** | Specifies the property to indicate whether to include the return value in debug messages. | [optional] 
**IncludeThrowableCause** | Pointer to **bool** | Specifies the property to indicate whether to include the cause of exceptions in exception thrown and caught messages. | [optional] 
**ThrowableStackFrames** | Pointer to **int64** | Specifies the property to indicate the number of stack frames to include in the stack trace for method entry and exception thrown messages. | [optional] 
**Description** | Pointer to **string** | A description for this Debug Target | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewDebugTargetResponse

`func NewDebugTargetResponse(id string, debugScope string, debugLevel EnumdebugTargetDebugLevelProp, ) *DebugTargetResponse`

NewDebugTargetResponse instantiates a new DebugTargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugTargetResponseWithDefaults

`func NewDebugTargetResponseWithDefaults() *DebugTargetResponse`

NewDebugTargetResponseWithDefaults instantiates a new DebugTargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DebugTargetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DebugTargetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DebugTargetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DebugTargetResponse) GetSchemas() []EnumdebugTargetSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DebugTargetResponse) GetSchemasOk() (*[]EnumdebugTargetSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DebugTargetResponse) SetSchemas(v []EnumdebugTargetSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DebugTargetResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDebugScope

`func (o *DebugTargetResponse) GetDebugScope() string`

GetDebugScope returns the DebugScope field if non-nil, zero value otherwise.

### GetDebugScopeOk

`func (o *DebugTargetResponse) GetDebugScopeOk() (*string, bool)`

GetDebugScopeOk returns a tuple with the DebugScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugScope

`func (o *DebugTargetResponse) SetDebugScope(v string)`

SetDebugScope sets DebugScope field to given value.


### GetDebugLevel

`func (o *DebugTargetResponse) GetDebugLevel() EnumdebugTargetDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *DebugTargetResponse) GetDebugLevelOk() (*EnumdebugTargetDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *DebugTargetResponse) SetDebugLevel(v EnumdebugTargetDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.


### GetDebugCategory

`func (o *DebugTargetResponse) GetDebugCategory() []EnumdebugTargetDebugCategoryProp`

GetDebugCategory returns the DebugCategory field if non-nil, zero value otherwise.

### GetDebugCategoryOk

`func (o *DebugTargetResponse) GetDebugCategoryOk() (*[]EnumdebugTargetDebugCategoryProp, bool)`

GetDebugCategoryOk returns a tuple with the DebugCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugCategory

`func (o *DebugTargetResponse) SetDebugCategory(v []EnumdebugTargetDebugCategoryProp)`

SetDebugCategory sets DebugCategory field to given value.

### HasDebugCategory

`func (o *DebugTargetResponse) HasDebugCategory() bool`

HasDebugCategory returns a boolean if a field has been set.

### GetOmitMethodEntryArguments

`func (o *DebugTargetResponse) GetOmitMethodEntryArguments() bool`

GetOmitMethodEntryArguments returns the OmitMethodEntryArguments field if non-nil, zero value otherwise.

### GetOmitMethodEntryArgumentsOk

`func (o *DebugTargetResponse) GetOmitMethodEntryArgumentsOk() (*bool, bool)`

GetOmitMethodEntryArgumentsOk returns a tuple with the OmitMethodEntryArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitMethodEntryArguments

`func (o *DebugTargetResponse) SetOmitMethodEntryArguments(v bool)`

SetOmitMethodEntryArguments sets OmitMethodEntryArguments field to given value.

### HasOmitMethodEntryArguments

`func (o *DebugTargetResponse) HasOmitMethodEntryArguments() bool`

HasOmitMethodEntryArguments returns a boolean if a field has been set.

### GetOmitMethodReturnValue

`func (o *DebugTargetResponse) GetOmitMethodReturnValue() bool`

GetOmitMethodReturnValue returns the OmitMethodReturnValue field if non-nil, zero value otherwise.

### GetOmitMethodReturnValueOk

`func (o *DebugTargetResponse) GetOmitMethodReturnValueOk() (*bool, bool)`

GetOmitMethodReturnValueOk returns a tuple with the OmitMethodReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitMethodReturnValue

`func (o *DebugTargetResponse) SetOmitMethodReturnValue(v bool)`

SetOmitMethodReturnValue sets OmitMethodReturnValue field to given value.

### HasOmitMethodReturnValue

`func (o *DebugTargetResponse) HasOmitMethodReturnValue() bool`

HasOmitMethodReturnValue returns a boolean if a field has been set.

### GetIncludeThrowableCause

`func (o *DebugTargetResponse) GetIncludeThrowableCause() bool`

GetIncludeThrowableCause returns the IncludeThrowableCause field if non-nil, zero value otherwise.

### GetIncludeThrowableCauseOk

`func (o *DebugTargetResponse) GetIncludeThrowableCauseOk() (*bool, bool)`

GetIncludeThrowableCauseOk returns a tuple with the IncludeThrowableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeThrowableCause

`func (o *DebugTargetResponse) SetIncludeThrowableCause(v bool)`

SetIncludeThrowableCause sets IncludeThrowableCause field to given value.

### HasIncludeThrowableCause

`func (o *DebugTargetResponse) HasIncludeThrowableCause() bool`

HasIncludeThrowableCause returns a boolean if a field has been set.

### GetThrowableStackFrames

`func (o *DebugTargetResponse) GetThrowableStackFrames() int64`

GetThrowableStackFrames returns the ThrowableStackFrames field if non-nil, zero value otherwise.

### GetThrowableStackFramesOk

`func (o *DebugTargetResponse) GetThrowableStackFramesOk() (*int64, bool)`

GetThrowableStackFramesOk returns a tuple with the ThrowableStackFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrowableStackFrames

`func (o *DebugTargetResponse) SetThrowableStackFrames(v int64)`

SetThrowableStackFrames sets ThrowableStackFrames field to given value.

### HasThrowableStackFrames

`func (o *DebugTargetResponse) HasThrowableStackFrames() bool`

HasThrowableStackFrames returns a boolean if a field has been set.

### GetDescription

`func (o *DebugTargetResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DebugTargetResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DebugTargetResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DebugTargetResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *DebugTargetResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DebugTargetResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DebugTargetResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DebugTargetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DebugTargetResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DebugTargetResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DebugTargetResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DebugTargetResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


