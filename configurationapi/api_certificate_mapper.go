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

// CertificateMapperAPIService CertificateMapperAPI service
type CertificateMapperAPIService service

type ApiAddCertificateMapperRequest struct {
	ctx                         context.Context
	ApiService                  *CertificateMapperAPIService
	addCertificateMapperRequest *AddCertificateMapperRequest
}

// Create a new Certificate Mapper in the config
func (r ApiAddCertificateMapperRequest) AddCertificateMapperRequest(addCertificateMapperRequest AddCertificateMapperRequest) ApiAddCertificateMapperRequest {
	r.addCertificateMapperRequest = &addCertificateMapperRequest
	return r
}

func (r ApiAddCertificateMapperRequest) Execute() (*AddCertificateMapper200Response, *http.Response, error) {
	return r.ApiService.AddCertificateMapperExecute(r)
}

/*
AddCertificateMapper Add a new Certificate Mapper to the config

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddCertificateMapperRequest
*/
func (a *CertificateMapperAPIService) AddCertificateMapper(ctx context.Context) ApiAddCertificateMapperRequest {
	return ApiAddCertificateMapperRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddCertificateMapper200Response
func (a *CertificateMapperAPIService) AddCertificateMapperExecute(r ApiAddCertificateMapperRequest) (*AddCertificateMapper200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCertificateMapper200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateMapperAPIService.AddCertificateMapper")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificate-mappers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addCertificateMapperRequest == nil {
		return localVarReturnValue, nil, reportError("addCertificateMapperRequest is required and must be specified")
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
	localVarPostBody = r.addCertificateMapperRequest
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

type ApiDeleteCertificateMapperRequest struct {
	ctx                   context.Context
	ApiService            *CertificateMapperAPIService
	certificateMapperName string
}

func (r ApiDeleteCertificateMapperRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCertificateMapperExecute(r)
}

/*
DeleteCertificateMapper Delete a Certificate Mapper

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param certificateMapperName Name of the Certificate Mapper
	@return ApiDeleteCertificateMapperRequest
*/
func (a *CertificateMapperAPIService) DeleteCertificateMapper(ctx context.Context, certificateMapperName string) ApiDeleteCertificateMapperRequest {
	return ApiDeleteCertificateMapperRequest{
		ApiService:            a,
		ctx:                   ctx,
		certificateMapperName: certificateMapperName,
	}
}

// Execute executes the request
func (a *CertificateMapperAPIService) DeleteCertificateMapperExecute(r ApiDeleteCertificateMapperRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateMapperAPIService.DeleteCertificateMapper")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificate-mappers/{certificate-mapper-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"certificate-mapper-name"+"}", url.PathEscape(parameterValueToString(r.certificateMapperName, "certificateMapperName")), -1)

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

type ApiGetCertificateMapperRequest struct {
	ctx                   context.Context
	ApiService            *CertificateMapperAPIService
	certificateMapperName string
}

func (r ApiGetCertificateMapperRequest) Execute() (*AddCertificateMapper200Response, *http.Response, error) {
	return r.ApiService.GetCertificateMapperExecute(r)
}

/*
GetCertificateMapper Returns a single Certificate Mapper

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param certificateMapperName Name of the Certificate Mapper
	@return ApiGetCertificateMapperRequest
*/
func (a *CertificateMapperAPIService) GetCertificateMapper(ctx context.Context, certificateMapperName string) ApiGetCertificateMapperRequest {
	return ApiGetCertificateMapperRequest{
		ApiService:            a,
		ctx:                   ctx,
		certificateMapperName: certificateMapperName,
	}
}

// Execute executes the request
//
//	@return AddCertificateMapper200Response
func (a *CertificateMapperAPIService) GetCertificateMapperExecute(r ApiGetCertificateMapperRequest) (*AddCertificateMapper200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCertificateMapper200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateMapperAPIService.GetCertificateMapper")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificate-mappers/{certificate-mapper-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"certificate-mapper-name"+"}", url.PathEscape(parameterValueToString(r.certificateMapperName, "certificateMapperName")), -1)

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

type ApiListCertificateMappersRequest struct {
	ctx        context.Context
	ApiService *CertificateMapperAPIService
	filter     *string
}

// SCIM filter
func (r ApiListCertificateMappersRequest) Filter(filter string) ApiListCertificateMappersRequest {
	r.filter = &filter
	return r
}

func (r ApiListCertificateMappersRequest) Execute() (*CertificateMapperListResponse, *http.Response, error) {
	return r.ApiService.ListCertificateMappersExecute(r)
}

/*
ListCertificateMappers Returns a list of all Certificate Mapper objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListCertificateMappersRequest
*/
func (a *CertificateMapperAPIService) ListCertificateMappers(ctx context.Context) ApiListCertificateMappersRequest {
	return ApiListCertificateMappersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertificateMapperListResponse
func (a *CertificateMapperAPIService) ListCertificateMappersExecute(r ApiListCertificateMappersRequest) (*CertificateMapperListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CertificateMapperListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateMapperAPIService.ListCertificateMappers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificate-mappers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
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

type ApiUpdateCertificateMapperRequest struct {
	ctx                   context.Context
	ApiService            *CertificateMapperAPIService
	certificateMapperName string
	updateRequest         *UpdateRequest
}

// Update an existing Certificate Mapper
func (r ApiUpdateCertificateMapperRequest) UpdateRequest(updateRequest UpdateRequest) ApiUpdateCertificateMapperRequest {
	r.updateRequest = &updateRequest
	return r
}

func (r ApiUpdateCertificateMapperRequest) Execute() (*AddCertificateMapper200Response, *http.Response, error) {
	return r.ApiService.UpdateCertificateMapperExecute(r)
}

/*
UpdateCertificateMapper Update an existing Certificate Mapper by name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param certificateMapperName Name of the Certificate Mapper
	@return ApiUpdateCertificateMapperRequest
*/
func (a *CertificateMapperAPIService) UpdateCertificateMapper(ctx context.Context, certificateMapperName string) ApiUpdateCertificateMapperRequest {
	return ApiUpdateCertificateMapperRequest{
		ApiService:            a,
		ctx:                   ctx,
		certificateMapperName: certificateMapperName,
	}
}

// Execute executes the request
//
//	@return AddCertificateMapper200Response
func (a *CertificateMapperAPIService) UpdateCertificateMapperExecute(r ApiUpdateCertificateMapperRequest) (*AddCertificateMapper200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCertificateMapper200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateMapperAPIService.UpdateCertificateMapper")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificate-mappers/{certificate-mapper-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"certificate-mapper-name"+"}", url.PathEscape(parameterValueToString(r.certificateMapperName, "certificateMapperName")), -1)

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
