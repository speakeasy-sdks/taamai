// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var CreatePromptTemplateServerList = []string{
	"http://127.0.0.1:8000/api",
}

type CreatePromptTemplateRequest struct {
	Activate      int    `queryParam:"style=form,explode=true,name=activate"`
	Category      string `queryParam:"style=form,explode=true,name=category"`
	Code0         string `queryParam:"style=form,explode=true,name=code[0]"`
	Edit          int    `queryParam:"style=form,explode=true,name=edit"`
	InputField0   string `queryParam:"style=form,explode=true,name=input_field[0]"`
	Language      string `queryParam:"style=form,explode=true,name=language"`
	Name          string `queryParam:"style=form,explode=true,name=name"`
	Names0        string `queryParam:"style=form,explode=true,name=names[0]"`
	Package       string `queryParam:"style=form,explode=true,name=package"`
	Placeholders0 string `queryParam:"style=form,explode=true,name=placeholders[0]"`
	Prompt        string `queryParam:"style=form,explode=true,name=prompt"`
	Public        int    `queryParam:"style=form,explode=true,name=public"`
	Tone          int    `queryParam:"style=form,explode=true,name=tone"`
	UserID        int    `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *CreatePromptTemplateRequest) GetActivate() int {
	if o == nil {
		return 0
	}
	return o.Activate
}

func (o *CreatePromptTemplateRequest) GetCategory() string {
	if o == nil {
		return ""
	}
	return o.Category
}

func (o *CreatePromptTemplateRequest) GetCode0() string {
	if o == nil {
		return ""
	}
	return o.Code0
}

func (o *CreatePromptTemplateRequest) GetEdit() int {
	if o == nil {
		return 0
	}
	return o.Edit
}

func (o *CreatePromptTemplateRequest) GetInputField0() string {
	if o == nil {
		return ""
	}
	return o.InputField0
}

func (o *CreatePromptTemplateRequest) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

func (o *CreatePromptTemplateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreatePromptTemplateRequest) GetNames0() string {
	if o == nil {
		return ""
	}
	return o.Names0
}

func (o *CreatePromptTemplateRequest) GetPackage() string {
	if o == nil {
		return ""
	}
	return o.Package
}

func (o *CreatePromptTemplateRequest) GetPlaceholders0() string {
	if o == nil {
		return ""
	}
	return o.Placeholders0
}

func (o *CreatePromptTemplateRequest) GetPrompt() string {
	if o == nil {
		return ""
	}
	return o.Prompt
}

func (o *CreatePromptTemplateRequest) GetPublic() int {
	if o == nil {
		return 0
	}
	return o.Public
}

func (o *CreatePromptTemplateRequest) GetTone() int {
	if o == nil {
		return 0
	}
	return o.Tone
}

func (o *CreatePromptTemplateRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

// CreatePromptTemplateCreatePromptTemplate - OK
type CreatePromptTemplateCreatePromptTemplate struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (o *CreatePromptTemplateCreatePromptTemplate) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *CreatePromptTemplateCreatePromptTemplate) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

type CreatePromptTemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreatePromptTemplate *CreatePromptTemplateCreatePromptTemplate
	Headers              map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePromptTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePromptTemplateResponse) GetCreatePromptTemplate() *CreatePromptTemplateCreatePromptTemplate {
	if o == nil {
		return nil
	}
	return o.CreatePromptTemplate
}

func (o *CreatePromptTemplateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreatePromptTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePromptTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
