// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package taamai

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"github.com/speakeasy-sdks/taamai/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/taamai/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type misc struct {
	sdkConfiguration sdkConfiguration
}

func newMisc(sdkConfig sdkConfiguration) *misc {
	return &misc{
		sdkConfiguration: sdkConfig,
	}
}

// AllCategories - All Categories
func (s *misc) AllCategories(ctx context.Context, request operations.AllCategoriesRequest, opts ...operations.Option) (*operations.AllCategoriesResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionServerURL,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(operations.AllCategoriesServerList[0], map[string]string{})
	if o.ServerURL != nil {
		baseURL = *o.ServerURL
	}

	url := strings.TrimSuffix(baseURL, "/") + "/all/categories"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.AllCategoriesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
