// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"net/http"
)

var TrashedPromptTemplatesServerList = []string{
	"http://127.0.0.1:8000/api",
}

type TrashedPromptTemplatesRequest struct {
	UserID int `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *TrashedPromptTemplatesRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// TrashedPromptTemplatesTrashedPromptTemplates - OK
type TrashedPromptTemplatesTrashedPromptTemplates struct {
	Data    []shared.Data16 `json:"data"`
	Message string          `json:"message"`
	Status  string          `json:"status"`
}

func (o *TrashedPromptTemplatesTrashedPromptTemplates) GetData() []shared.Data16 {
	if o == nil {
		return []shared.Data16{}
	}
	return o.Data
}

func (o *TrashedPromptTemplatesTrashedPromptTemplates) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *TrashedPromptTemplatesTrashedPromptTemplates) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type TrashedPromptTemplatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TrashedPromptTemplates *TrashedPromptTemplatesTrashedPromptTemplates
}

func (o *TrashedPromptTemplatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TrashedPromptTemplatesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *TrashedPromptTemplatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TrashedPromptTemplatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TrashedPromptTemplatesResponse) GetTrashedPromptTemplates() *TrashedPromptTemplatesTrashedPromptTemplates {
	if o == nil {
		return nil
	}
	return o.TrashedPromptTemplates
}