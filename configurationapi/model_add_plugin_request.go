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

// AddPluginRequest - struct for AddPluginRequest
type AddPluginRequest struct {
	AddAttributeMapperPluginRequest                                  *AddAttributeMapperPluginRequest
	AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest
	AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest     *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest
	AddCleanUpInactivePingfederatePersistentSessionsPluginRequest    *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest
	AddComposedAttributePluginRequest                                *AddComposedAttributePluginRequest
	AddDelayPluginRequest                                            *AddDelayPluginRequest
	AddDnMapperPluginRequest                                         *AddDnMapperPluginRequest
	AddGroovyScriptedPluginRequest                                   *AddGroovyScriptedPluginRequest
	AddInternalSearchRatePluginRequest                               *AddInternalSearchRatePluginRequest
	AddModifiablePasswordPolicyStatePluginRequest                    *AddModifiablePasswordPolicyStatePluginRequest
	AddPassThroughAuthenticationPluginRequest                        *AddPassThroughAuthenticationPluginRequest
	AddPeriodicGcPluginRequest                                       *AddPeriodicGcPluginRequest
	AddPeriodicStatsLoggerPluginRequest                              *AddPeriodicStatsLoggerPluginRequest
	AddPingOnePassThroughAuthenticationPluginRequest                 *AddPingOnePassThroughAuthenticationPluginRequest
	AddPluggablePassThroughAuthenticationPluginRequest               *AddPluggablePassThroughAuthenticationPluginRequest
	AddPurgeExpiredDataPluginRequest                                 *AddPurgeExpiredDataPluginRequest
	AddReferentialIntegrityPluginRequest                             *AddReferentialIntegrityPluginRequest
	AddReferralOnUpdatePluginRequest                                 *AddReferralOnUpdatePluginRequest
	AddSearchShutdownPluginRequest                                   *AddSearchShutdownPluginRequest
	AddSevenBitCleanPluginRequest                                    *AddSevenBitCleanPluginRequest
	AddSimpleToExternalBindPluginRequest                             *AddSimpleToExternalBindPluginRequest
	AddSnmpSubagentPluginRequest                                     *AddSnmpSubagentPluginRequest
	AddSubOperationTimingPluginRequest                               *AddSubOperationTimingPluginRequest
	AddThirdPartyPluginRequest                                       *AddThirdPartyPluginRequest
	AddUniqueAttributePluginRequest                                  *AddUniqueAttributePluginRequest
}

// AddAttributeMapperPluginRequestAsAddPluginRequest is a convenience function that returns AddAttributeMapperPluginRequest wrapped in AddPluginRequest
func AddAttributeMapperPluginRequestAsAddPluginRequest(v *AddAttributeMapperPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddAttributeMapperPluginRequest: v,
	}
}

// AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequestAsAddPluginRequest is a convenience function that returns AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest wrapped in AddPluginRequest
func AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequestAsAddPluginRequest(v *AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest: v,
	}
}

// AddCleanUpExpiredPingfederatePersistentSessionsPluginRequestAsAddPluginRequest is a convenience function that returns AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest wrapped in AddPluginRequest
func AddCleanUpExpiredPingfederatePersistentSessionsPluginRequestAsAddPluginRequest(v *AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest: v,
	}
}

// AddCleanUpInactivePingfederatePersistentSessionsPluginRequestAsAddPluginRequest is a convenience function that returns AddCleanUpInactivePingfederatePersistentSessionsPluginRequest wrapped in AddPluginRequest
func AddCleanUpInactivePingfederatePersistentSessionsPluginRequestAsAddPluginRequest(v *AddCleanUpInactivePingfederatePersistentSessionsPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddCleanUpInactivePingfederatePersistentSessionsPluginRequest: v,
	}
}

// AddComposedAttributePluginRequestAsAddPluginRequest is a convenience function that returns AddComposedAttributePluginRequest wrapped in AddPluginRequest
func AddComposedAttributePluginRequestAsAddPluginRequest(v *AddComposedAttributePluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddComposedAttributePluginRequest: v,
	}
}

// AddDelayPluginRequestAsAddPluginRequest is a convenience function that returns AddDelayPluginRequest wrapped in AddPluginRequest
func AddDelayPluginRequestAsAddPluginRequest(v *AddDelayPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddDelayPluginRequest: v,
	}
}

