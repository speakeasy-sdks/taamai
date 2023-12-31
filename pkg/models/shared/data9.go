// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Data9 struct {
	Code             string  `json:"code"`
	CreatedAt        string  `json:"created_at"`
	DeletedAt        *string `json:"deleted_at"`
	ID               int     `json:"id"`
	Instructions     string  `json:"instructions"`
	Model            string  `json:"model"`
	Title            string  `json:"title"`
	UpdatedAt        string  `json:"updated_at"`
	UserID           int     `json:"user_id"`
	WorkbookFolderID int     `json:"workbook_folder_id"`
	WorkbookID       int     `json:"workbook_id"`
}

func (o *Data9) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *Data9) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *Data9) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Data9) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Data9) GetInstructions() string {
	if o == nil {
		return ""
	}
	return o.Instructions
}

func (o *Data9) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *Data9) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Data9) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *Data9) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *Data9) GetWorkbookFolderID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookFolderID
}

func (o *Data9) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}
