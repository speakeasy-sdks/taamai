<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/taamai"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"log"
)

func main() {
	s := taamai.New(
		taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.WorkbookAndFolders.AddandremovefromfavDocument(ctx, operations.AddandremovefromfavDocumentRequest{
		ID:     6,
		Type:   "document",
		UserID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.AddandremovefromfavDocument != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->