// AddDnMapperPluginRequestAsAddPluginRequest is a convenience function that returns AddDnMapperPluginRequest wrapped in AddPluginRequest
func AddDnMapperPluginRequestAsAddPluginRequest(v *AddDnMapperPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddDnMapperPluginRequest: v,
	}
}

// AddGroovyScriptedPluginRequestAsAddPluginRequest is a convenience function that returns AddGroovyScriptedPluginRequest wrapped in AddPluginRequest
func AddGroovyScriptedPluginRequestAsAddPluginRequest(v *AddGroovyScriptedPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddGroovyScriptedPluginRequest: v,
	}
}

// AddInternalSearchRatePluginRequestAsAddPluginRequest is a convenience function that returns AddInternalSearchRatePluginRequest wrapped in AddPluginRequest
func AddInternalSearchRatePluginRequestAsAddPluginRequest(v *AddInternalSearchRatePluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddInternalSearchRatePluginRequest: v,
	}
}

// AddModifiablePasswordPolicyStatePluginRequestAsAddPluginRequest is a convenience function that returns AddModifiablePasswordPolicyStatePluginRequest wrapped in AddPluginRequest
func AddModifiablePasswordPolicyStatePluginRequestAsAddPluginRequest(v *AddModifiablePasswordPolicyStatePluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddModifiablePasswordPolicyStatePluginRequest: v,
	}
}

// AddPassThroughAuthenticationPluginRequestAsAddPluginRequest is a convenience function that returns AddPassThroughAuthenticationPluginRequest wrapped in AddPluginRequest
func AddPassThroughAuthenticationPluginRequestAsAddPluginRequest(v *AddPassThroughAuthenticationPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddPassThroughAuthenticationPluginRequest: v,
	}
}

// AddPeriodicGcPluginRequestAsAddPluginRequest is a convenience function that returns AddPeriodicGcPluginRequest wrapped in AddPluginRequest
func AddPeriodicGcPluginRequestAsAddPluginRequest(v *AddPeriodicGcPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddPeriodicGcPluginRequest: v,
	}
}

// AddPeriodicStatsLoggerPluginRequestAsAddPluginRequest is a convenience function that returns AddPeriodicStatsLoggerPluginRequest wrapped in AddPluginRequest
func AddPeriodicStatsLoggerPluginRequestAsAddPluginRequest(v *AddPeriodicStatsLoggerPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddPeriodicStatsLoggerPluginRequest: v,
	}
}

// AddPingOnePassThroughAuthenticationPluginRequestAsAddPluginRequest is a convenience function that returns AddPingOnePassThroughAuthenticationPluginRequest wrapped in AddPluginRequest
func AddPingOnePassThroughAuthenticationPluginRequestAsAddPluginRequest(v *AddPingOnePassThroughAuthenticationPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddPingOnePassThroughAuthenticationPluginRequest: v,
	}
}

// AddPluggablePassThroughAuthenticationPluginRequestAsAddPluginRequest is a convenience function that returns AddPluggablePassThroughAuthenticationPluginRequest wrapped in AddPluginRequest
func AddPluggablePassThroughAuthenticationPluginRequestAsAddPluginRequest(v *AddPluggablePassThroughAuthenticationPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddPluggablePassThroughAuthenticationPluginRequest: v,
	}
}

// AddPurgeExpiredDataPluginRequestAsAddPluginRequest is a convenience function that returns AddPurgeExpiredDataPluginRequest wrapped in AddPluginRequest
func AddPurgeExpiredDataPluginRequestAsAddPluginRequest(v *AddPurgeExpiredDataPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddPurgeExpiredDataPluginRequest: v,
	}
}

// AddReferentialIntegrityPluginRequestAsAddPluginRequest is a convenience function that returns AddReferentialIntegrityPluginRequest wrapped in AddPluginRequest
func AddReferentialIntegrityPluginRequestAsAddPluginRequest(v *AddReferentialIntegrityPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddReferentialIntegrityPluginRequest: v,
	}
}

// AddReferralOnUpdatePluginRequestAsAddPluginRequest is a convenience function that returns AddReferralOnUpdatePluginRequest wrapped in AddPluginRequest
func AddReferralOnUpdatePluginRequestAsAddPluginRequest(v *AddReferralOnUpdatePluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddReferralOnUpdatePluginRequest: v,
	}
}

