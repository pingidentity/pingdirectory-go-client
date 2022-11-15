# ErrorLogFieldMappingShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumerrorLogFieldMappingSchemaUrn**](EnumerrorLogFieldMappingSchemaUrn.md) |  | 
**LogFieldTimestamp** | Pointer to **string** | The time that the log message was generated. | [optional] 
**LogFieldProductName** | Pointer to **string** | The name for this Directory Server product, which may be used to identify which product was used to log the message if multiple products log to the same database table. | [optional] 
**LogFieldInstanceName** | Pointer to **string** | A name that uniquely identifies this Directory Server instance, which may be used to identify which instance was used to log the message if multiple server instances log to the same database table. | [optional] 
**LogFieldStartupid** | Pointer to **string** | The startup ID for the Directory Server. A different value will be generated each time the server is started. | [optional] 
**LogFieldCategory** | Pointer to **string** | The category for the log message. | [optional] 
**LogFieldSeverity** | Pointer to **string** | The severity for the log message. | [optional] 
**LogFieldMessageID** | Pointer to **string** | The numeric value which uniquely identifies the type of message. | [optional] 
**LogFieldMessage** | Pointer to **string** | The text of the log message. | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Mapping | [optional] 

## Methods

### NewErrorLogFieldMappingShared

`func NewErrorLogFieldMappingShared(schemas []EnumerrorLogFieldMappingSchemaUrn, ) *ErrorLogFieldMappingShared`

NewErrorLogFieldMappingShared instantiates a new ErrorLogFieldMappingShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogFieldMappingSharedWithDefaults

`func NewErrorLogFieldMappingSharedWithDefaults() *ErrorLogFieldMappingShared`

