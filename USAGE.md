<!-- Start SDK Example Usage -->


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
		taamai.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.AddonFeatures.GenerateCode(ctx, operations.GenerateCodeRequest{
		Document:     "new checking",
		Instructions: "generate a code to store image",
		Language:     "php",
		UserID:       1,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.GenerateCode != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->