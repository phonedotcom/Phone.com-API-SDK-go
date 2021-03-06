/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type CalleridsApi struct {
	Configuration *Configuration
}

func NewCalleridsApi() *CalleridsApi {
	configuration := NewConfiguration()
	return &CalleridsApi{
		Configuration: configuration,
	}
}

func NewCalleridsApiWithBasePath(basePath string) *CalleridsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &CalleridsApi{
		Configuration: configuration,
	}
}

/**
 * Show the Caller ID options a given extension can use.
 * Show the Caller ID options a given extension can use. See Intro to Caller IDs for more on the properties.
 *
 * @param accountId Account ID
 * @param extensionId Extension ID
 * @param filtersNumber Number filter
 * @param filtersName Name filter
 * @param sortNumber Number sorting
 * @param sortName Name sorting
 * @param limit Max results
 * @param offset Results to skip
 * @param fields Field set
 * @return *ListCallerIds
 */
func (a CalleridsApi) GetCallerIds(accountId int32, extensionId int32, filtersNumber []string, filtersName []string, sortNumber string, sortName string, limit int32, offset int32, fields string) (*ListCallerIds, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/accounts/{account_id}/extensions/{extension_id}/caller-ids"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extension_id"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(apiKey)' required
	// set key with prefix in header
	localVarHeaderParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var filtersNumberCollectionFormat = "multi"
	localVarQueryParams.Add("filters[number]", a.Configuration.APIClient.ParameterToString(filtersNumber, filtersNumberCollectionFormat))

	var filtersNameCollectionFormat = "multi"
	localVarQueryParams.Add("filters[name]", a.Configuration.APIClient.ParameterToString(filtersName, filtersNameCollectionFormat))

	localVarQueryParams.Add("sort[number]", a.Configuration.APIClient.ParameterToString(sortNumber, ""))
	localVarQueryParams.Add("sort[name]", a.Configuration.APIClient.ParameterToString(sortName, ""))
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
	localVarQueryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))

	clearEmptyParams(localVarQueryParams)

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ListCallerIds)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetCallerIds", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

