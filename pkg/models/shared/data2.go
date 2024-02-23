// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Data2 struct {
	CreatedAt string `json:"created_at"`
	Default   int    `json:"default"`
	DeletedAt string `json:"deleted_at"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
	UserID    int    `json:"user_id"`
}

func (o *Data2) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *Data2) GetDefault() int {
	if o == nil {
		return 0
	}
	return o.Default
}

func (o *Data2) GetDeletedAt() string {
	if o == nil {
		return ""
	}
	return o.DeletedAt
}

func (o *Data2) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Data2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Data2) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *Data2) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}
