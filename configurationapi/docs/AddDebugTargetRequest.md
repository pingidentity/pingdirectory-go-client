# AddDebugTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetName** | **string** | Name of the new Debug Target | 
**Schemas** | Pointer to [**[]EnumdebugTargetSchemaUrn**](EnumdebugTargetSchemaUrn.md) |  | [optional] 
**DebugScope** | **string** | Specifies the fully-qualified Java package, class, or method affected by the settings in this target definition. Use the number character (#) to separate the class name and the method name (that is, com.unboundid.directory.server.core.DirectoryServer#startUp). | 
**DebugLevel** | [**EnumdebugTargetDebugLevelProp**](EnumdebugTargetDebugLevelProp.md) |  | 
**DebugCategory** | Pointer to [**[]EnumdebugTargetDebugCategoryProp**](EnumdebugTargetDebugCategoryProp.md) |  | [optional] 
**OmitMethodEntryArguments** | Pointer to **bool** | Specifies the property to indicate whether to include method arguments in debug messages. | [optional] 
**OmitMethodReturnValue** | Pointer to **bool** | Specifies the property to indicate whether to include the return value in debug messages. | [optional] 
**IncludeThrowableCause** | Pointer to **bool** | Specifies the property to indicate whether to include the cause of exceptions in exception thrown and caught messages. | [optional] 
**ThrowableStackFrames** | Pointer to **int64** | Specifies the property to indicate the number of stack frames to include in the stack trace for method entry and exception thrown messages. | [optional] 
**Description** | Pointer to **string** | A description for this Debug Target | [optional] 

## Methods

### NewAddDebugTargetRequest

`func NewAddDebugTargetRequest(targetName string, debugScope string, debugLevel EnumdebugTargetDebugLevelProp, ) *AddDebugTargetRequest`

NewAddDebugTargetRequest instantiates a new AddDebugTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDebugTargetRequestWithDefaults

`func NewAddDebugTargetRequestWithDefaults() *AddDebugTargetRequest`

NewAddDebugTargetRequestWithDefaults instantiates a new AddDebugTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetName

`func (o *AddDebugTargetRequest) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *AddDebugTargetRequest) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *AddDebugTargetRequest) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetSchemas

`func (o *AddDebugTargetRequest) GetSchemas() []EnumdebugTargetSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDebugTargetRequest) GetSchemasOk() (*[]EnumdebugTargetSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDebugTargetRequest) SetSchemas(v []EnumdebugTargetSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddDebugTargetRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDebugScope

`func (o *AddDebugTargetRequest) GetDebugScope() string`

GetDebugScope returns the DebugScope field if non-nil, zero value otherwise.

### GetDebugScopeOk

`func (o *AddDebugTargetRequest) GetDebugScopeOk() (*string, bool)`

GetDebugScopeOk returns a tuple with the DebugScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugScope

`func (o *AddDebugTargetRequest) SetDebugScope(v string)`

SetDebugScope sets DebugScope field to given value.


### GetDebugLevel

`func (o *AddDebugTargetRequest) GetDebugLevel() EnumdebugTargetDebugLevelProp`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *AddDebugTargetRequest) GetDebugLevelOk() (*EnumdebugTargetDebugLevelProp, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *AddDebugTargetRequest) SetDebugLevel(v EnumdebugTargetDebugLevelProp)`

SetDebugLevel sets DebugLevel field to given value.


### GetDebugCategory

`func (o *AddDebugTargetRequest) GetDebugCategory() []EnumdebugTargetDebugCategoryProp`

GetDebugCategory returns the DebugCategory field if non-nil, zero value otherwise.

### GetDebugCategoryOk

`func (o *AddDebugTargetRequest) GetDebugCategoryOk() (*[]EnumdebugTargetDebugCategoryProp, bool)`

GetDebugCategoryOk returns a tuple with the DebugCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugCategory

`func (o *AddDebugTargetRequest) SetDebugCategory(v []EnumdebugTargetDebugCategoryProp)`

SetDebugCategory sets DebugCategory field to given value.

### HasDebugCategory

`func (o *AddDebugTargetRequest) HasDebugCategory() bool`

HasDebugCategory returns a boolean if a field has been set.

### GetOmitMethodEntryArguments

`func (o *AddDebugTargetRequest) GetOmitMethodEntryArguments() bool`

GetOmitMethodEntryArguments returns the OmitMethodEntryArguments field if non-nil, zero value otherwise.

### GetOmitMethodEntryArgumentsOk

`func (o *AddDebugTargetRequest) GetOmitMethodEntryArgumentsOk() (*bool, bool)`

GetOmitMethodEntryArgumentsOk returns a tuple with the OmitMethodEntryArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitMethodEntryArguments

`func (o *AddDebugTargetRequest) SetOmitMethodEntryArguments(v bool)`

SetOmitMethodEntryArguments sets OmitMethodEntryArguments field to given value.

### HasOmitMethodEntryArguments

`func (o *AddDebugTargetRequest) HasOmitMethodEntryArguments() bool`

HasOmitMethodEntryArguments returns a boolean if a field has been set.

### GetOmitMethodReturnValue

`func (o *AddDebugTargetRequest) GetOmitMethodReturnValue() bool`

GetOmitMethodReturnValue returns the OmitMethodReturnValue field if non-nil, zero value otherwise.

### GetOmitMethodReturnValueOk

`func (o *AddDebugTargetRequest) GetOmitMethodReturnValueOk() (*bool, bool)`

GetOmitMethodReturnValueOk returns a tuple with the OmitMethodReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitMethodReturnValue

`func (o *AddDebugTargetRequest) SetOmitMethodReturnValue(v bool)`

SetOmitMethodReturnValue sets OmitMethodReturnValue field to given value.

### HasOmitMethodReturnValue

`func (o *AddDebugTargetRequest) HasOmitMethodReturnValue() bool`

HasOmitMethodReturnValue returns a boolean if a field has been set.

### GetIncludeThrowableCause

`func (o *AddDebugTargetRequest) GetIncludeThrowableCause() bool`

GetIncludeThrowableCause returns the IncludeThrowableCause field if non-nil, zero value otherwise.

### GetIncludeThrowableCauseOk

`func (o *AddDebugTargetRequest) GetIncludeThrowableCauseOk() (*bool, bool)`

GetIncludeThrowableCauseOk returns a tuple with the IncludeThrowableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeThrowableCause

`func (o *AddDebugTargetRequest) SetIncludeThrowableCause(v bool)`

SetIncludeThrowableCause sets IncludeThrowableCause field to given value.

### HasIncludeThrowableCause

`func (o *AddDebugTargetRequest) HasIncludeThrowableCause() bool`

HasIncludeThrowableCause returns a boolean if a field has been set.

### GetThrowableStackFrames

`func (o *AddDebugTargetRequest) GetThrowableStackFrames() int64`

GetThrowableStackFrames returns the ThrowableStackFrames field if non-nil, zero value otherwise.

### GetThrowableStackFramesOk

`func (o *AddDebugTargetRequest) GetThrowableStackFramesOk() (*int64, bool)`

GetThrowableStackFramesOk returns a tuple with the ThrowableStackFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrowableStackFrames

`func (o *AddDebugTargetRequest) SetThrowableStackFrames(v int64)`

SetThrowableStackFrames sets ThrowableStackFrames field to given value.

### HasThrowableStackFrames

`func (o *AddDebugTargetRequest) HasThrowableStackFrames() bool`

HasThrowableStackFrames returns a boolean if a field has been set.

### GetDescription

`func (o *AddDebugTargetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDebugTargetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDebugTargetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDebugTargetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


