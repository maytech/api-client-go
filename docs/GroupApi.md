# \GroupApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactGroupGet**](GroupApi.md#ContactGroupGet) | **Get** /contact/group | All contact groups
[**GroupGet**](GroupApi.md#GroupGet) | **Get** /group | List user groups
[**GroupMetadataIdGet**](GroupApi.md#GroupMetadataIdGet) | **Get** /group/metadata/{id} | Group metadata
[**UserGroupGet**](GroupApi.md#UserGroupGet) | **Get** /user/group | All user groups


# **ContactGroupGet**
> []ContactGroupRespItems ContactGroupGet(ctx, )
All contact groups

List all contact groups 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContactGroupRespItems**](ContactGroupRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupGet**
> []GroupResp GroupGet(ctx, )
List user groups

List user groups 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GroupResp**](GroupResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupMetadataIdGet**
> GroupResp GroupMetadataIdGet(ctx, id)
Group metadata

Get group metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a group | 

### Return type

[**GroupResp**](GroupResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGroupGet**
> []Group UserGroupGet(ctx, )
All user groups

List all user groups 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Group**](Group.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

