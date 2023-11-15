// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var JoinworkbookServerList = []string{
	"http://127.0.0.1:8000/api",
}

type JoinworkbookRequest struct {
	UserID     int `queryParam:"style=form,explode=true,name=user_id"`
	WorkbookID int `queryParam:"style=form,explode=true,name=workbook_id"`
}

func (o *JoinworkbookRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *JoinworkbookRequest) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}

// JoinworkbookJoinworkbook - OK
type JoinworkbookJoinworkbook struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (o *JoinworkbookJoinworkbook) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *JoinworkbookJoinworkbook) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type JoinworkbookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	Joinworkbook *JoinworkbookJoinworkbook
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *JoinworkbookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *JoinworkbookResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *JoinworkbookResponse) GetJoinworkbook() *JoinworkbookJoinworkbook {
	if o == nil {
		return nil
	}
	return o.Joinworkbook
}

func (o *JoinworkbookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *JoinworkbookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
