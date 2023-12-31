// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var PrompttemplatelikeorremovefromlikeServerList = []string{
	"http://127.0.0.1:8000/api",
}

type PrompttemplatelikeorremovefromlikeRequest struct {
	TemplateID int `queryParam:"style=form,explode=true,name=template_id"`
	UserID     int `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *PrompttemplatelikeorremovefromlikeRequest) GetTemplateID() int {
	if o == nil {
		return 0
	}
	return o.TemplateID
}

func (o *PrompttemplatelikeorremovefromlikeRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// PrompttemplatelikeorremovefromlikePrompttemplatelikeorremovefromlike - OK
type PrompttemplatelikeorremovefromlikePrompttemplatelikeorremovefromlike struct {
	Count   int    `json:"count"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (o *PrompttemplatelikeorremovefromlikePrompttemplatelikeorremovefromlike) GetCount() int {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *PrompttemplatelikeorremovefromlikePrompttemplatelikeorremovefromlike) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *PrompttemplatelikeorremovefromlikePrompttemplatelikeorremovefromlike) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type PrompttemplatelikeorremovefromlikeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Prompttemplatelikeorremovefromlike *PrompttemplatelikeorremovefromlikePrompttemplatelikeorremovefromlike
}

func (o *PrompttemplatelikeorremovefromlikeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PrompttemplatelikeorremovefromlikeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PrompttemplatelikeorremovefromlikeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PrompttemplatelikeorremovefromlikeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PrompttemplatelikeorremovefromlikeResponse) GetPrompttemplatelikeorremovefromlike() *PrompttemplatelikeorremovefromlikePrompttemplatelikeorremovefromlike {
	if o == nil {
		return nil
	}
	return o.Prompttemplatelikeorremovefromlike
}
