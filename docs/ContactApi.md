# \ContactApi

All URIs are relative to *https://api.quatrix.it/api/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactCreatePost**](ContactApi.md#ContactCreatePost) | **Post** /contact/create | Create a new contact
[**ContactDeletePost**](ContactApi.md#ContactDeletePost) | **Post** /contact/delete | Delete a contact
[**ContactEditIdPost**](ContactApi.md#ContactEditIdPost) | **Post** /contact/edit/{id} | Edit contact metadata
[**ContactGet**](ContactApi.md#ContactGet) | **Get** /contact | List user contacts
[**ContactGroupGet**](ContactApi.md#ContactGroupGet) | **Get** /contact/group | List available contact groups.
[**ContactMetadataIdGet**](ContactApi.md#ContactMetadataIdGet) | **Get** /contact/metadata/{id} | Get contact metadata
[**ContactPgpKeyIdGet**](ContactApi.md#ContactPgpKeyIdGet) | **Get** /contact/pgp-key/{id} | Get contact&#39;s PGP key
[**ContactUpgradeIdGet**](ContactApi.md#ContactUpgradeIdGet) | **Get** /contact/upgrade/{id} | Upgrade a contact


# **ContactCreatePost**
> ContactResp ContactCreatePost(ctx, body)
Create a new contact

Add a new contact to the account. The contact will be assigned a unique ID that can be used to look up the contact inside of Quatrix later. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ContactCreateReq**](ContactCreateReq.md)|  | 

### Return type

[**ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactDeletePost**
> []ContactDeleteRespItems ContactDeletePost(ctx, body)
Delete a contact

Delete an existing contact from the account. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdsReq**](IdsReq.md)| IDs of a contacts | 

### Return type

[**[]ContactDeleteRespItems**](ContactDeleteRespItems.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactEditIdPost**
> ContactResp ContactEditIdPost(ctx, id, optional)
Edit contact metadata

Use to edit the existing contact details. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)| ID of a contact | 
 **body** | [**ContactEditResp**](ContactEditResp.md)|  | 

### Return type

[**ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactGet**
> []ContactResp ContactGet(ctx, )
List user contacts

Retrieve the list of contacts (personal and site) available in the current account. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **ContactMetadataIdGet**
> ContactResp ContactMetadataIdGet(ctx, id)
Get contact metadata

Retrieve contact details containing the ID, name, email, status, time it was created, PGP key details, assigned group and permissions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 

### Return type

[**ContactResp**](ContactResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactPgpKeyIdGet**
> PgpKeyResp ContactPgpKeyIdGet(ctx, id)
Get contact's PGP key

Use to retrieve PGP key generated by the contact. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 

### Return type

[**PgpKeyResp**](PgpKeyResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactUpgradeIdGet**
> UserResp ContactUpgradeIdGet(ctx, id)
Upgrade a contact

Convert an existing contact to the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)| ID of a contact | 

### Return type

[**UserResp**](UserResp.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

