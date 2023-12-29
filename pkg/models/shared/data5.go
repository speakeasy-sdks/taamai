// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Data5 struct {
	CreatedAt   string  `json:"created_at"`
	DeletedAt   string  `json:"deleted_at"`
	Description *string `json:"description"`
	Icon        string  `json:"icon"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Status      int     `json:"status"`
	UpdatedAt   string  `json:"updated_at"`
	UserID      int     `json:"user_id"`
	WorkbookID  int     `json:"workbook_id"`
}

func (o *Data5) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *Data5) GetDeletedAt() string {
	if o == nil {
		return ""
	}
	return o.DeletedAt
}

func (o *Data5) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Data5) GetIcon() string {
	if o == nil {
		return ""
	}
	return o.Icon
}

func (o *Data5) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Data5) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Data5) GetStatus() int {
	if o == nil {
		return 0
	}
	return o.Status
}

func (o *Data5) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *Data5) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *Data5) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}
