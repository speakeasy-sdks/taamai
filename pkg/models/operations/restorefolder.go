// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var RestorefolderServerList = []string{
	"https://taam.ai/api",
}

type RestorefolderRequest struct {
	FolderID int `queryParam:"style=form,explode=true,name=folder_id"`
	UserID   int `queryParam:"style=form,explode=true,name=user_id"`
}

func (o *RestorefolderRequest) GetFolderID() int {
	if o == nil {
		return 0
	}
	return o.FolderID
}

func (o *RestorefolderRequest) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

type RestorefolderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RestorefolderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RestorefolderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RestorefolderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
