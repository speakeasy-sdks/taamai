// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var GenerateprompttemplateServerList = []string{
	"http://127.0.0.1:8000/api",
}

type GenerateprompttemplateRequest struct {
	Creativity   float64 `queryParam:"style=form,explode=true,name=creativity"`
	Description  string  `queryParam:"style=form,explode=true,name=description"`
	FolderID     int     `queryParam:"style=form,explode=true,name=folder_id"`
	Language     string  `queryParam:"style=form,explode=true,name=language"`
	MaxResults   int     `queryParam:"style=form,explode=true,name=max_results"`
	TemplateCode string  `queryParam:"style=form,explode=true,name=template_code"`
	Text1        string  `queryParam:"style=form,explode=true,name=text1"`
	Title        string  `queryParam:"style=form,explode=true,name=title"`
	UserID       int     `queryParam:"style=form,explode=true,name=user_id"`
	Words        int     `queryParam:"style=form,explode=true,name=words"`
	WorkbookID   int     `queryParam:"style=form,explode=true,name=workbook_id"`
}

func (o *GenerateprompttemplateRequest) GetCreativity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Creativity
}

func (o *GenerateprompttemplateRequest) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *GenerateprompttemplateRequest) GetFolderID() int {
	if o == nil {
		return 0
	}
	return o.FolderID
}

func (o *GenerateprompttemplateRequest) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

func (o *GenerateprompttemplateRequest) GetMaxResults() int {
	if o == nil {
		return 0
	}
	return o.MaxResults
}

func (o *GenerateprompttemplateRequest) GetTemplateCode() string {
	if o == nil {
		return ""
	}
	return o.TemplateCode
}

func (o *GenerateprompttemplateRequest) GetText1() string {
	if o == nil {
		return ""
	}
	return o.Text1
}

func (o *GenerateprompttemplateRequest) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *GenerateprompttemplateRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *GenerateprompttemplateRequest) GetWords() int {
	if o == nil {
		return 0
	}
	return o.Words
}

func (o *GenerateprompttemplateRequest) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}

// GenerateprompttemplateGenerateprompttemplate - OK
type GenerateprompttemplateGenerateprompttemplate struct {
	ID          int    `json:"id"`
	MaxResults  string `json:"max_results"`
	MaxWords    string `json:"max_words"`
	Status      string `json:"status"`
	Temperature string `json:"temperature"`
}

func (o *GenerateprompttemplateGenerateprompttemplate) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GenerateprompttemplateGenerateprompttemplate) GetMaxResults() string {
	if o == nil {
		return ""
	}
	return o.MaxResults
}

func (o *GenerateprompttemplateGenerateprompttemplate) GetMaxWords() string {
	if o == nil {
		return ""
	}
	return o.MaxWords
}

func (o *GenerateprompttemplateGenerateprompttemplate) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GenerateprompttemplateGenerateprompttemplate) GetTemperature() string {
	if o == nil {
		return ""
	}
	return o.Temperature
}

type GenerateprompttemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	Generateprompttemplate *GenerateprompttemplateGenerateprompttemplate
	Headers                map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GenerateprompttemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateprompttemplateResponse) GetGenerateprompttemplate() *GenerateprompttemplateGenerateprompttemplate {
	if o == nil {
		return nil
	}
	return o.Generateprompttemplate
}

func (o *GenerateprompttemplateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GenerateprompttemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateprompttemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
