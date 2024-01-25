# Templates
(*Templates*)

### Available Operations

* [AllTemplates](#alltemplates) - All Templates
* [FavTemplates](#favtemplates) - Fav Templates
* [GenerateTemplate](#generatetemplate) - Generate Template
* [ProcessTemplate](#processtemplate) - Process Template
* [TemplateDetail](#templatedetail) - Template Detail
* [Templategroups](#templategroups) - Template groups

## AllTemplates

All Templates

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Templates.AllTemplates(ctx, operations.AllTemplatesRequest{
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AllTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.AllTemplatesRequest](../../pkg/models/operations/alltemplatesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.AllTemplatesResponse](../../pkg/models/operations/alltemplatesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## FavTemplates

Fav Templates

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
    res, err := s.Templates.FavTemplates(ctx, operations.FavTemplatesRequest{
        UserID: 12,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.FavTemplatesRequest](../../pkg/models/operations/favtemplatesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.FavTemplatesResponse](../../pkg/models/operations/favtemplatesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GenerateTemplate

Generate Template

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Templates.GenerateTemplate(ctx, operations.GenerateTemplateRequest{
        Creativity: 0.5,
        Description: "tell me about code",
        FolderID: 1,
        Language: "en-US",
        MaxResults: 1,
        TemplateCode: "EEKZF",
        Title: "Tsmokeshope",
        UserID: 40,
        Words: 100,
        WorkbookID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GenerateTemplateRequest](../../pkg/models/operations/generatetemplaterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.GenerateTemplateResponse](../../pkg/models/operations/generatetemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ProcessTemplate

Process Template

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"github.com/speakeasy-sdks/taamai/pkg/models/operations"
	"log"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Templates.ProcessTemplate(ctx, operations.ProcessTemplateRequest{
        ContentID: 286,
        MaxResults: 1,
        MaxWords: 100,
        Temperature: 0.5,
        UserID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProcessTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ProcessTemplateRequest](../../pkg/models/operations/processtemplaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.ProcessTemplateResponse](../../pkg/models/operations/processtemplateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## TemplateDetail

Template Detail

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
    res, err := s.Templates.TemplateDetail(ctx, operations.TemplateDetailRequest{
        TemplateID: 1,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.TemplateDetailRequest](../../pkg/models/operations/templatedetailrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.TemplateDetailResponse](../../pkg/models/operations/templatedetailresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Templategroups

Template groups

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/taamai/pkg/models/shared"
	"github.com/speakeasy-sdks/taamai"
	"context"
	"log"
	"net/http"
)

func main() {
    s := taamai.New(
        taamai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Templates.Templategroups(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |


### Response

**[*operations.TemplategroupsResponse](../../pkg/models/operations/templategroupsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