// AddSearchShutdownPluginRequestAsAddPluginRequest is a convenience function that returns AddSearchShutdownPluginRequest wrapped in AddPluginRequest
func AddSearchShutdownPluginRequestAsAddPluginRequest(v *AddSearchShutdownPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddSearchShutdownPluginRequest: v,
	}
}

// AddSevenBitCleanPluginRequestAsAddPluginRequest is a convenience function that returns AddSevenBitCleanPluginRequest wrapped in AddPluginRequest
func AddSevenBitCleanPluginRequestAsAddPluginRequest(v *AddSevenBitCleanPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddSevenBitCleanPluginRequest: v,
	}
}

// AddSimpleToExternalBindPluginRequestAsAddPluginRequest is a convenience function that returns AddSimpleToExternalBindPluginRequest wrapped in AddPluginRequest
func AddSimpleToExternalBindPluginRequestAsAddPluginRequest(v *AddSimpleToExternalBindPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddSimpleToExternalBindPluginRequest: v,
	}
}

// AddSnmpSubagentPluginRequestAsAddPluginRequest is a convenience function that returns AddSnmpSubagentPluginRequest wrapped in AddPluginRequest
func AddSnmpSubagentPluginRequestAsAddPluginRequest(v *AddSnmpSubagentPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddSnmpSubagentPluginRequest: v,
	}
}

// AddSubOperationTimingPluginRequestAsAddPluginRequest is a convenience function that returns AddSubOperationTimingPluginRequest wrapped in AddPluginRequest
func AddSubOperationTimingPluginRequestAsAddPluginRequest(v *AddSubOperationTimingPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddSubOperationTimingPluginRequest: v,
	}
}

// AddThirdPartyPluginRequestAsAddPluginRequest is a convenience function that returns AddThirdPartyPluginRequest wrapped in AddPluginRequest
func AddThirdPartyPluginRequestAsAddPluginRequest(v *AddThirdPartyPluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddThirdPartyPluginRequest: v,
	}
}

