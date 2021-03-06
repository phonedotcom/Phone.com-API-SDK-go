# \CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountCall**](CallsApi.md#CreateAccountCall) | **Post** /accounts/{account_id}/calls | Make a phone call


# **CreateAccountCall**
> CallFull CreateAccountCall($accountId, $data)

Make a phone call

Make a phone call. See Calls for more details and how to setup caller id's. Note: This API is for users with Account Owner scope access token. Users with Extension User scope token should invoke the Calls API with the following definition: POST https://api.phone.com/v4/accounts/:account_id/extensions/:extension_id/calls


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateCallParams**](CreateCallParams.md)| Call data | [optional] 

### Return type

[**CallFull**](CallFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

