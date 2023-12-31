// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/taamai/pkg/utils"
)

type Data4WordsType string

const (
	Data4WordsTypeStr   Data4WordsType = "str"
	Data4WordsTypeInt32 Data4WordsType = "int32"
	Data4WordsTypeAny   Data4WordsType = "any"
)

type Data4Words struct {
	Str   *string
	Int32 *int
	Any   interface{}

	Type Data4WordsType
}

func CreateData4WordsStr(str string) Data4Words {
	typ := Data4WordsTypeStr

	return Data4Words{
		Str:  &str,
		Type: typ,
	}
}

func CreateData4WordsInt32(int32T int) Data4Words {
	typ := Data4WordsTypeInt32

	return Data4Words{
		Int32: &int32T,
		Type:  typ,
	}
}

func CreateData4WordsAny(any interface{}) Data4Words {
	typ := Data4WordsTypeAny

	return Data4Words{
		Any:  &any,
		Type: typ,
	}
}

func (u *Data4Words) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = Data4WordsTypeStr
		return nil
	}

	int32Var := new(int)
	if err := utils.UnmarshalJSON(data, &int32Var, "", true, true); err == nil {
		u.Int32 = int32Var
		u.Type = Data4WordsTypeInt32
		return nil
	}

	any := new(interface{})
	if err := utils.UnmarshalJSON(data, &any, "", true, true); err == nil {
		u.Any = any
		u.Type = Data4WordsTypeAny
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Data4Words) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Int32 != nil {
		return utils.MarshalJSON(u.Int32, "", true)
	}

	if u.Any != nil {
		return utils.MarshalJSON(u.Any, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Data4 struct {
	CreatedAt        string      `json:"created_at"`
	DeletedAt        *string     `json:"deleted_at"`
	Draft            int         `json:"draft"`
	Group            string      `json:"group"`
	Icon             string      `json:"icon"`
	ID               int         `json:"id"`
	InputText        *string     `json:"input_text"`
	Language         string      `json:"language"`
	LanguageFlag     string      `json:"language_flag"`
	LanguageName     string      `json:"language_name"`
	Model            *string     `json:"model"`
	PlanType         string      `json:"plan_type"`
	ResultText       *string     `json:"result_text"`
	Tags             *string     `json:"tags"`
	TemplateCode     *string     `json:"template_code"`
	TemplateName     string      `json:"template_name"`
	Title            *string     `json:"title"`
	Tokens           int         `json:"tokens"`
	UpdatedAt        string      `json:"updated_at"`
	UserID           int         `json:"user_id"`
	Words            *Data4Words `json:"words"`
	Workbook         *string     `json:"workbook"`
	WorkbookFolderID int         `json:"workbook_folder_id"`
	WorkbookID       int         `json:"workbook_id"`
}

func (o *Data4) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *Data4) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Data4) GetDraft() int {
	if o == nil {
		return 0
	}
	return o.Draft
}

func (o *Data4) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

func (o *Data4) GetIcon() string {
	if o == nil {
		return ""
	}
	return o.Icon
}

func (o *Data4) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Data4) GetInputText() *string {
	if o == nil {
		return nil
	}
	return o.InputText
}

func (o *Data4) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

func (o *Data4) GetLanguageFlag() string {
	if o == nil {
		return ""
	}
	return o.LanguageFlag
}

func (o *Data4) GetLanguageName() string {
	if o == nil {
		return ""
	}
	return o.LanguageName
}

func (o *Data4) GetModel() *string {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *Data4) GetPlanType() string {
	if o == nil {
		return ""
	}
	return o.PlanType
}

func (o *Data4) GetResultText() *string {
	if o == nil {
		return nil
	}
	return o.ResultText
}

func (o *Data4) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Data4) GetTemplateCode() *string {
	if o == nil {
		return nil
	}
	return o.TemplateCode
}

func (o *Data4) GetTemplateName() string {
	if o == nil {
		return ""
	}
	return o.TemplateName
}

func (o *Data4) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Data4) GetTokens() int {
	if o == nil {
		return 0
	}
	return o.Tokens
}

func (o *Data4) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *Data4) GetUserID() int {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *Data4) GetWords() *Data4Words {
	if o == nil {
		return nil
	}
	return o.Words
}

func (o *Data4) GetWorkbook() *string {
	if o == nil {
		return nil
	}
	return o.Workbook
}

func (o *Data4) GetWorkbookFolderID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookFolderID
}

func (o *Data4) GetWorkbookID() int {
	if o == nil {
		return 0
	}
	return o.WorkbookID
}
