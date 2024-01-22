/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddVelocityTemplateLoaderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVelocityTemplateLoaderRequest{}

// AddVelocityTemplateLoaderRequest struct for AddVelocityTemplateLoaderRequest
type AddVelocityTemplateLoaderRequest struct {
	Schemas []EnumvelocityTemplateLoaderSchemaUrn `json:"schemas,omitempty"`
	// Indicates whether this Velocity Template Loader is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// This property determines the evaluation order for determining the correct Velocity Template Loader to load a template for generating content for a particular request.
	EvaluationOrderIndex *int64 `json:"evaluationOrderIndex,omitempty"`
	// Specifies a media type for matching Accept request-header values.
	MimeTypeMatcher string `json:"mimeTypeMatcher"`
	// Specifies a the value that will be used in the response's Content-Type header that indicates the type of content to return.
	MimeType *string `json:"mimeType,omitempty"`
	// Specifies the suffix to append to the requested resource name when searching for the template file with which to form a response.
	TemplateSuffix *string `json:"templateSuffix,omitempty"`
	// Specifies the directory in which to search for the template files.
	TemplateDirectory *string `json:"templateDirectory,omitempty"`
	// Name of the new Velocity Template Loader
	LoaderName string `json:"loaderName"`
}

// NewAddVelocityTemplateLoaderRequest instantiates a new AddVelocityTemplateLoaderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVelocityTemplateLoaderRequest(mimeTypeMatcher string, loaderName string) *AddVelocityTemplateLoaderRequest {
	this := AddVelocityTemplateLoaderRequest{}
	this.MimeTypeMatcher = mimeTypeMatcher
	this.LoaderName = loaderName
	return &this
}

// NewAddVelocityTemplateLoaderRequestWithDefaults instantiates a new AddVelocityTemplateLoaderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVelocityTemplateLoaderRequestWithDefaults() *AddVelocityTemplateLoaderRequest {
	this := AddVelocityTemplateLoaderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddVelocityTemplateLoaderRequest) GetSchemas() []EnumvelocityTemplateLoaderSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumvelocityTemplateLoaderSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetSchemasOk() ([]EnumvelocityTemplateLoaderSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddVelocityTemplateLoaderRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumvelocityTemplateLoaderSchemaUrn and assigns it to the Schemas field.
func (o *AddVelocityTemplateLoaderRequest) SetSchemas(v []EnumvelocityTemplateLoaderSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddVelocityTemplateLoaderRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddVelocityTemplateLoaderRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddVelocityTemplateLoaderRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value if set, zero value otherwise.
func (o *AddVelocityTemplateLoaderRequest) GetEvaluationOrderIndex() int64 {
	if o == nil || IsNil(o.EvaluationOrderIndex) {
		var ret int64
		return ret
	}
	return *o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.EvaluationOrderIndex) {
		return nil, false
	}
	return o.EvaluationOrderIndex, true
}

// HasEvaluationOrderIndex returns a boolean if a field has been set.
func (o *AddVelocityTemplateLoaderRequest) HasEvaluationOrderIndex() bool {
	if o != nil && !IsNil(o.EvaluationOrderIndex) {
		return true
	}

	return false
}

// SetEvaluationOrderIndex gets a reference to the given int64 and assigns it to the EvaluationOrderIndex field.
func (o *AddVelocityTemplateLoaderRequest) SetEvaluationOrderIndex(v int64) {
	o.EvaluationOrderIndex = &v
}

// GetMimeTypeMatcher returns the MimeTypeMatcher field value
func (o *AddVelocityTemplateLoaderRequest) GetMimeTypeMatcher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeTypeMatcher
}

// GetMimeTypeMatcherOk returns a tuple with the MimeTypeMatcher field value
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetMimeTypeMatcherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeTypeMatcher, true
}

// SetMimeTypeMatcher sets field value
func (o *AddVelocityTemplateLoaderRequest) SetMimeTypeMatcher(v string) {
	o.MimeTypeMatcher = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *AddVelocityTemplateLoaderRequest) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *AddVelocityTemplateLoaderRequest) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *AddVelocityTemplateLoaderRequest) SetMimeType(v string) {
	o.MimeType = &v
}

