/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// CipherStreamProviderListResponseResourcesInner - struct for CipherStreamProviderListResponseResourcesInner
type CipherStreamProviderListResponseResourcesInner struct {
	AmazonKeyManagementServiceCipherStreamProviderResponse *AmazonKeyManagementServiceCipherStreamProviderResponse
	AmazonSecretsManagerCipherStreamProviderResponse       *AmazonSecretsManagerCipherStreamProviderResponse
	AzureKeyVaultCipherStreamProviderResponse              *AzureKeyVaultCipherStreamProviderResponse
	ConjurCipherStreamProviderResponse                     *ConjurCipherStreamProviderResponse
	DefaultCipherStreamProviderResponse                    *DefaultCipherStreamProviderResponse
	FileBasedCipherStreamProviderResponse                  *FileBasedCipherStreamProviderResponse
	Pkcs11CipherStreamProviderResponse                     *Pkcs11CipherStreamProviderResponse
	ThirdPartyCipherStreamProviderResponse                 *ThirdPartyCipherStreamProviderResponse
	VaultCipherStreamProviderResponse                      *VaultCipherStreamProviderResponse
	WaitForPassphraseCipherStreamProviderResponse          *WaitForPassphraseCipherStreamProviderResponse
}

// AmazonKeyManagementServiceCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns AmazonKeyManagementServiceCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func AmazonKeyManagementServiceCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *AmazonKeyManagementServiceCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		AmazonKeyManagementServiceCipherStreamProviderResponse: v,
	}
}

// AmazonSecretsManagerCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns AmazonSecretsManagerCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func AmazonSecretsManagerCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *AmazonSecretsManagerCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		AmazonSecretsManagerCipherStreamProviderResponse: v,
	}
}

// AzureKeyVaultCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns AzureKeyVaultCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func AzureKeyVaultCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *AzureKeyVaultCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		AzureKeyVaultCipherStreamProviderResponse: v,
	}
}

// ConjurCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns ConjurCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func ConjurCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *ConjurCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		ConjurCipherStreamProviderResponse: v,
	}
}

// DefaultCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns DefaultCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func DefaultCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *DefaultCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		DefaultCipherStreamProviderResponse: v,
	}
}

// FileBasedCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns FileBasedCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func FileBasedCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *FileBasedCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		FileBasedCipherStreamProviderResponse: v,
	}
}

// Pkcs11CipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns Pkcs11CipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func Pkcs11CipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *Pkcs11CipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		Pkcs11CipherStreamProviderResponse: v,
	}
}

// ThirdPartyCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns ThirdPartyCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func ThirdPartyCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *ThirdPartyCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		ThirdPartyCipherStreamProviderResponse: v,
	}
}

// VaultCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns VaultCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func VaultCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *VaultCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		VaultCipherStreamProviderResponse: v,
	}
}

// WaitForPassphraseCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner is a convenience function that returns WaitForPassphraseCipherStreamProviderResponse wrapped in CipherStreamProviderListResponseResourcesInner
func WaitForPassphraseCipherStreamProviderResponseAsCipherStreamProviderListResponseResourcesInner(v *WaitForPassphraseCipherStreamProviderResponse) CipherStreamProviderListResponseResourcesInner {
	return CipherStreamProviderListResponseResourcesInner{
		WaitForPassphraseCipherStreamProviderResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CipherStreamProviderListResponseResourcesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AmazonKeyManagementServiceCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.AmazonKeyManagementServiceCipherStreamProviderResponse)
	if err == nil {
		jsonAmazonKeyManagementServiceCipherStreamProviderResponse, _ := json.Marshal(dst.AmazonKeyManagementServiceCipherStreamProviderResponse)
		if string(jsonAmazonKeyManagementServiceCipherStreamProviderResponse) == "{}" { // empty struct
			dst.AmazonKeyManagementServiceCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.AmazonKeyManagementServiceCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into AmazonSecretsManagerCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.AmazonSecretsManagerCipherStreamProviderResponse)
	if err == nil {
		jsonAmazonSecretsManagerCipherStreamProviderResponse, _ := json.Marshal(dst.AmazonSecretsManagerCipherStreamProviderResponse)
		if string(jsonAmazonSecretsManagerCipherStreamProviderResponse) == "{}" { // empty struct
			dst.AmazonSecretsManagerCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.AmazonSecretsManagerCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into AzureKeyVaultCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.AzureKeyVaultCipherStreamProviderResponse)
	if err == nil {
		jsonAzureKeyVaultCipherStreamProviderResponse, _ := json.Marshal(dst.AzureKeyVaultCipherStreamProviderResponse)
		if string(jsonAzureKeyVaultCipherStreamProviderResponse) == "{}" { // empty struct
			dst.AzureKeyVaultCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.AzureKeyVaultCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into ConjurCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.ConjurCipherStreamProviderResponse)
	if err == nil {
		jsonConjurCipherStreamProviderResponse, _ := json.Marshal(dst.ConjurCipherStreamProviderResponse)
		if string(jsonConjurCipherStreamProviderResponse) == "{}" { // empty struct
			dst.ConjurCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.ConjurCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into DefaultCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.DefaultCipherStreamProviderResponse)
	if err == nil {
		jsonDefaultCipherStreamProviderResponse, _ := json.Marshal(dst.DefaultCipherStreamProviderResponse)
		if string(jsonDefaultCipherStreamProviderResponse) == "{}" { // empty struct
			dst.DefaultCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.DefaultCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into FileBasedCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.FileBasedCipherStreamProviderResponse)
	if err == nil {
		jsonFileBasedCipherStreamProviderResponse, _ := json.Marshal(dst.FileBasedCipherStreamProviderResponse)
		if string(jsonFileBasedCipherStreamProviderResponse) == "{}" { // empty struct
			dst.FileBasedCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.FileBasedCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into Pkcs11CipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.Pkcs11CipherStreamProviderResponse)
	if err == nil {
		jsonPkcs11CipherStreamProviderResponse, _ := json.Marshal(dst.Pkcs11CipherStreamProviderResponse)
		if string(jsonPkcs11CipherStreamProviderResponse) == "{}" { // empty struct
			dst.Pkcs11CipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.Pkcs11CipherStreamProviderResponse = nil
	}

	// try to unmarshal data into ThirdPartyCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.ThirdPartyCipherStreamProviderResponse)
	if err == nil {
		jsonThirdPartyCipherStreamProviderResponse, _ := json.Marshal(dst.ThirdPartyCipherStreamProviderResponse)
		if string(jsonThirdPartyCipherStreamProviderResponse) == "{}" { // empty struct
			dst.ThirdPartyCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.ThirdPartyCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into VaultCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.VaultCipherStreamProviderResponse)
	if err == nil {
		jsonVaultCipherStreamProviderResponse, _ := json.Marshal(dst.VaultCipherStreamProviderResponse)
		if string(jsonVaultCipherStreamProviderResponse) == "{}" { // empty struct
			dst.VaultCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.VaultCipherStreamProviderResponse = nil
	}

	// try to unmarshal data into WaitForPassphraseCipherStreamProviderResponse
	err = newStrictDecoder(data).Decode(&dst.WaitForPassphraseCipherStreamProviderResponse)
	if err == nil {
		jsonWaitForPassphraseCipherStreamProviderResponse, _ := json.Marshal(dst.WaitForPassphraseCipherStreamProviderResponse)
		if string(jsonWaitForPassphraseCipherStreamProviderResponse) == "{}" { // empty struct
			dst.WaitForPassphraseCipherStreamProviderResponse = nil
		} else {
			match++
		}
	} else {
		dst.WaitForPassphraseCipherStreamProviderResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AmazonKeyManagementServiceCipherStreamProviderResponse = nil
		dst.AmazonSecretsManagerCipherStreamProviderResponse = nil
		dst.AzureKeyVaultCipherStreamProviderResponse = nil
		dst.ConjurCipherStreamProviderResponse = nil
		dst.DefaultCipherStreamProviderResponse = nil
		dst.FileBasedCipherStreamProviderResponse = nil
		dst.Pkcs11CipherStreamProviderResponse = nil
		dst.ThirdPartyCipherStreamProviderResponse = nil
		dst.VaultCipherStreamProviderResponse = nil
		dst.WaitForPassphraseCipherStreamProviderResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CipherStreamProviderListResponseResourcesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CipherStreamProviderListResponseResourcesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CipherStreamProviderListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	if src.AmazonKeyManagementServiceCipherStreamProviderResponse != nil {
		return json.Marshal(&src.AmazonKeyManagementServiceCipherStreamProviderResponse)
	}

	if src.AmazonSecretsManagerCipherStreamProviderResponse != nil {
		return json.Marshal(&src.AmazonSecretsManagerCipherStreamProviderResponse)
	}

	if src.AzureKeyVaultCipherStreamProviderResponse != nil {
		return json.Marshal(&src.AzureKeyVaultCipherStreamProviderResponse)
	}

	if src.ConjurCipherStreamProviderResponse != nil {
		return json.Marshal(&src.ConjurCipherStreamProviderResponse)
	}

	if src.DefaultCipherStreamProviderResponse != nil {
		return json.Marshal(&src.DefaultCipherStreamProviderResponse)
	}

	if src.FileBasedCipherStreamProviderResponse != nil {
		return json.Marshal(&src.FileBasedCipherStreamProviderResponse)
	}

	if src.Pkcs11CipherStreamProviderResponse != nil {
		return json.Marshal(&src.Pkcs11CipherStreamProviderResponse)
	}

	if src.ThirdPartyCipherStreamProviderResponse != nil {
		return json.Marshal(&src.ThirdPartyCipherStreamProviderResponse)
	}

	if src.VaultCipherStreamProviderResponse != nil {
		return json.Marshal(&src.VaultCipherStreamProviderResponse)
	}

	if src.WaitForPassphraseCipherStreamProviderResponse != nil {
		return json.Marshal(&src.WaitForPassphraseCipherStreamProviderResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CipherStreamProviderListResponseResourcesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AmazonKeyManagementServiceCipherStreamProviderResponse != nil {
		return obj.AmazonKeyManagementServiceCipherStreamProviderResponse
	}

	if obj.AmazonSecretsManagerCipherStreamProviderResponse != nil {
		return obj.AmazonSecretsManagerCipherStreamProviderResponse
	}

	if obj.AzureKeyVaultCipherStreamProviderResponse != nil {
		return obj.AzureKeyVaultCipherStreamProviderResponse
	}

	if obj.ConjurCipherStreamProviderResponse != nil {
		return obj.ConjurCipherStreamProviderResponse
	}

	if obj.DefaultCipherStreamProviderResponse != nil {
		return obj.DefaultCipherStreamProviderResponse
	}

	if obj.FileBasedCipherStreamProviderResponse != nil {
		return obj.FileBasedCipherStreamProviderResponse
	}

	if obj.Pkcs11CipherStreamProviderResponse != nil {
		return obj.Pkcs11CipherStreamProviderResponse
	}

	if obj.ThirdPartyCipherStreamProviderResponse != nil {
		return obj.ThirdPartyCipherStreamProviderResponse
	}

	if obj.VaultCipherStreamProviderResponse != nil {
		return obj.VaultCipherStreamProviderResponse
	}

	if obj.WaitForPassphraseCipherStreamProviderResponse != nil {
		return obj.WaitForPassphraseCipherStreamProviderResponse
	}

	// all schemas are nil
	return nil
}

type NullableCipherStreamProviderListResponseResourcesInner struct {
	value *CipherStreamProviderListResponseResourcesInner
	isSet bool
}

func (v NullableCipherStreamProviderListResponseResourcesInner) Get() *CipherStreamProviderListResponseResourcesInner {
	return v.value
}

func (v *NullableCipherStreamProviderListResponseResourcesInner) Set(val *CipherStreamProviderListResponseResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCipherStreamProviderListResponseResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCipherStreamProviderListResponseResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCipherStreamProviderListResponseResourcesInner(val *CipherStreamProviderListResponseResourcesInner) *NullableCipherStreamProviderListResponseResourcesInner {
	return &NullableCipherStreamProviderListResponseResourcesInner{value: val, isSet: true}
}

func (v NullableCipherStreamProviderListResponseResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCipherStreamProviderListResponseResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
