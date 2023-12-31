// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"net/http"
)

var TrashedworkspacesServerList = []string{
	"http://127.0.0.1:8000/api",
}

type TrashedworkspacesRequest struct {
	UserID int `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *TrashedworkspacesRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// TrashedworkspacesTrashedworkspaceslive - OK
type TrashedworkspacesTrashedworkspaceslive struct {
	Data    []shared.Data2 `json:"data"`
	Message string         `json:"message"`
	Status  string         `json:"status"`
}

func (o *TrashedworkspacesTrashedworkspaceslive) GetData() []shared.Data2 {
	if o == nil {
		return []shared.Data2{}
	}
	return o.Data
}

func (o *TrashedworkspacesTrashedworkspaceslive) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *TrashedworkspacesTrashedworkspaceslive) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type TrashedworkspacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Trashedworkspaceslive *TrashedworkspacesTrashedworkspaceslive
}

func (o *TrashedworkspacesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TrashedworkspacesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *TrashedworkspacesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TrashedworkspacesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TrashedworkspacesResponse) GetTrashedworkspaceslive() *TrashedworkspacesTrashedworkspaceslive {
	if o == nil {
		return nil
	}
	return o.Trashedworkspaceslive
}
