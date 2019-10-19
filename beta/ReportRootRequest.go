// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReportRootRequestBuilder is request builder for ReportRoot
type ReportRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns ReportRootRequest
func (b *ReportRootRequestBuilder) Request() *ReportRootRequest {
	return &ReportRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ReportRootRequest is request for ReportRoot
type ReportRootRequest struct{ BaseRequest }

// Do performs HTTP request for ReportRoot
func (r *ReportRootRequest) Do(method, path string, reqObj interface{}) (resObj *ReportRoot, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ReportRoot
func (r *ReportRootRequest) Get() (*ReportRoot, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ReportRoot
func (r *ReportRootRequest) Update(reqObj *ReportRoot) (*ReportRoot, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ReportRoot
func (r *ReportRootRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// ApplicationSignInDetailedSummary returns request builder for ApplicationSignInDetailedSummary collection
func (b *ReportRootRequestBuilder) ApplicationSignInDetailedSummary() *ReportRootApplicationSignInDetailedSummaryCollectionRequestBuilder {
	bb := &ReportRootApplicationSignInDetailedSummaryCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/applicationSignInDetailedSummary"
	return bb
}

// ReportRootApplicationSignInDetailedSummaryCollectionRequestBuilder is request builder for ApplicationSignInDetailedSummary collection
type ReportRootApplicationSignInDetailedSummaryCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ApplicationSignInDetailedSummary collection
func (b *ReportRootApplicationSignInDetailedSummaryCollectionRequestBuilder) Request() *ReportRootApplicationSignInDetailedSummaryCollectionRequest {
	return &ReportRootApplicationSignInDetailedSummaryCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ApplicationSignInDetailedSummary item
func (b *ReportRootApplicationSignInDetailedSummaryCollectionRequestBuilder) ID(id string) *ApplicationSignInDetailedSummaryRequestBuilder {
	bb := &ApplicationSignInDetailedSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ReportRootApplicationSignInDetailedSummaryCollectionRequest is request for ApplicationSignInDetailedSummary collection
type ReportRootApplicationSignInDetailedSummaryCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ApplicationSignInDetailedSummary collection
func (r *ReportRootApplicationSignInDetailedSummaryCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ApplicationSignInDetailedSummary, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ApplicationSignInDetailedSummary collection
func (r *ReportRootApplicationSignInDetailedSummaryCollectionRequest) Paging(method, path string, obj interface{}) ([]ApplicationSignInDetailedSummary, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ApplicationSignInDetailedSummary
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ApplicationSignInDetailedSummary
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ApplicationSignInDetailedSummary collection
func (r *ReportRootApplicationSignInDetailedSummaryCollectionRequest) Get() ([]ApplicationSignInDetailedSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ApplicationSignInDetailedSummary collection
func (r *ReportRootApplicationSignInDetailedSummaryCollectionRequest) Add(reqObj *ApplicationSignInDetailedSummary) (*ApplicationSignInDetailedSummary, error) {
	return r.Do("POST", "", reqObj)
}

// CredentialUserRegistrationDetails returns request builder for CredentialUserRegistrationDetails collection
func (b *ReportRootRequestBuilder) CredentialUserRegistrationDetails() *ReportRootCredentialUserRegistrationDetailsCollectionRequestBuilder {
	bb := &ReportRootCredentialUserRegistrationDetailsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/credentialUserRegistrationDetails"
	return bb
}

// ReportRootCredentialUserRegistrationDetailsCollectionRequestBuilder is request builder for CredentialUserRegistrationDetails collection
type ReportRootCredentialUserRegistrationDetailsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CredentialUserRegistrationDetails collection
func (b *ReportRootCredentialUserRegistrationDetailsCollectionRequestBuilder) Request() *ReportRootCredentialUserRegistrationDetailsCollectionRequest {
	return &ReportRootCredentialUserRegistrationDetailsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CredentialUserRegistrationDetails item
func (b *ReportRootCredentialUserRegistrationDetailsCollectionRequestBuilder) ID(id string) *CredentialUserRegistrationDetailsRequestBuilder {
	bb := &CredentialUserRegistrationDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ReportRootCredentialUserRegistrationDetailsCollectionRequest is request for CredentialUserRegistrationDetails collection
type ReportRootCredentialUserRegistrationDetailsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for CredentialUserRegistrationDetails collection
func (r *ReportRootCredentialUserRegistrationDetailsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *CredentialUserRegistrationDetails, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for CredentialUserRegistrationDetails collection
func (r *ReportRootCredentialUserRegistrationDetailsCollectionRequest) Paging(method, path string, obj interface{}) ([]CredentialUserRegistrationDetails, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []CredentialUserRegistrationDetails
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []CredentialUserRegistrationDetails
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for CredentialUserRegistrationDetails collection
func (r *ReportRootCredentialUserRegistrationDetailsCollectionRequest) Get() ([]CredentialUserRegistrationDetails, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for CredentialUserRegistrationDetails collection
func (r *ReportRootCredentialUserRegistrationDetailsCollectionRequest) Add(reqObj *CredentialUserRegistrationDetails) (*CredentialUserRegistrationDetails, error) {
	return r.Do("POST", "", reqObj)
}

// UserCredentialUsageDetails returns request builder for UserCredentialUsageDetails collection
func (b *ReportRootRequestBuilder) UserCredentialUsageDetails() *ReportRootUserCredentialUsageDetailsCollectionRequestBuilder {
	bb := &ReportRootUserCredentialUsageDetailsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userCredentialUsageDetails"
	return bb
}

// ReportRootUserCredentialUsageDetailsCollectionRequestBuilder is request builder for UserCredentialUsageDetails collection
type ReportRootUserCredentialUsageDetailsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserCredentialUsageDetails collection
func (b *ReportRootUserCredentialUsageDetailsCollectionRequestBuilder) Request() *ReportRootUserCredentialUsageDetailsCollectionRequest {
	return &ReportRootUserCredentialUsageDetailsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserCredentialUsageDetails item
func (b *ReportRootUserCredentialUsageDetailsCollectionRequestBuilder) ID(id string) *UserCredentialUsageDetailsRequestBuilder {
	bb := &UserCredentialUsageDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ReportRootUserCredentialUsageDetailsCollectionRequest is request for UserCredentialUsageDetails collection
type ReportRootUserCredentialUsageDetailsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for UserCredentialUsageDetails collection
func (r *ReportRootUserCredentialUsageDetailsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *UserCredentialUsageDetails, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for UserCredentialUsageDetails collection
func (r *ReportRootUserCredentialUsageDetailsCollectionRequest) Paging(method, path string, obj interface{}) ([]UserCredentialUsageDetails, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UserCredentialUsageDetails
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []UserCredentialUsageDetails
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for UserCredentialUsageDetails collection
func (r *ReportRootUserCredentialUsageDetailsCollectionRequest) Get() ([]UserCredentialUsageDetails, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for UserCredentialUsageDetails collection
func (r *ReportRootUserCredentialUsageDetailsCollectionRequest) Add(reqObj *UserCredentialUsageDetails) (*UserCredentialUsageDetails, error) {
	return r.Do("POST", "", reqObj)
}