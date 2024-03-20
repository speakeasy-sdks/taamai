// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var ParmenentdeleteProductServerList = []string{
	"http://127.0.0.1:8000/api",
}

type ParmenentdeleteProductRequest struct {
	ProductID int `queryParam:"style=form,explode=true,name=product_id"`
	UserID    int `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *ParmenentdeleteProductRequest) GetProductID() int {
	if o == nil {
		return 0
	}
	return o.ProductID
}

func (o *ParmenentdeleteProductRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// ParmenentdeleteProductNewRequest1 - OK
type ParmenentdeleteProductNewRequest1 struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (o *ParmenentdeleteProductNewRequest1) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *ParmenentdeleteProductNewRequest1) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type ParmenentdeleteProductResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	NewRequest1 *ParmenentdeleteProductNewRequest1
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ParmenentdeleteProductResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ParmenentdeleteProductResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ParmenentdeleteProductResponse) GetNewRequest1() *ParmenentdeleteProductNewRequest1 {
	if o == nil {
		return nil
	}
	return o.NewRequest1
}

func (o *ParmenentdeleteProductResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ParmenentdeleteProductResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
