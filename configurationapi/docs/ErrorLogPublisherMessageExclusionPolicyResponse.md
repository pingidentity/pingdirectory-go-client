# ErrorLogPublisherMessageExclusionPolicyResponse

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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Log Publisher Message Exclusion Policy | 

## Methods

### NewErrorLogPublisherMessageExclusionPolicyResponse

`func NewErrorLogPublisherMessageExclusionPolicyResponse(schemas []EnumerrorLogPublisherMessageExclusionPolicySchemaUrn, logMessageCategory EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp, logMessageSeverity EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp, logMessageRegex string, enabled bool, id string, ) *ErrorLogPublisherMessageExclusionPolicyResponse`

NewErrorLogPublisherMessageExclusionPolicyResponse instantiates a new ErrorLogPublisherMessageExclusionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogPublisherMessageExclusionPolicyResponseWithDefaults

`func NewErrorLogPublisherMessageExclusionPolicyResponseWithDefaults() *ErrorLogPublisherMessageExclusionPolicyResponse`

NewErrorLogPublisherMessageExclusionPolicyResponseWithDefaults instantiates a new ErrorLogPublisherMessageExclusionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetSchemas() []EnumerrorLogPublisherMessageExclusionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetSchemasOk() (*[]EnumerrorLogPublisherMessageExclusionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetSchemas(v []EnumerrorLogPublisherMessageExclusionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogMessageCategory

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageCategory() EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp`

GetLogMessageCategory returns the LogMessageCategory field if non-nil, zero value otherwise.

### GetLogMessageCategoryOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageCategoryOk() (*EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp, bool)`

GetLogMessageCategoryOk returns a tuple with the LogMessageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageCategory

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetLogMessageCategory(v EnumlogPublisherMessageExclusionPolicyLogMessageCategoryProp)`

SetLogMessageCategory sets LogMessageCategory field to given value.


### GetLogMessageSeverity

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageSeverity() EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp`

GetLogMessageSeverity returns the LogMessageSeverity field if non-nil, zero value otherwise.

### GetLogMessageSeverityOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageSeverityOk() (*EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp, bool)`

GetLogMessageSeverityOk returns a tuple with the LogMessageSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageSeverity

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetLogMessageSeverity(v EnumlogPublisherMessageExclusionPolicyLogMessageSeverityProp)`

SetLogMessageSeverity sets LogMessageSeverity field to given value.


### GetLogMessageID

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageID() int64`

GetLogMessageID returns the LogMessageID field if non-nil, zero value otherwise.

### GetLogMessageIDOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageIDOk() (*int64, bool)`

GetLogMessageIDOk returns a tuple with the LogMessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageID

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetLogMessageID(v int64)`

SetLogMessageID sets LogMessageID field to given value.

### HasLogMessageID

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) HasLogMessageID() bool`

HasLogMessageID returns a boolean if a field has been set.

### GetLogMessageRegex

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageRegex() string`

GetLogMessageRegex returns the LogMessageRegex field if non-nil, zero value otherwise.

### GetLogMessageRegexOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetLogMessageRegexOk() (*string, bool)`

GetLogMessageRegexOk returns a tuple with the LogMessageRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessageRegex

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetLogMessageRegex(v string)`

SetLogMessageRegex sets LogMessageRegex field to given value.


### GetDescription

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorLogPublisherMessageExclusionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


