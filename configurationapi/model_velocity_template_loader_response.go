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

// checks if the VelocityTemplateLoaderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VelocityTemplateLoaderResponse{}

// VelocityTemplateLoaderResponse struct for VelocityTemplateLoaderResponse
type VelocityTemplateLoaderResponse struct {
	// Name of the HTTP Servlet Extension
	Id      string                                `json:"id"`
	Schemas []EnumvelocityTemplateLoaderSchemaUrn `json:"schemas,omitempty"`
	// Indicates whether this Velocity Template Loader is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// This property determines the evaluation order for determining the correct Velocity Template Loader to load a template for generating content for a particular request.
	EvaluationOrderIndex int32 `json:"evaluationOrderIndex"`
	// Specifies a media type for matching Accept request-header values.
	MimeTypeMatcher string `json:"mimeTypeMatcher"`
	// Specifies a the value that will be used in the response's Content-Type header that indicates the type of content to return.
	MimeType *string `json:"mimeType,omitempty"`
	// Specifies the suffix to append to the requested resource name when searching for the template file with which to form a response.
	TemplateSuffix *string `json:"templateSuffix,omitempty"`
	// Specifies the directory in which to search for the template files.
	TemplateDirectory                             *string                                            `json:"templateDirectory,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewVelocityTemplateLoaderResponse instantiates a new VelocityTemplateLoaderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVelocityTemplateLoaderResponse(id string, evaluationOrderIndex int32, mimeTypeMatcher string) *VelocityTemplateLoaderResponse {
	this := VelocityTemplateLoaderResponse{}
	this.Id = id
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.MimeTypeMatcher = mimeTypeMatcher
	return &this
}

// NewVelocityTemplateLoaderResponseWithDefaults instantiates a new VelocityTemplateLoaderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelocityTemplateLoaderResponseWithDefaults() *VelocityTemplateLoaderResponse {
	this := VelocityTemplateLoaderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *VelocityTemplateLoaderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VelocityTemplateLoaderResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *VelocityTemplateLoaderResponse) GetSchemas() []EnumvelocityTemplateLoaderSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumvelocityTemplateLoaderSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetSchemasOk() ([]EnumvelocityTemplateLoaderSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *VelocityTemplateLoaderResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumvelocityTemplateLoaderSchemaUrn and assigns it to the Schemas field.
func (o *VelocityTemplateLoaderResponse) SetSchemas(v []EnumvelocityTemplateLoaderSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VelocityTemplateLoaderResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VelocityTemplateLoaderResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VelocityTemplateLoaderResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *VelocityTemplateLoaderResponse) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *VelocityTemplateLoaderResponse) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

// GetMimeTypeMatcher returns the MimeTypeMatcher field value
func (o *VelocityTemplateLoaderResponse) GetMimeTypeMatcher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeTypeMatcher
}

// GetMimeTypeMatcherOk returns a tuple with the MimeTypeMatcher field value
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetMimeTypeMatcherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeTypeMatcher, true
}

// SetMimeTypeMatcher sets field value
func (o *VelocityTemplateLoaderResponse) SetMimeTypeMatcher(v string) {
	o.MimeTypeMatcher = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *VelocityTemplateLoaderResponse) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *VelocityTemplateLoaderResponse) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *VelocityTemplateLoaderResponse) SetMimeType(v string) {
	o.MimeType = &v
}

// GetTemplateSuffix returns the TemplateSuffix field value if set, zero value otherwise.
func (o *VelocityTemplateLoaderResponse) GetTemplateSuffix() string {
	if o == nil || IsNil(o.TemplateSuffix) {
		var ret string
		return ret
	}
	return *o.TemplateSuffix
}

// GetTemplateSuffixOk returns a tuple with the TemplateSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetTemplateSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateSuffix) {
		return nil, false
	}
	return o.TemplateSuffix, true
}

// HasTemplateSuffix returns a boolean if a field has been set.
func (o *VelocityTemplateLoaderResponse) HasTemplateSuffix() bool {
	if o != nil && !IsNil(o.TemplateSuffix) {
		return true
	}

	return false
}

// SetTemplateSuffix gets a reference to the given string and assigns it to the TemplateSuffix field.
func (o *VelocityTemplateLoaderResponse) SetTemplateSuffix(v string) {
	o.TemplateSuffix = &v
}

// GetTemplateDirectory returns the TemplateDirectory field value if set, zero value otherwise.
func (o *VelocityTemplateLoaderResponse) GetTemplateDirectory() string {
	if o == nil || IsNil(o.TemplateDirectory) {
		var ret string
		return ret
	}
	return *o.TemplateDirectory
}

// GetTemplateDirectoryOk returns a tuple with the TemplateDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetTemplateDirectoryOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateDirectory) {
		return nil, false
	}
	return o.TemplateDirectory, true
}

// HasTemplateDirectory returns a boolean if a field has been set.
func (o *VelocityTemplateLoaderResponse) HasTemplateDirectory() bool {
	if o != nil && !IsNil(o.TemplateDirectory) {
		return true
	}

	return false
}

// SetTemplateDirectory gets a reference to the given string and assigns it to the TemplateDirectory field.
func (o *VelocityTemplateLoaderResponse) SetTemplateDirectory(v string) {
	o.TemplateDirectory = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *VelocityTemplateLoaderResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *VelocityTemplateLoaderResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *VelocityTemplateLoaderResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *VelocityTemplateLoaderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityTemplateLoaderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *VelocityTemplateLoaderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *VelocityTemplateLoaderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o VelocityTemplateLoaderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VelocityTemplateLoaderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableVelocityTemplateLoaderResponse struct {
	value *VelocityTemplateLoaderResponse
	isSet bool
}

func (v NullableVelocityTemplateLoaderResponse) Get() *VelocityTemplateLoaderResponse {
	return v.value
}

func (v *NullableVelocityTemplateLoaderResponse) Set(val *VelocityTemplateLoaderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityTemplateLoaderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityTemplateLoaderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityTemplateLoaderResponse(val *VelocityTemplateLoaderResponse) *NullableVelocityTemplateLoaderResponse {
	return &NullableVelocityTemplateLoaderResponse{value: val, isSet: true}
}

func (v NullableVelocityTemplateLoaderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityTemplateLoaderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
