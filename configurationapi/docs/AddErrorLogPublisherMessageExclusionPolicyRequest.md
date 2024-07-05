# AddErrorLogPublisherMessageExclusionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumerrorLogPublisherMessageExclusionPolicySchemaUrn**](EnumerrorLogPublisherMessageExclusionPolicySchemaUrn.md) |  | 
**LogMessageCategory** | [**EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp**](EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp.md) |  | 
**LogMessageSeverity** | [**EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp**](EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp.md) |  | 
**LogMessageID** | Pointer to **int64** | Indicates the message id the Error Log Publisher Message Exclusion Policy will check to exclude. | [optional] 
**LogMessageRegex** | **string** | The message regex the Error Log Publisher Message Exclusion Policy will check to exclude. | 
**Description** | Pointer to **string** | A description for this Log Publisher Message Exclusion Policy | [optional] 
**Enabled** | **bool** | Indicates whether the Log Publisher Message Exclusion Policy is enabled for use. | 
**PolicyName** | **string** | Name of the new Log Publisher Message Exclusion Policy | 

## Methods

### NewAddErrorLogPublisherMessageExclusionPolicyRequest

`func NewAddErrorLogPublisherMessageExclusionPolicyRequest(schemas []EnumerrorLogPublisherMessageExclusionPolicySchemaUrn, logMessageCategory EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp, logMessageSeverity EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp, logMessageRegex string, enabled bool, policyName string, ) *AddErrorLogPublisherMessageExclusionPolicyRequest`

NewAddErrorLogPublisherMessageExclusionPolicyRequest instantiates a new AddErrorLogPublisherMessageExclusionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddErrorLogPublisherMessageExclusionPolicyRequestWithDefaults

`func NewAddErrorLogPublisherMessageExclusionPolicyRequestWithDefaults() *AddErrorLogPublisherMessageExclusionPolicyRequest`

NewAddErrorLogPublisherMessageExclusionPolicyRequestWithDefaults instantiates a new AddErrorLogPublisherMessageExclusionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetSchemas() []EnumerrorLogPublisherMessageExclusionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetSchemasOk() (*[]EnumerrorLogPublisherMessageExclusionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetSchemas(v []EnumerrorLogPublisherMessageExclusionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogMessageCategory

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageCategory() EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp`

GetLogMessageCategory returns the LogMessageCategory field if non-nil, zero value otherwise.

### GetLogMessageCategoryOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageCategoryOk() (*EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp, bool)`

GetLogMessageCategoryOk returns a tuple with the LogMessageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageCategory

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetLogMessageCategory(v EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp)`

SetLogMessageCategory sets LogMessageCategory field to given value.


### GetLogMessageSeverity

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageSeverity() EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp`

GetLogMessageSeverity returns the LogMessageSeverity field if non-nil, zero value otherwise.

### GetLogMessageSeverityOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageSeverityOk() (*EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp, bool)`

GetLogMessageSeverityOk returns a tuple with the LogMessageSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageSeverity

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetLogMessageSeverity(v EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp)`

SetLogMessageSeverity sets LogMessageSeverity field to given value.


### GetLogMessageID

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageID() int64`

GetLogMessageID returns the LogMessageID field if non-nil, zero value otherwise.

### GetLogMessageIDOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageIDOk() (*int64, bool)`

GetLogMessageIDOk returns a tuple with the LogMessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageID

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetLogMessageID(v int64)`

SetLogMessageID sets LogMessageID field to given value.

### HasLogMessageID

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) HasLogMessageID() bool`

HasLogMessageID returns a boolean if a field has been set.

### GetLogMessageRegex

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageRegex() string`

GetLogMessageRegex returns the LogMessageRegex field if non-nil, zero value otherwise.

### GetLogMessageRegexOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetLogMessageRegexOk() (*string, bool)`

GetLogMessageRegexOk returns a tuple with the LogMessageRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageRegex

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetLogMessageRegex(v string)`

SetLogMessageRegex sets LogMessageRegex field to given value.


### GetDescription

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPolicyName

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddErrorLogPublisherMessageExclusionPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


