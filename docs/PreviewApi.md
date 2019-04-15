# \PreviewApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilePreviewIdGet**](PreviewApi.md#FilePreviewIdGet) | **Get** /file/preview/{id} | Get a file preview
[**PreviewIdGet**](PreviewApi.md#PreviewIdGet) | **Get** /preview/{id} | Get binary preview data


# **FilePreviewIdGet**
> FilePreviewResp FilePreviewIdGet(ctx, id)
Get a file preview

Retrieve a file preview by the given ID of the file. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a file | 

### Return type

[**FilePreviewResp**](FilePreviewResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreviewIdGet**
> PreviewIdGet(ctx, id)
Get binary preview data

Get a preview with the response content type based on the file type e.g. image/jpeg for images. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| File ID | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: image/jpeg, video/mp4, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

