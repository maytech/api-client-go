# \ProjectFolderApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectFolderAddUsersIdPost**](ProjectFolderApi.md#ProjectFolderAddUsersIdPost) | **Post** /project-folder/add-users/{id} | Add users to the project folder
[**ProjectFolderCreatePost**](ProjectFolderApi.md#ProjectFolderCreatePost) | **Post** /project-folder/create | Create a project folder
[**ProjectFolderDeleteIdGet**](ProjectFolderApi.md#ProjectFolderDeleteIdGet) | **Get** /project-folder/delete/{id} | Convert a project folder to a folder
[**ProjectFolderDeleteUsersPost**](ProjectFolderApi.md#ProjectFolderDeleteUsersPost) | **Post** /project-folder/delete-users/ | Remove project folder users
[**ProjectFolderEditUsersIdPost**](ProjectFolderApi.md#ProjectFolderEditUsersIdPost) | **Post** /project-folder/edit-users/{id} | Update users’ permissions of the project folder
[**ProjectFolderGet**](ProjectFolderApi.md#ProjectFolderGet) | **Get** /project-folder | List available project folders for a logged-in user
[**ProjectFolderMetadataIdGet**](ProjectFolderApi.md#ProjectFolderMetadataIdGet) | **Get** /project-folder/metadata/{id} | Get project folder metadata
[**ProjectFolderProjectUsersPost**](ProjectFolderApi.md#ProjectFolderProjectUsersPost) | **Post** /project-folder/project-users | List all project folders for given users
[**ProjectFolderSetUsersPost**](ProjectFolderApi.md#ProjectFolderSetUsersPost) | **Post** /project-folder/set-users | Add users to project folders.
[**ProjectFolderUsersIdGet**](ProjectFolderApi.md#ProjectFolderUsersIdGet) | **Get** /project-folder/users/{id} | List users of the project folder


# **ProjectFolderAddUsersIdPost**
> []UserPermissionResp ProjectFolderAddUsersIdPost(ctx, id, body)
Add users to the project folder

Add users to the specified project folder. 

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
Create a project folder

Create a new project folder for a user. 

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
Convert a project folder to a folder

Convert the project folder to regular folder keeping the same name and location. 

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

Delete users of the given project folder. 

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
Update users’ permissions of the project folder

Edit users’ permissions of the given project folder. 

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
List available project folders for a logged-in user

Retrieve a list of all project folders of a current logged-in user. 

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
Get project folder metadata

Retrieve the metadata of the project folder. 

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
List all project folders for given users

Get a list of all project folders for given users. 

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
Add users to project folders.

Add users to specified project folders. If the users exist in the given project folder, their permissions will be updated based on given parameters. 

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
List users of the project folder

Retrieve a list of users by given project folder ID. 

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

