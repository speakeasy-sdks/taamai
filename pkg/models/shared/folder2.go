// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Folder2 struct {
	CreatedAt   string  `json:"created_at"`
	DeletedAt   *string `json:"deleted_at"`
	Description *string `json:"description"`
	Icon        string  `json:"icon"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Status      int     `json:"status"`
	UpdatedAt   string  `json:"updated_at"`
	UserID      int     `json:"user_id"`
	WorkbookID  int     `json:"workbook_id"`
}

func (o *Folder2) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *Folder2) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Folder2) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Folder2) GetIcon() string {
	if o == nil {
		return ""
	}
	return o.Icon
}

func (o *Folder2) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Folder2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Folder2) GetStatus() int {
	if o == nil {
		return 0
	}
	return o.Status
}

func (o *Folder2) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *Folder2) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *Folder2) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}
