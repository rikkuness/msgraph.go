// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsVarARequestBuilder struct{ BaseRequestBuilder }

// VarA action undocumented
func (b *WorkbookFunctionsRequestBuilder) VarA(reqObj *WorkbookFunctionsVarARequestParameter) *WorkbookFunctionsVarARequestBuilder {
	bb := &WorkbookFunctionsVarARequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/varA"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsVarARequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsVarARequestBuilder) Request() *WorkbookFunctionsVarARequest {
	return &WorkbookFunctionsVarARequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsVarARequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsVarPARequestBuilder struct{ BaseRequestBuilder }

// VarPA action undocumented
func (b *WorkbookFunctionsRequestBuilder) VarPA(reqObj *WorkbookFunctionsVarPARequestParameter) *WorkbookFunctionsVarPARequestBuilder {
	bb := &WorkbookFunctionsVarPARequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/varPA"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsVarPARequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsVarPARequestBuilder) Request() *WorkbookFunctionsVarPARequest {
	return &WorkbookFunctionsVarPARequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsVarPARequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}