// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"net/http"
)

var AllworkbooksServerList = []string{
	"https://taam.ai/api",
}

type AllworkbooksRequest struct {
	UserID int `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *AllworkbooksRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// AllworkbooksAllworkbookslive - OK
type AllworkbooksAllworkbookslive struct {
	Code    int             `json:"code"`
	Data    []shared.Datum1 `json:"data"`
	Mesaage string          `json:"mesaage"`
	Status  string          `json:"status"`
}

func (o *AllworkbooksAllworkbookslive) GetCode() int {
	if o == nil {
		return 0
	}
	return o.Code
}

func (o *AllworkbooksAllworkbookslive) GetData() []shared.Datum1 {
	if o == nil {
		return []shared.Datum1{}
	}
	return o.Data
}

func (o *AllworkbooksAllworkbookslive) GetMesaage() string {
	if o == nil {
		return ""
	}
	return o.Mesaage
}

func (o *AllworkbooksAllworkbookslive) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type AllworkbooksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Allworkbookslive *AllworkbooksAllworkbookslive
}

func (o *AllworkbooksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AllworkbooksResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *AllworkbooksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AllworkbooksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AllworkbooksResponse) GetAllworkbookslive() *AllworkbooksAllworkbookslive {
	if o == nil {
		return nil
	}
	return o.Allworkbookslive
}
