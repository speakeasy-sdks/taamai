// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"net/http"
)

var WorkbookvoiceoversServerList = []string{
	"https://taam.ai/api",
}

type WorkbookvoiceoversRequest struct {
	FolderID   int    `queryParam:"style=form,explode=true,name=folder_id"`
	Type       string `queryParam:"style=form,explode=true,name=type"`
	UserID     int    `queryParam:"style=form,explode=true,name=user_id"`
	WorkbookID int    `queryParam:"style=form,explode=true,name=workbook_id"`
}

func (o *WorkbookvoiceoversRequest) GetFolderID() int {
	if o == nil {
		return 0
	}
	return o.FolderID
}

func (o *WorkbookvoiceoversRequest) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *WorkbookvoiceoversRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *WorkbookvoiceoversRequest) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}

// WorkbookvoiceoversWorkbookvoiceovers - OK
type WorkbookvoiceoversWorkbookvoiceovers struct {
	Count   int            `json:"count"`
	Data    []shared.Data6 `json:"data"`
	Message string         `json:"message"`
	Status  string         `json:"status"`
}

func (o *WorkbookvoiceoversWorkbookvoiceovers) GetCount() int {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *WorkbookvoiceoversWorkbookvoiceovers) GetData() []shared.Data6 {
	if o == nil {
		return []shared.Data6{}
	}
	return o.Data
}

func (o *WorkbookvoiceoversWorkbookvoiceovers) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *WorkbookvoiceoversWorkbookvoiceovers) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type WorkbookvoiceoversResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Workbookvoiceovers *WorkbookvoiceoversWorkbookvoiceovers
}

func (o *WorkbookvoiceoversResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WorkbookvoiceoversResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *WorkbookvoiceoversResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WorkbookvoiceoversResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WorkbookvoiceoversResponse) GetWorkbookvoiceovers() *WorkbookvoiceoversWorkbookvoiceovers {
	if o == nil {
		return nil
	}
	return o.Workbookvoiceovers
}