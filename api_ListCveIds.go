package cveservices_go_sdk

import (
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
ListCveIds
Lists all CVE Ids associated to an entity. No roles are needed to access the endpoint.
Expected Behavior:
Secretariat - Can see CVE-IDs owned by any Organization.
Admin User - Can only see CVE-IDs owned by the Organization it belongs to.
Regular User - Can only see CVE-IDs owned by the Organization it belongs to.
*/
func (a *APIClient) ListCveIds(localVarOptionals *types.ListCveIdsOpts) (types.ListCveIdsResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue types.ListCveIdsResponse
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/cve-id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarQueryParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CveIdYear.IsSet() {
		localVarQueryParams.Add("cve_id_year", parameterToString(localVarOptionals.CveIdYear.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimeReservedLt.IsSet() {
		localVarQueryParams.Add("time_reserved.lt", parameterToString(localVarOptionals.TimeReservedLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimeReservedGt.IsSet() {
		localVarQueryParams.Add("time_reserved.gt", parameterToString(localVarOptionals.TimeReservedGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimeModifiedLt.IsSet() {
		localVarQueryParams.Add("time_modified.lt", parameterToString(localVarOptionals.TimeModifiedLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimeModifiedGt.IsSet() {
		localVarQueryParams.Add("time_modified.gt", parameterToString(localVarOptionals.TimeModifiedGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	} else if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v types.ListCveIdsResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		} else if localVarHttpResponse.StatusCode >= 400 {
			var v types.ErrorResponse
			err = a.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			if v.Message != "" {
				newErr.error = v.Message
			} else {
				newErr.error = v.Message
			}
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