NewErrorLogFieldMappingSharedWithDefaults instantiates a new ErrorLogFieldMappingShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ErrorLogFieldMappingShared) GetSchemas() []EnumerrorLogFieldMappingSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ErrorLogFieldMappingShared) GetSchemasOk() (*[]EnumerrorLogFieldMappingSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ErrorLogFieldMappingShared) SetSchemas(v []EnumerrorLogFieldMappingSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogFieldTimestamp

`func (o *ErrorLogFieldMappingShared) GetLogFieldTimestamp() string`

GetLogFieldTimestamp returns the LogFieldTimestamp field if non-nil, zero value otherwise.

### GetLogFieldTimestampOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldTimestampOk() (*string, bool)`

GetLogFieldTimestampOk returns a tuple with the LogFieldTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTimestamp

`func (o *ErrorLogFieldMappingShared) SetLogFieldTimestamp(v string)`

SetLogFieldTimestamp sets LogFieldTimestamp field to given value.

### HasLogFieldTimestamp

`func (o *ErrorLogFieldMappingShared) HasLogFieldTimestamp() bool`

HasLogFieldTimestamp returns a boolean if a field has been set.

### GetLogFieldProductName

`func (o *ErrorLogFieldMappingShared) GetLogFieldProductName() string`

GetLogFieldProductName returns the LogFieldProductName field if non-nil, zero value otherwise.

### GetLogFieldProductNameOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldProductNameOk() (*string, bool)`

GetLogFieldProductNameOk returns a tuple with the LogFieldProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProductName

`func (o *ErrorLogFieldMappingShared) SetLogFieldProductName(v string)`

SetLogFieldProductName sets LogFieldProductName field to given value.

### HasLogFieldProductName

`func (o *ErrorLogFieldMappingShared) HasLogFieldProductName() bool`

HasLogFieldProductName returns a boolean if a field has been set.

### GetLogFieldInstanceName

`func (o *ErrorLogFieldMappingShared) GetLogFieldInstanceName() string`

GetLogFieldInstanceName returns the LogFieldInstanceName field if non-nil, zero value otherwise.

### GetLogFieldInstanceNameOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldInstanceNameOk() (*string, bool)`

GetLogFieldInstanceNameOk returns a tuple with the LogFieldInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldInstanceName

`func (o *ErrorLogFieldMappingShared) SetLogFieldInstanceName(v string)`

SetLogFieldInstanceName sets LogFieldInstanceName field to given value.

### HasLogFieldInstanceName

`func (o *ErrorLogFieldMappingShared) HasLogFieldInstanceName() bool`

HasLogFieldInstanceName returns a boolean if a field has been set.

### GetLogFieldStartupid

`func (o *ErrorLogFieldMappingShared) GetLogFieldStartupid() string`

GetLogFieldStartupid returns the LogFieldStartupid field if non-nil, zero value otherwise.

### GetLogFieldStartupidOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldStartupidOk() (*string, bool)`

GetLogFieldStartupidOk returns a tuple with the LogFieldStartupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldStartupid

`func (o *ErrorLogFieldMappingShared) SetLogFieldStartupid(v string)`

SetLogFieldStartupid sets LogFieldStartupid field to given value.

### HasLogFieldStartupid

`func (o *ErrorLogFieldMappingShared) HasLogFieldStartupid() bool`

HasLogFieldStartupid returns a boolean if a field has been set.

### GetLogFieldCategory

`func (o *ErrorLogFieldMappingShared) GetLogFieldCategory() string`

GetLogFieldCategory returns the LogFieldCategory field if non-nil, zero value otherwise.

### GetLogFieldCategoryOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldCategoryOk() (*string, bool)`

GetLogFieldCategoryOk returns a tuple with the LogFieldCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldCategory

`func (o *ErrorLogFieldMappingShared) SetLogFieldCategory(v string)`

SetLogFieldCategory sets LogFieldCategory field to given value.

### HasLogFieldCategory

`func (o *ErrorLogFieldMappingShared) HasLogFieldCategory() bool`

HasLogFieldCategory returns a boolean if a field has been set.

### GetLogFieldSeverity

`func (o *ErrorLogFieldMappingShared) GetLogFieldSeverity() string`

GetLogFieldSeverity returns the LogFieldSeverity field if non-nil, zero value otherwise.

### GetLogFieldSeverityOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldSeverityOk() (*string, bool)`

GetLogFieldSeverityOk returns a tuple with the LogFieldSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSeverity

`func (o *ErrorLogFieldMappingShared) SetLogFieldSeverity(v string)`

SetLogFieldSeverity sets LogFieldSeverity field to given value.

### HasLogFieldSeverity

`func (o *ErrorLogFieldMappingShared) HasLogFieldSeverity() bool`

HasLogFieldSeverity returns a boolean if a field has been set.

### GetLogFieldMessageID

`func (o *ErrorLogFieldMappingShared) GetLogFieldMessageID() string`

GetLogFieldMessageID returns the LogFieldMessageID field if non-nil, zero value otherwise.

### GetLogFieldMessageIDOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldMessageIDOk() (*string, bool)`

GetLogFieldMessageIDOk returns a tuple with the LogFieldMessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageID

`func (o *ErrorLogFieldMappingShared) SetLogFieldMessageID(v string)`

SetLogFieldMessageID sets LogFieldMessageID field to given value.

### HasLogFieldMessageID

`func (o *ErrorLogFieldMappingShared) HasLogFieldMessageID() bool`

HasLogFieldMessageID returns a boolean if a field has been set.

### GetLogFieldMessage

`func (o *ErrorLogFieldMappingShared) GetLogFieldMessage() string`

GetLogFieldMessage returns the LogFieldMessage field if non-nil, zero value otherwise.

### GetLogFieldMessageOk

`func (o *ErrorLogFieldMappingShared) GetLogFieldMessageOk() (*string, bool)`

GetLogFieldMessageOk returns a tuple with the LogFieldMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessage

`func (o *ErrorLogFieldMappingShared) SetLogFieldMessage(v string)`

SetLogFieldMessage sets LogFieldMessage field to given value.

### HasLogFieldMessage

`func (o *ErrorLogFieldMappingShared) HasLogFieldMessage() bool`

HasLogFieldMessage returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorLogFieldMappingShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorLogFieldMappingShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorLogFieldMappingShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorLogFieldMappingShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


