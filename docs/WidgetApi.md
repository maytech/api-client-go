# \WidgetApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WidgetFinalizeUploadIdGet**](WidgetApi.md#WidgetFinalizeUploadIdGet) | **Get** /widget/finalize-upload/{id} | Finalize chunked upload of the widget
[**WidgetMetadataIdGet**](WidgetApi.md#WidgetMetadataIdGet) | **Get** /widget/metadata/{id} | Get all widget metadata
[**WidgetUploadLinkIdPost**](WidgetApi.md#WidgetUploadLinkIdPost) | **Post** /widget/upload-link/{id} | Get widget upload link


# **WidgetFinalizeUploadIdGet**
> WidgetFinalizeUploadResp WidgetFinalizeUploadIdGet(ctx, id)
Finalize chunked upload of the widget

Complete the chunked upload of the widget. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a widget | 

### Return type

[**WidgetFinalizeUploadResp**](WidgetFinalizeUploadResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WidgetMetadataIdGet**
> IdResp WidgetMetadataIdGet(ctx, id)
Get all widget metadata

Retrieve available metadata of the widget. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a widget | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WidgetUploadLinkIdPost**
> WidgetUploadLinkResp WidgetUploadLinkIdPost(ctx, id, optional)
Get widget upload link

Retrieve a link for uploading the widget. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a widget | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| ID of a widget | 
 **body** | [**WidgetUploadLinkReq**](WidgetUploadLinkReq.md)|  | 

### Return type

[**WidgetUploadLinkResp**](WidgetUploadLinkResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

