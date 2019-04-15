# \GroupApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactGroupGet**](GroupApi.md#ContactGroupGet) | **Get** /contact/group | List available contact groups.
[**GroupGet**](GroupApi.md#GroupGet) | **Get** /group | List available user groups
[**GroupMetadataIdGet**](GroupApi.md#GroupMetadataIdGet) | **Get** /group/metadata/{id} | Get group metadata
[**UserGroupGet**](GroupApi.md#UserGroupGet) | **Get** /user/group | List all user groups


# **ContactGroupGet**
> []ContactGroupRespItems ContactGroupGet(ctx, )
List available contact groups.

Get a list of available contact groups. 

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
List available user groups

Get a list of available contact groups. 

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
Get group metadata

Retrieve information about the group by the specified group ID. 

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
List all user groups

Retrieve a list of all user groups that can be set for users. 

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

