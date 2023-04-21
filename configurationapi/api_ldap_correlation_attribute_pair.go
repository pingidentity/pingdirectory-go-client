/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// LdapCorrelationAttributePairApiService LdapCorrelationAttributePairApi service
type LdapCorrelationAttributePairApiService service

type ApiAddLdapCorrelationAttributePairRequest struct {
	ctx                                    context.Context
	ApiService                             *LdapCorrelationAttributePairApiService
	correlatedLdapDataViewName             string
	scimResourceTypeName                   string
	addLdapCorrelationAttributePairRequest *AddLdapCorrelationAttributePairRequest
}

// Create a new LDAP Correlation Attribute Pair in the config
func (r ApiAddLdapCorrelationAttributePairRequest) AddLdapCorrelationAttributePairRequest(addLdapCorrelationAttributePairRequest AddLdapCorrelationAttributePairRequest) ApiAddLdapCorrelationAttributePairRequest {
	r.addLdapCorrelationAttributePairRequest = &addLdapCorrelationAttributePairRequest
	return r
}

func (r ApiAddLdapCorrelationAttributePairRequest) Execute() (*LdapCorrelationAttributePairResponse, *http.Response, error) {
	return r.ApiService.AddLdapCorrelationAttributePairExecute(r)
}

/*
AddLdapCorrelationAttributePair Add a new LDAP Correlation Attribute Pair to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param correlatedLdapDataViewName Name of the Correlated LDAP Data View
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiAddLdapCorrelationAttributePairRequest
*/
func (a *LdapCorrelationAttributePairApiService) AddLdapCorrelationAttributePair(ctx context.Context, correlatedLdapDataViewName string, scimResourceTypeName string) ApiAddLdapCorrelationAttributePairRequest {
	return ApiAddLdapCorrelationAttributePairRequest{
		ApiService:                 a,
		ctx:                        ctx,
		correlatedLdapDataViewName: correlatedLdapDataViewName,
		scimResourceTypeName:       scimResourceTypeName,
	}
}

