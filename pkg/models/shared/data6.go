// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Data6 struct {
	AudioType        string  `json:"audio_type"`
	Characters       int     `json:"characters"`
	CreatedAt        string  `json:"created_at"`
	DeletedAt        *string `json:"deleted_at"`
	ExpiresAt        *string `json:"expires_at"`
	FileName         string  `json:"file_name"`
	Gender           string  `json:"gender"`
	ID               int     `json:"id"`
	Language         string  `json:"language"`
	LanguageFlag     string  `json:"language_flag"`
	Mode             string  `json:"mode"`
	PlanType         string  `json:"plan_type"`
	Project          *string `json:"project"`
	ResultExt        string  `json:"result_ext"`
	ResultURL        string  `json:"result_url"`
	Storage          string  `json:"storage"`
	Text             string  `json:"text"`
	TextRaw          string  `json:"text_raw"`
	Title            string  `json:"title"`
	UpdatedAt        string  `json:"updated_at"`
	UserID           int     `json:"user_id"`
	Vendor           string  `json:"vendor"`
	VendorID         string  `json:"vendor_id"`
	Voice            string  `json:"voice"`
	VoiceID          string  `json:"voice_id"`
	VoiceType        string  `json:"voice_type"`
	WorkbookFolderID int     `json:"workbook_folder_id"`
	WorkbookID       int     `json:"workbook_id"`
}

func (o *Data6) GetAudioType() string {
	if o == nil {
		return ""
	}
	return o.AudioType
}

func (o *Data6) GetCharacters() int {
	if o == nil {
		return 0
	}
	return o.Characters
}

func (o *Data6) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *Data6) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Data6) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *Data6) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *Data6) GetGender() string {
	if o == nil {
		return ""
	}
	return o.Gender
}

func (o *Data6) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Data6) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

func (o *Data6) GetLanguageFlag() string {
	if o == nil {
		return ""
	}
	return o.LanguageFlag
}

func (o *Data6) GetMode() string {
	if o == nil {
		return ""
	}
	return o.Mode
}

func (o *Data6) GetPlanType() string {
	if o == nil {
		return ""
	}
	return o.PlanType
}

func (o *Data6) GetProject() *string {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *Data6) GetResultExt() string {
	if o == nil {
		return ""
	}
	return o.ResultExt
}

func (o *Data6) GetResultURL() string {
	if o == nil {
		return ""
	}
	return o.ResultURL
}

func (o *Data6) GetStorage() string {
	if o == nil {
		return ""
	}
	return o.Storage
}

func (o *Data6) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

func (o *Data6) GetTextRaw() string {
	if o == nil {
		return ""
	}
	return o.TextRaw
}

func (o *Data6) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Data6) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *Data6) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *Data6) GetVendor() string {
	if o == nil {
		return ""
	}
	return o.Vendor
}

func (o *Data6) GetVendorID() string {
	if o == nil {
		return ""
	}
	return o.VendorID
}

func (o *Data6) GetVoice() string {
	if o == nil {
		return ""
	}
	return o.Voice
}

func (o *Data6) GetVoiceID() string {
	if o == nil {
		return ""
	}
	return o.VoiceID
}

func (o *Data6) GetVoiceType() string {
	if o == nil {
		return ""
	}
	return o.VoiceType
}

func (o *Data6) GetWorkbookFolderID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookFolderID
}

func (o *Data6) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}