// GetTemplateSuffix returns the TemplateSuffix field value if set, zero value otherwise.
func (o *AddVelocityTemplateLoaderRequest) GetTemplateSuffix() string {
	if o == nil || IsNil(o.TemplateSuffix) {
		var ret string
		return ret
	}
	return *o.TemplateSuffix
}

// GetTemplateSuffixOk returns a tuple with the TemplateSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetTemplateSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateSuffix) {
		return nil, false
	}
	return o.TemplateSuffix, true
}

// HasTemplateSuffix returns a boolean if a field has been set.
func (o *AddVelocityTemplateLoaderRequest) HasTemplateSuffix() bool {
	if o != nil && !IsNil(o.TemplateSuffix) {
		return true
	}

	return false
}

// SetTemplateSuffix gets a reference to the given string and assigns it to the TemplateSuffix field.
func (o *AddVelocityTemplateLoaderRequest) SetTemplateSuffix(v string) {
	o.TemplateSuffix = &v
}

// GetTemplateDirectory returns the TemplateDirectory field value if set, zero value otherwise.
func (o *AddVelocityTemplateLoaderRequest) GetTemplateDirectory() string {
	if o == nil || IsNil(o.TemplateDirectory) {
		var ret string
		return ret
	}
	return *o.TemplateDirectory
}

// GetTemplateDirectoryOk returns a tuple with the TemplateDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetTemplateDirectoryOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateDirectory) {
		return nil, false
	}
	return o.TemplateDirectory, true
}

// HasTemplateDirectory returns a boolean if a field has been set.
func (o *AddVelocityTemplateLoaderRequest) HasTemplateDirectory() bool {
	if o != nil && !IsNil(o.TemplateDirectory) {
		return true
	}

	return false
}

// SetTemplateDirectory gets a reference to the given string and assigns it to the TemplateDirectory field.
func (o *AddVelocityTemplateLoaderRequest) SetTemplateDirectory(v string) {
	o.TemplateDirectory = &v
}

// GetLoaderName returns the LoaderName field value
func (o *AddVelocityTemplateLoaderRequest) GetLoaderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoaderName
}

// GetLoaderNameOk returns a tuple with the LoaderName field value
// and a boolean to check if the value has been set.
func (o *AddVelocityTemplateLoaderRequest) GetLoaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoaderName, true
}

// SetLoaderName sets field value
func (o *AddVelocityTemplateLoaderRequest) SetLoaderName(v string) {
	o.LoaderName = v
}

func (o AddVelocityTemplateLoaderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVelocityTemplateLoaderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EvaluationOrderIndex) {
		toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	}
	toSerialize["mimeTypeMatcher"] = o.MimeTypeMatcher
	if !IsNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
	}
	if !IsNil(o.TemplateSuffix) {
		toSerialize["templateSuffix"] = o.TemplateSuffix
	}
	if !IsNil(o.TemplateDirectory) {
		toSerialize["templateDirectory"] = o.TemplateDirectory
	}
	toSerialize["loaderName"] = o.LoaderName
	return toSerialize, nil
}

type NullableAddVelocityTemplateLoaderRequest struct {
	value *AddVelocityTemplateLoaderRequest
	isSet bool
}

func (v NullableAddVelocityTemplateLoaderRequest) Get() *AddVelocityTemplateLoaderRequest {
	return v.value
}

func (v *NullableAddVelocityTemplateLoaderRequest) Set(val *AddVelocityTemplateLoaderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVelocityTemplateLoaderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVelocityTemplateLoaderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVelocityTemplateLoaderRequest(val *AddVelocityTemplateLoaderRequest) *NullableAddVelocityTemplateLoaderRequest {
	return &NullableAddVelocityTemplateLoaderRequest{value: val, isSet: true}
}

func (v NullableAddVelocityTemplateLoaderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVelocityTemplateLoaderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
