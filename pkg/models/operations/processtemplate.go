// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var ProcessTemplateServerList = []string{
	"https://taam.ai/api",
}

type ProcessTemplateRequest struct {
	ContentID   int     `queryParam:"style=form,explode=true,name=content_id"`
	MaxResults  int     `queryParam:"style=form,explode=true,name=max_results"`
	MaxWords    int     `queryParam:"style=form,explode=true,name=max_words"`
	Temperature float64 `queryParam:"style=form,explode=true,name=temperature"`
	UserID      int     `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *ProcessTemplateRequest) GetContentID() int {
	if o == nil {
		return 0
	}
	return o.ContentID
}

func (o *ProcessTemplateRequest) GetMaxResults() int {
	if o == nil {
		return 0
	}
	return o.MaxResults
}

func (o *ProcessTemplateRequest) GetMaxWords() int {
	if o == nil {
		return 0
	}
	return o.MaxWords
}

func (o *ProcessTemplateRequest) GetTemperature() float64 {
	if o == nil {
		return 0.0
	}
	return o.Temperature
}

func (o *ProcessTemplateRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// ProcessTemplateProcessTemplate - OK
type ProcessTemplateProcessTemplate struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}

func (o *ProcessTemplateProcessTemplate) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}

func (o *ProcessTemplateProcessTemplate) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type ProcessTemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// OK
	ProcessTemplate *ProcessTemplateProcessTemplate
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ProcessTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ProcessTemplateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ProcessTemplateResponse) GetProcessTemplate() *ProcessTemplateProcessTemplate {
	if o == nil {
		return nil
	}
	return o.ProcessTemplate
}

func (o *ProcessTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ProcessTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
