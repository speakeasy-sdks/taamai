// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var AddandremovefromfavDocumentServerList = []string{
	"http://127.0.0.1:8000/api",
}

type AddandremovefromfavDocumentRequest struct {
	ID     int    `queryParam:"style=form,explode=true,name=id"`
	Type   string `queryParam:"style=form,explode=true,name=type"`
	UserID int    `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *AddandremovefromfavDocumentRequest) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *AddandremovefromfavDocumentRequest) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *AddandremovefromfavDocumentRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// AddandremovefromfavDocumentAddandremovefromfavDocument - OK
type AddandremovefromfavDocumentAddandremovefromfavDocument struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (o *AddandremovefromfavDocumentAddandremovefromfavDocument) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *AddandremovefromfavDocumentAddandremovefromfavDocument) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type AddandremovefromfavDocumentResponse struct {
	// OK
	AddandremovefromfavDocument *AddandremovefromfavDocumentAddandremovefromfavDocument
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddandremovefromfavDocumentResponse) GetAddandremovefromfavDocument() *AddandremovefromfavDocumentAddandremovefromfavDocument {
	if o == nil {
		return nil
	}
	return o.AddandremovefromfavDocument
}

func (o *AddandremovefromfavDocumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddandremovefromfavDocumentResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *AddandremovefromfavDocumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddandremovefromfavDocumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
