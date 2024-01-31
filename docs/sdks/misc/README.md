# Misc
(*Misc*)

### Available Operations

* [AllCategories](#allcategories) - All Categories

## AllCategories

All Categories

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Misc.AllCategories(ctx, operations.AllCategoriesRequest{
        UserID: 40,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AllCategoriesRequest](../../pkg/models/operations/allcategoriesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.AllCategoriesResponse](../../pkg/models/operations/allcategoriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