// AddUniqueAttributePluginRequestAsAddPluginRequest is a convenience function that returns AddUniqueAttributePluginRequest wrapped in AddPluginRequest
func AddUniqueAttributePluginRequestAsAddPluginRequest(v *AddUniqueAttributePluginRequest) AddPluginRequest {
	return AddPluginRequest{
		AddUniqueAttributePluginRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddPluginRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddAttributeMapperPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddAttributeMapperPluginRequest)
	if err == nil {
		jsonAddAttributeMapperPluginRequest, _ := json.Marshal(dst.AddAttributeMapperPluginRequest)
		if string(jsonAddAttributeMapperPluginRequest) == "{}" { // empty struct
			dst.AddAttributeMapperPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddAttributeMapperPluginRequest = nil
	}

	// try to unmarshal data into AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest)
	if err == nil {
		jsonAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest, _ := json.Marshal(dst.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest)
		if string(jsonAddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest) == "{}" { // empty struct
			dst.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest = nil
	}

	// try to unmarshal data into AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest)
	if err == nil {
		jsonAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest, _ := json.Marshal(dst.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest)
		if string(jsonAddCleanUpExpiredPingfederatePersistentSessionsPluginRequest) == "{}" { // empty struct
			dst.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest = nil
	}

	// try to unmarshal data into AddCleanUpInactivePingfederatePersistentSessionsPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest)
	if err == nil {
		jsonAddCleanUpInactivePingfederatePersistentSessionsPluginRequest, _ := json.Marshal(dst.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest)
		if string(jsonAddCleanUpInactivePingfederatePersistentSessionsPluginRequest) == "{}" { // empty struct
			dst.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest = nil
	}

	// try to unmarshal data into AddComposedAttributePluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddComposedAttributePluginRequest)
	if err == nil {
		jsonAddComposedAttributePluginRequest, _ := json.Marshal(dst.AddComposedAttributePluginRequest)
		if string(jsonAddComposedAttributePluginRequest) == "{}" { // empty struct
			dst.AddComposedAttributePluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddComposedAttributePluginRequest = nil
	}

	// try to unmarshal data into AddDelayPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddDelayPluginRequest)
	if err == nil {
		jsonAddDelayPluginRequest, _ := json.Marshal(dst.AddDelayPluginRequest)
		if string(jsonAddDelayPluginRequest) == "{}" { // empty struct
			dst.AddDelayPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddDelayPluginRequest = nil
	}

	// try to unmarshal data into AddDnMapperPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddDnMapperPluginRequest)
	if err == nil {
		jsonAddDnMapperPluginRequest, _ := json.Marshal(dst.AddDnMapperPluginRequest)
		if string(jsonAddDnMapperPluginRequest) == "{}" { // empty struct
			dst.AddDnMapperPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddDnMapperPluginRequest = nil
	}

	// try to unmarshal data into AddGroovyScriptedPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddGroovyScriptedPluginRequest)
	if err == nil {
		jsonAddGroovyScriptedPluginRequest, _ := json.Marshal(dst.AddGroovyScriptedPluginRequest)
		if string(jsonAddGroovyScriptedPluginRequest) == "{}" { // empty struct
			dst.AddGroovyScriptedPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddGroovyScriptedPluginRequest = nil
	}

	// try to unmarshal data into AddInternalSearchRatePluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddInternalSearchRatePluginRequest)
	if err == nil {
		jsonAddInternalSearchRatePluginRequest, _ := json.Marshal(dst.AddInternalSearchRatePluginRequest)
		if string(jsonAddInternalSearchRatePluginRequest) == "{}" { // empty struct
			dst.AddInternalSearchRatePluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddInternalSearchRatePluginRequest = nil
	}

	// try to unmarshal data into AddModifiablePasswordPolicyStatePluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddModifiablePasswordPolicyStatePluginRequest)
	if err == nil {
		jsonAddModifiablePasswordPolicyStatePluginRequest, _ := json.Marshal(dst.AddModifiablePasswordPolicyStatePluginRequest)
		if string(jsonAddModifiablePasswordPolicyStatePluginRequest) == "{}" { // empty struct
			dst.AddModifiablePasswordPolicyStatePluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddModifiablePasswordPolicyStatePluginRequest = nil
	}

	// try to unmarshal data into AddPassThroughAuthenticationPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddPassThroughAuthenticationPluginRequest)
	if err == nil {
		jsonAddPassThroughAuthenticationPluginRequest, _ := json.Marshal(dst.AddPassThroughAuthenticationPluginRequest)
		if string(jsonAddPassThroughAuthenticationPluginRequest) == "{}" { // empty struct
			dst.AddPassThroughAuthenticationPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPassThroughAuthenticationPluginRequest = nil
	}

	// try to unmarshal data into AddPeriodicGcPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddPeriodicGcPluginRequest)
	if err == nil {
		jsonAddPeriodicGcPluginRequest, _ := json.Marshal(dst.AddPeriodicGcPluginRequest)
		if string(jsonAddPeriodicGcPluginRequest) == "{}" { // empty struct
			dst.AddPeriodicGcPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPeriodicGcPluginRequest = nil
	}

	// try to unmarshal data into AddPeriodicStatsLoggerPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddPeriodicStatsLoggerPluginRequest)
	if err == nil {
		jsonAddPeriodicStatsLoggerPluginRequest, _ := json.Marshal(dst.AddPeriodicStatsLoggerPluginRequest)
		if string(jsonAddPeriodicStatsLoggerPluginRequest) == "{}" { // empty struct
			dst.AddPeriodicStatsLoggerPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPeriodicStatsLoggerPluginRequest = nil
	}

	// try to unmarshal data into AddPingOnePassThroughAuthenticationPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddPingOnePassThroughAuthenticationPluginRequest)
	if err == nil {
		jsonAddPingOnePassThroughAuthenticationPluginRequest, _ := json.Marshal(dst.AddPingOnePassThroughAuthenticationPluginRequest)
		if string(jsonAddPingOnePassThroughAuthenticationPluginRequest) == "{}" { // empty struct
			dst.AddPingOnePassThroughAuthenticationPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPingOnePassThroughAuthenticationPluginRequest = nil
	}

	// try to unmarshal data into AddPluggablePassThroughAuthenticationPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddPluggablePassThroughAuthenticationPluginRequest)
	if err == nil {
		jsonAddPluggablePassThroughAuthenticationPluginRequest, _ := json.Marshal(dst.AddPluggablePassThroughAuthenticationPluginRequest)
		if string(jsonAddPluggablePassThroughAuthenticationPluginRequest) == "{}" { // empty struct
			dst.AddPluggablePassThroughAuthenticationPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPluggablePassThroughAuthenticationPluginRequest = nil
	}

	// try to unmarshal data into AddPurgeExpiredDataPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddPurgeExpiredDataPluginRequest)
	if err == nil {
		jsonAddPurgeExpiredDataPluginRequest, _ := json.Marshal(dst.AddPurgeExpiredDataPluginRequest)
		if string(jsonAddPurgeExpiredDataPluginRequest) == "{}" { // empty struct
			dst.AddPurgeExpiredDataPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddPurgeExpiredDataPluginRequest = nil
	}

	// try to unmarshal data into AddReferentialIntegrityPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddReferentialIntegrityPluginRequest)
	if err == nil {
		jsonAddReferentialIntegrityPluginRequest, _ := json.Marshal(dst.AddReferentialIntegrityPluginRequest)
		if string(jsonAddReferentialIntegrityPluginRequest) == "{}" { // empty struct
			dst.AddReferentialIntegrityPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddReferentialIntegrityPluginRequest = nil
	}

	// try to unmarshal data into AddReferralOnUpdatePluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddReferralOnUpdatePluginRequest)
	if err == nil {
		jsonAddReferralOnUpdatePluginRequest, _ := json.Marshal(dst.AddReferralOnUpdatePluginRequest)
		if string(jsonAddReferralOnUpdatePluginRequest) == "{}" { // empty struct
			dst.AddReferralOnUpdatePluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddReferralOnUpdatePluginRequest = nil
	}

	// try to unmarshal data into AddSearchShutdownPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddSearchShutdownPluginRequest)
	if err == nil {
		jsonAddSearchShutdownPluginRequest, _ := json.Marshal(dst.AddSearchShutdownPluginRequest)
		if string(jsonAddSearchShutdownPluginRequest) == "{}" { // empty struct
			dst.AddSearchShutdownPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSearchShutdownPluginRequest = nil
	}

	// try to unmarshal data into AddSevenBitCleanPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddSevenBitCleanPluginRequest)
	if err == nil {
		jsonAddSevenBitCleanPluginRequest, _ := json.Marshal(dst.AddSevenBitCleanPluginRequest)
		if string(jsonAddSevenBitCleanPluginRequest) == "{}" { // empty struct
			dst.AddSevenBitCleanPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSevenBitCleanPluginRequest = nil
	}

	// try to unmarshal data into AddSimpleToExternalBindPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddSimpleToExternalBindPluginRequest)
	if err == nil {
		jsonAddSimpleToExternalBindPluginRequest, _ := json.Marshal(dst.AddSimpleToExternalBindPluginRequest)
		if string(jsonAddSimpleToExternalBindPluginRequest) == "{}" { // empty struct
			dst.AddSimpleToExternalBindPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSimpleToExternalBindPluginRequest = nil
	}

	// try to unmarshal data into AddSnmpSubagentPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddSnmpSubagentPluginRequest)
	if err == nil {
		jsonAddSnmpSubagentPluginRequest, _ := json.Marshal(dst.AddSnmpSubagentPluginRequest)
		if string(jsonAddSnmpSubagentPluginRequest) == "{}" { // empty struct
			dst.AddSnmpSubagentPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSnmpSubagentPluginRequest = nil
	}

	// try to unmarshal data into AddSubOperationTimingPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddSubOperationTimingPluginRequest)
	if err == nil {
		jsonAddSubOperationTimingPluginRequest, _ := json.Marshal(dst.AddSubOperationTimingPluginRequest)
		if string(jsonAddSubOperationTimingPluginRequest) == "{}" { // empty struct
			dst.AddSubOperationTimingPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddSubOperationTimingPluginRequest = nil
	}

	// try to unmarshal data into AddThirdPartyPluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddThirdPartyPluginRequest)
	if err == nil {
		jsonAddThirdPartyPluginRequest, _ := json.Marshal(dst.AddThirdPartyPluginRequest)
		if string(jsonAddThirdPartyPluginRequest) == "{}" { // empty struct
			dst.AddThirdPartyPluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddThirdPartyPluginRequest = nil
	}

	// try to unmarshal data into AddUniqueAttributePluginRequest
	err = newStrictDecoder(data).Decode(&dst.AddUniqueAttributePluginRequest)
	if err == nil {
		jsonAddUniqueAttributePluginRequest, _ := json.Marshal(dst.AddUniqueAttributePluginRequest)
		if string(jsonAddUniqueAttributePluginRequest) == "{}" { // empty struct
			dst.AddUniqueAttributePluginRequest = nil
		} else {
			match++
		}
	} else {
		dst.AddUniqueAttributePluginRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddAttributeMapperPluginRequest = nil
		dst.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest = nil
		dst.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest = nil
		dst.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest = nil
		dst.AddComposedAttributePluginRequest = nil
		dst.AddDelayPluginRequest = nil
		dst.AddDnMapperPluginRequest = nil
		dst.AddGroovyScriptedPluginRequest = nil
		dst.AddInternalSearchRatePluginRequest = nil
		dst.AddModifiablePasswordPolicyStatePluginRequest = nil
		dst.AddPassThroughAuthenticationPluginRequest = nil
		dst.AddPeriodicGcPluginRequest = nil
		dst.AddPeriodicStatsLoggerPluginRequest = nil
		dst.AddPingOnePassThroughAuthenticationPluginRequest = nil
		dst.AddPluggablePassThroughAuthenticationPluginRequest = nil
		dst.AddPurgeExpiredDataPluginRequest = nil
		dst.AddReferentialIntegrityPluginRequest = nil
		dst.AddReferralOnUpdatePluginRequest = nil
		dst.AddSearchShutdownPluginRequest = nil
		dst.AddSevenBitCleanPluginRequest = nil
		dst.AddSimpleToExternalBindPluginRequest = nil
		dst.AddSnmpSubagentPluginRequest = nil
		dst.AddSubOperationTimingPluginRequest = nil
		dst.AddThirdPartyPluginRequest = nil
		dst.AddUniqueAttributePluginRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AddPluginRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AddPluginRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddPluginRequest) MarshalJSON() ([]byte, error) {
	if src.AddAttributeMapperPluginRequest != nil {
		return json.Marshal(&src.AddAttributeMapperPluginRequest)
	}

	if src.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest != nil {
		return json.Marshal(&src.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest)
	}

	if src.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest != nil {
		return json.Marshal(&src.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest)
	}

	if src.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest != nil {
		return json.Marshal(&src.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest)
	}

	if src.AddComposedAttributePluginRequest != nil {
		return json.Marshal(&src.AddComposedAttributePluginRequest)
	}

	if src.AddDelayPluginRequest != nil {
		return json.Marshal(&src.AddDelayPluginRequest)
	}

	if src.AddDnMapperPluginRequest != nil {
		return json.Marshal(&src.AddDnMapperPluginRequest)
	}

	if src.AddGroovyScriptedPluginRequest != nil {
		return json.Marshal(&src.AddGroovyScriptedPluginRequest)
	}

	if src.AddInternalSearchRatePluginRequest != nil {
		return json.Marshal(&src.AddInternalSearchRatePluginRequest)
	}

	if src.AddModifiablePasswordPolicyStatePluginRequest != nil {
		return json.Marshal(&src.AddModifiablePasswordPolicyStatePluginRequest)
	}

	if src.AddPassThroughAuthenticationPluginRequest != nil {
		return json.Marshal(&src.AddPassThroughAuthenticationPluginRequest)
	}

	if src.AddPeriodicGcPluginRequest != nil {
		return json.Marshal(&src.AddPeriodicGcPluginRequest)
	}

	if src.AddPeriodicStatsLoggerPluginRequest != nil {
		return json.Marshal(&src.AddPeriodicStatsLoggerPluginRequest)
	}

	if src.AddPingOnePassThroughAuthenticationPluginRequest != nil {
		return json.Marshal(&src.AddPingOnePassThroughAuthenticationPluginRequest)
	}

	if src.AddPluggablePassThroughAuthenticationPluginRequest != nil {
		return json.Marshal(&src.AddPluggablePassThroughAuthenticationPluginRequest)
	}

	if src.AddPurgeExpiredDataPluginRequest != nil {
		return json.Marshal(&src.AddPurgeExpiredDataPluginRequest)
	}

	if src.AddReferentialIntegrityPluginRequest != nil {
		return json.Marshal(&src.AddReferentialIntegrityPluginRequest)
	}

	if src.AddReferralOnUpdatePluginRequest != nil {
		return json.Marshal(&src.AddReferralOnUpdatePluginRequest)
	}

	if src.AddSearchShutdownPluginRequest != nil {
		return json.Marshal(&src.AddSearchShutdownPluginRequest)
	}

	if src.AddSevenBitCleanPluginRequest != nil {
		return json.Marshal(&src.AddSevenBitCleanPluginRequest)
	}

	if src.AddSimpleToExternalBindPluginRequest != nil {
		return json.Marshal(&src.AddSimpleToExternalBindPluginRequest)
	}

	if src.AddSnmpSubagentPluginRequest != nil {
		return json.Marshal(&src.AddSnmpSubagentPluginRequest)
	}

	if src.AddSubOperationTimingPluginRequest != nil {
		return json.Marshal(&src.AddSubOperationTimingPluginRequest)
	}

	if src.AddThirdPartyPluginRequest != nil {
		return json.Marshal(&src.AddThirdPartyPluginRequest)
	}

	if src.AddUniqueAttributePluginRequest != nil {
		return json.Marshal(&src.AddUniqueAttributePluginRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddPluginRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddAttributeMapperPluginRequest != nil {
		return obj.AddAttributeMapperPluginRequest
	}

	if obj.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest != nil {
		return obj.AddCleanUpExpiredPingfederatePersistentAccessGrantsPluginRequest
	}

	if obj.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest != nil {
		return obj.AddCleanUpExpiredPingfederatePersistentSessionsPluginRequest
	}

	if obj.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest != nil {
		return obj.AddCleanUpInactivePingfederatePersistentSessionsPluginRequest
	}

	if obj.AddComposedAttributePluginRequest != nil {
		return obj.AddComposedAttributePluginRequest
	}

	if obj.AddDelayPluginRequest != nil {
		return obj.AddDelayPluginRequest
	}

	if obj.AddDnMapperPluginRequest != nil {
		return obj.AddDnMapperPluginRequest
	}

	if obj.AddGroovyScriptedPluginRequest != nil {
		return obj.AddGroovyScriptedPluginRequest
	}

	if obj.AddInternalSearchRatePluginRequest != nil {
		return obj.AddInternalSearchRatePluginRequest
	}

	if obj.AddModifiablePasswordPolicyStatePluginRequest != nil {
		return obj.AddModifiablePasswordPolicyStatePluginRequest
	}

	if obj.AddPassThroughAuthenticationPluginRequest != nil {
		return obj.AddPassThroughAuthenticationPluginRequest
	}

	if obj.AddPeriodicGcPluginRequest != nil {
		return obj.AddPeriodicGcPluginRequest
	}

	if obj.AddPeriodicStatsLoggerPluginRequest != nil {
		return obj.AddPeriodicStatsLoggerPluginRequest
	}

	if obj.AddPingOnePassThroughAuthenticationPluginRequest != nil {
		return obj.AddPingOnePassThroughAuthenticationPluginRequest
	}

	if obj.AddPluggablePassThroughAuthenticationPluginRequest != nil {
		return obj.AddPluggablePassThroughAuthenticationPluginRequest
	}

	if obj.AddPurgeExpiredDataPluginRequest != nil {
		return obj.AddPurgeExpiredDataPluginRequest
	}

	if obj.AddReferentialIntegrityPluginRequest != nil {
		return obj.AddReferentialIntegrityPluginRequest
	}

	if obj.AddReferralOnUpdatePluginRequest != nil {
		return obj.AddReferralOnUpdatePluginRequest
	}

	if obj.AddSearchShutdownPluginRequest != nil {
		return obj.AddSearchShutdownPluginRequest
	}

	if obj.AddSevenBitCleanPluginRequest != nil {
		return obj.AddSevenBitCleanPluginRequest
	}

	if obj.AddSimpleToExternalBindPluginRequest != nil {
		return obj.AddSimpleToExternalBindPluginRequest
	}

	if obj.AddSnmpSubagentPluginRequest != nil {
		return obj.AddSnmpSubagentPluginRequest
	}

	if obj.AddSubOperationTimingPluginRequest != nil {
		return obj.AddSubOperationTimingPluginRequest
	}

	if obj.AddThirdPartyPluginRequest != nil {
		return obj.AddThirdPartyPluginRequest
	}

	if obj.AddUniqueAttributePluginRequest != nil {
		return obj.AddUniqueAttributePluginRequest
	}

	// all schemas are nil
	return nil
}

type NullableAddPluginRequest struct {
	value *AddPluginRequest
	isSet bool
}

func (v NullableAddPluginRequest) Get() *AddPluginRequest {
	return v.value
}

func (v *NullableAddPluginRequest) Set(val *AddPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPluginRequest(val *AddPluginRequest) *NullableAddPluginRequest {
	return &NullableAddPluginRequest{value: val, isSet: true}
}

func (v NullableAddPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}