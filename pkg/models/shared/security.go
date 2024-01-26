// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	Bearer string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *Security) GetBearer() string {
	if o == nil {
		return ""
	}
	return o.Bearer
}