// Execute executes the request
//
//	@return LdapCorrelationAttributePairResponse
func (a *LdapCorrelationAttributePairApiService) AddLdapCorrelationAttributePairExecute(r ApiAddLdapCorrelationAttributePairRequest) (*LdapCorrelationAttributePairResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LdapCorrelationAttributePairResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LdapCorrelationAttributePairApiService.AddLdapCorrelationAttributePair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs"
	localVarPath = strings.Replace(localVarPath, "{"+"correlated-ldap-data-view-name"+"}", url.PathEscape(parameterValueToString(r.correlatedLdapDataViewName, "correlatedLdapDataViewName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addLdapCorrelationAttributePairRequest == nil {
		return localVarReturnValue, nil, reportError("addLdapCorrelationAttributePairRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addLdapCorrelationAttributePairRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteLdapCorrelationAttributePairRequest struct {
	ctx                              context.Context
	ApiService                       *LdapCorrelationAttributePairApiService
	ldapCorrelationAttributePairName string
	correlatedLdapDataViewName       string
	scimResourceTypeName             string
}

func (r ApiDeleteLdapCorrelationAttributePairRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLdapCorrelationAttributePairExecute(r)
}

/*
DeleteLdapCorrelationAttributePair Delete a LDAP Correlation Attribute Pair

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ldapCorrelationAttributePairName Name of the LDAP Correlation Attribute Pair
	@param correlatedLdapDataViewName Name of the Correlated LDAP Data View
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiDeleteLdapCorrelationAttributePairRequest
*/
func (a *LdapCorrelationAttributePairApiService) DeleteLdapCorrelationAttributePair(ctx context.Context, ldapCorrelationAttributePairName string, correlatedLdapDataViewName string, scimResourceTypeName string) ApiDeleteLdapCorrelationAttributePairRequest {
	return ApiDeleteLdapCorrelationAttributePairRequest{
		ApiService:                       a,
		ctx:                              ctx,
		ldapCorrelationAttributePairName: ldapCorrelationAttributePairName,
		correlatedLdapDataViewName:       correlatedLdapDataViewName,
		scimResourceTypeName:             scimResourceTypeName,
	}
}

// Execute executes the request
func (a *LdapCorrelationAttributePairApiService) DeleteLdapCorrelationAttributePairExecute(r ApiDeleteLdapCorrelationAttributePairRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LdapCorrelationAttributePairApiService.DeleteLdapCorrelationAttributePair")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs/{ldap-correlation-attribute-pair-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"ldap-correlation-attribute-pair-name"+"}", url.PathEscape(parameterValueToString(r.ldapCorrelationAttributePairName, "ldapCorrelationAttributePairName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"correlated-ldap-data-view-name"+"}", url.PathEscape(parameterValueToString(r.correlatedLdapDataViewName, "correlatedLdapDataViewName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetLdapCorrelationAttributePairRequest struct {
	ctx                              context.Context
	ApiService                       *LdapCorrelationAttributePairApiService
	ldapCorrelationAttributePairName string
	correlatedLdapDataViewName       string
	scimResourceTypeName             string
}

func (r ApiGetLdapCorrelationAttributePairRequest) Execute() (*LdapCorrelationAttributePairResponse, *http.Response, error) {
	return r.ApiService.GetLdapCorrelationAttributePairExecute(r)
}

/*
GetLdapCorrelationAttributePair Returns a single LDAP Correlation Attribute Pair

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ldapCorrelationAttributePairName Name of the LDAP Correlation Attribute Pair
	@param correlatedLdapDataViewName Name of the Correlated LDAP Data View
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiGetLdapCorrelationAttributePairRequest
*/
func (a *LdapCorrelationAttributePairApiService) GetLdapCorrelationAttributePair(ctx context.Context, ldapCorrelationAttributePairName string, correlatedLdapDataViewName string, scimResourceTypeName string) ApiGetLdapCorrelationAttributePairRequest {
	return ApiGetLdapCorrelationAttributePairRequest{
		ApiService:                       a,
		ctx:                              ctx,
		ldapCorrelationAttributePairName: ldapCorrelationAttributePairName,
		correlatedLdapDataViewName:       correlatedLdapDataViewName,
		scimResourceTypeName:             scimResourceTypeName,
	}
}

// Execute executes the request
//
//	@return LdapCorrelationAttributePairResponse
func (a *LdapCorrelationAttributePairApiService) GetLdapCorrelationAttributePairExecute(r ApiGetLdapCorrelationAttributePairRequest) (*LdapCorrelationAttributePairResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LdapCorrelationAttributePairResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LdapCorrelationAttributePairApiService.GetLdapCorrelationAttributePair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs/{ldap-correlation-attribute-pair-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"ldap-correlation-attribute-pair-name"+"}", url.PathEscape(parameterValueToString(r.ldapCorrelationAttributePairName, "ldapCorrelationAttributePairName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"correlated-ldap-data-view-name"+"}", url.PathEscape(parameterValueToString(r.correlatedLdapDataViewName, "correlatedLdapDataViewName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateLdapCorrelationAttributePairRequest struct {
	ctx                              context.Context
	ApiService                       *LdapCorrelationAttributePairApiService
	ldapCorrelationAttributePairName string
	correlatedLdapDataViewName       string
	scimResourceTypeName             string
	updateRequest                    *UpdateRequest
}

// Update an existing LDAP Correlation Attribute Pair
func (r ApiUpdateLdapCorrelationAttributePairRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateLdapCorrelationAttributePairRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateLdapCorrelationAttributePairRequest) Execute() (*LdapCorrelationAttributePairResponse, *http.Response, error) {
	return r.ApiService.UpdateLdapCorrelationAttributePairExecute(r)
}

/*
UpdateLdapCorrelationAttributePair Update an existing LDAP Correlation Attribute Pair by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ldapCorrelationAttributePairName Name of the LDAP Correlation Attribute Pair
	@param correlatedLdapDataViewName Name of the Correlated LDAP Data View
	@param scimResourceTypeName Name of the SCIM Resource Type
	@return ApiUpdateLdapCorrelationAttributePairRequest
*/
func (a *LdapCorrelationAttributePairApiService) UpdateLdapCorrelationAttributePair(ctx context.Context, ldapCorrelationAttributePairName string, correlatedLdapDataViewName string, scimResourceTypeName string) ApiUpdateLdapCorrelationAttributePairRequest {
	return ApiUpdateLdapCorrelationAttributePairRequest{
		ApiService:                       a,
		ctx:                              ctx,
		ldapCorrelationAttributePairName: ldapCorrelationAttributePairName,
		correlatedLdapDataViewName:       correlatedLdapDataViewName,
		scimResourceTypeName:             scimResourceTypeName,
	}
}

// Execute executes the request
//
//	@return LdapCorrelationAttributePairResponse
func (a *LdapCorrelationAttributePairApiService) UpdateLdapCorrelationAttributePairExecute(r ApiUpdateLdapCorrelationAttributePairRequest) (*LdapCorrelationAttributePairResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LdapCorrelationAttributePairResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LdapCorrelationAttributePairApiService.UpdateLdapCorrelationAttributePair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs/{ldap-correlation-attribute-pair-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"ldap-correlation-attribute-pair-name"+"}", url.PathEscape(parameterValueToString(r.ldapCorrelationAttributePairName, "ldapCorrelationAttributePairName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"correlated-ldap-data-view-name"+"}", url.PathEscape(parameterValueToString(r.correlatedLdapDataViewName, "correlatedLdapDataViewName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scim-resource-type-name"+"}", url.PathEscape(parameterValueToString(r.scimResourceTypeName, "scimResourceTypeName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRequest == nil {
		return localVarReturnValue, nil, reportError("updateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
