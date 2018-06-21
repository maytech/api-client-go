# \ProjectFolderApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectFolderAddUsersIdPost**](ProjectFolderApi.md#ProjectFolderAddUsersIdPost) | **Post** /project-folder/add-users/{id} | Add project folder users
[**ProjectFolderCreatePost**](ProjectFolderApi.md#ProjectFolderCreatePost) | **Post** /project-folder/create | Create project folder
[**ProjectFolderDeleteIdGet**](ProjectFolderApi.md#ProjectFolderDeleteIdGet) | **Get** /project-folder/delete/{id} | Remove project folder
[**ProjectFolderDeleteUsersPost**](ProjectFolderApi.md#ProjectFolderDeleteUsersPost) | **Post** /project-folder/delete-users/ | Remove project folder users
[**ProjectFolderEditUsersIdPost**](ProjectFolderApi.md#ProjectFolderEditUsersIdPost) | **Post** /project-folder/edit-users/{id} | Update project folder users
[**ProjectFolderGet**](ProjectFolderApi.md#ProjectFolderGet) | **Get** /project-folder | List of valid project folder for current user
[**ProjectFolderMetadataIdGet**](ProjectFolderApi.md#ProjectFolderMetadataIdGet) | **Get** /project-folder/metadata/{id} | Project folder metadata
[**ProjectFolderProjectUsersPost**](ProjectFolderApi.md#ProjectFolderProjectUsersPost) | **Post** /project-folder/project-users | List of project folders for users
[**ProjectFolderSetUsersPost**](ProjectFolderApi.md#ProjectFolderSetUsersPost) | **Post** /project-folder/set-users | Add users to project folders. Replace exists permissions if users exists
[**ProjectFolderUsersIdGet**](ProjectFolderApi.md#ProjectFolderUsersIdGet) | **Get** /project-folder/users/{id} | List project folder users


# **ProjectFolderAddUsersIdPost**
> []UserPermissionResp ProjectFolderAddUsersIdPost(ctx, id, body)
Add project folder users

Add project folder users 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)|  | 
  **body** | [**PfaddUsersReq**](PfaddUsersReq.md)|  | 

### Return type

[**[]UserPermissionResp**](UserPermissionResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderCreatePost**
> PfCreateResp ProjectFolderCreatePost(ctx, body)
Create project folder

Create new project folder 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**PfcreateReq**](PfcreateReq.md)|  | 

### Return type

[**PfCreateResp**](PfCreateResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderDeleteIdGet**
> IdResp ProjectFolderDeleteIdGet(ctx, id)
Remove project folder

Remove project folder 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**IdResp**](IdResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderDeleteUsersPost**
> []PfDeleteUsersRespItems ProjectFolderDeleteUsersPost(ctx, body)
Remove project folder users

Remove project folder users 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**PfdeleteUsersReq**](PfdeleteUsersReq.md)|  | 

### Return type

[**[]PfDeleteUsersRespItems**](PfDeleteUsersRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderEditUsersIdPost**
> []UserPermissionResp ProjectFolderEditUsersIdPost(ctx, id, body)
Update project folder users

Update project folder users 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)|  | 
  **body** | [**PfeditUsersReq**](PfeditUsersReq.md)|  | 

### Return type

[**[]UserPermissionResp**](UserPermissionResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderGet**
> []ProjectFoldersListRespItem ProjectFolderGet(ctx, )
List of valid project folder for current user

Get list of project folders 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProjectFoldersListRespItem**](ProjectFoldersListRespItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderMetadataIdGet**
> PfMetadataResp ProjectFolderMetadataIdGet(ctx, id)
Project folder metadata

Project folder metadata 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| project folder ID | 

### Return type

[**PfMetadataResp**](PfMetadataResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderProjectUsersPost**
> []PfUsersListRespItems ProjectFolderProjectUsersPost(ctx, body)
List of project folders for users

List of project folders for users 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**PfUsersListReq**](PfUsersListReq.md)|  | 

### Return type

[**[]PfUsersListRespItems**](PfUsersListRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderSetUsersPost**
> []PfSetUsersRespItems ProjectFolderSetUsersPost(ctx, body)
Add users to project folders. Replace exists permissions if users exists

Set users to project folders 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**PfSetUsersReq**](PfSetUsersReq.md)|  | 

### Return type

[**[]PfSetUsersRespItems**](PfSetUsersRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectFolderUsersIdGet**
> []UserPermissionResp ProjectFolderUsersIdGet(ctx, id)
List project folder users

List project folder users 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| project folder ID | 

### Return type

[**[]UserPermissionResp**](UserPermissionResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

