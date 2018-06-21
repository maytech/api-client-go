# \ServiceApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceGet**](ServiceApi.md#ServiceGet) | **Get** /service | List services
[**ServiceMetadataIdGet**](ServiceApi.md#ServiceMetadataIdGet) | **Get** /service/metadata/{id} | Service metadata


# **ServiceGet**
> []ServiceResp ServiceGet(ctx, )
List services

List services 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ServiceResp**](ServiceResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceMetadataIdGet**
> ServiceResp ServiceMetadataIdGet(ctx, id)
Service metadata

Get service metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| Service ID | 

### Return type

[**ServiceResp**](ServiceResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

