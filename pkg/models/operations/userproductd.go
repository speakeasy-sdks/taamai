// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"net/http"
)

var UserProductdServerList = []string{
	"http://127.0.0.1:8000/api",
}

type UserProductdRequest struct {
	UserID int `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *UserProductdRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// UserProductdUserProductd - OK
type UserProductdUserProductd struct {
	Data    []shared.Data13 `json:"data"`
	Message string          `json:"message"`
	Status  string          `json:"status"`
}

func (o *UserProductdUserProductd) GetData() []shared.Data13 {
	if o == nil {
		return []shared.Data13{}
	}
	return o.Data
}

func (o *UserProductdUserProductd) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *UserProductdUserProductd) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type UserProductdResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UserProductd *UserProductdUserProductd
}

func (o *UserProductdResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UserProductdResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *UserProductdResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UserProductdResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UserProductdResponse) GetUserProductd() *UserProductdUserProductd {
	if o == nil {
		return nil
	}
	return o.UserProductd
}