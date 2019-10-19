// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ManagementConditionStatementRequestBuilder is request builder for ManagementConditionStatement
type ManagementConditionStatementRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagementConditionStatementRequest
func (b *ManagementConditionStatementRequestBuilder) Request() *ManagementConditionStatementRequest {
	return &ManagementConditionStatementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagementConditionStatementRequest is request for ManagementConditionStatement
type ManagementConditionStatementRequest struct{ BaseRequest }

// Do performs HTTP request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Do(method, path string, reqObj interface{}) (resObj *ManagementConditionStatement, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Get() (*ManagementConditionStatement, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Update(reqObj *ManagementConditionStatement) (*ManagementConditionStatement, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// ManagementConditions returns request builder for ManagementCondition collection
func (b *ManagementConditionStatementRequestBuilder) ManagementConditions() *ManagementConditionStatementManagementConditionsCollectionRequestBuilder {
	bb := &ManagementConditionStatementManagementConditionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managementConditions"
	return bb
}

// ManagementConditionStatementManagementConditionsCollectionRequestBuilder is request builder for ManagementCondition collection
type ManagementConditionStatementManagementConditionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagementCondition collection
func (b *ManagementConditionStatementManagementConditionsCollectionRequestBuilder) Request() *ManagementConditionStatementManagementConditionsCollectionRequest {
	return &ManagementConditionStatementManagementConditionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagementCondition item
func (b *ManagementConditionStatementManagementConditionsCollectionRequestBuilder) ID(id string) *ManagementConditionRequestBuilder {
	bb := &ManagementConditionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagementConditionStatementManagementConditionsCollectionRequest is request for ManagementCondition collection
type ManagementConditionStatementManagementConditionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ManagementCondition, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagementCondition, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagementCondition
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ManagementCondition
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

// Get performs GET request for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Get() ([]ManagementCondition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Add(reqObj *ManagementCondition) (*ManagementCondition, error) {
	return r.Do("POST", "", reqObj)
}