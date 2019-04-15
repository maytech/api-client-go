# \UploadApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileModifyPost**](UploadApi.md#FileModifyPost) | **Post** /file/modify | Get file modification link
[**SettingsUploadLogoLinkGet**](UploadApi.md#SettingsUploadLogoLinkGet) | **Get** /settings/upload-logo-link | Get a new logo upload link
[**UploadFinalizeIdGet**](UploadApi.md#UploadFinalizeIdGet) | **Get** /upload/finalize/{id} | Finalize chunked file upload
[**UploadLinkPost**](UploadApi.md#UploadLinkPost) | **Post** /upload/link | Get file upload link


# **FileModifyPost**
> FileModifyResp FileModifyPost(ctx, body)
Get file modification link

Get file modification link 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FileModifyReq**](FileModifyReq.md)|  | 

### Return type

[**FileModifyResp**](FileModifyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsUploadLogoLinkGet**
> SettingsUploadLogoLinkResp SettingsUploadLogoLinkGet(ctx, )
Get a new logo upload link

Get a unique key for uploading a new logo 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsUploadLogoLinkResp**](SettingsUploadLogoLinkResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFinalizeIdGet**
> UploadFinalizeResp UploadFinalizeIdGet(ctx, id, optional)
Finalize chunked file upload

Complete the chunked upload of a file. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Upload key | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| Upload key | 
 **mtime** | **float32**| File modification timestamp | 

### Return type

[**UploadFinalizeResp**](UploadFinalizeResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadLinkPost**
> FileModifyResp UploadLinkPost(ctx, body)
Get file upload link

Retrieve a link for uploading a file. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**UploadLinkReq**](UploadLinkReq.md)|  | 

### Return type

[**FileModifyResp**](FileModifyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

